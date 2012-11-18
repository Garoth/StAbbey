package order

/* Representation of player/monster abilities in the game */

import (
    "log"

    "stabbey/interfaces"
)

var ACTIONS = map[byte] func(partialAction *Action) {
    '.' : IdleAction,
    'm' : MoveAction,
    'p' : PushAction,
}

type Action struct {
    actionString string
    count int
    act func(e interfaces.Entity, g interfaces.Game)
}

/* Smartly creates a new action based on the given action string */
func NewAction(at string) interfaces.Action {

    me := &Action{}
    me.actionString = at
    me.count = 0
    me.act = func(e interfaces.Entity, g interfaces.Game) {
    }
    ACTIONS[at[0]](me)

    return me
}

func (a *Action) ActionString() string {
    return a.actionString
}

/* Wrapper around the act member to work with interfaces */
func (a *Action) Act(e interfaces.Entity, g interfaces.Game) {
    a.act(e, g)
}

/* Makes you do nothing for one turn */
func IdleAction(me *Action) {
}

/* Moves your entity over one */
func MoveAction(me *Action) {
    me.act = func(e interfaces.Entity, g interfaces.Game) {
        boardId, x, y := e.GetPosition()

        x2, y2 := getDirectionCoords(me.actionString[1], x, y)
        if g.IsSpaceEmpty(x2, y2) {
            e.SetPosition(boardId, x2, y2)
        } else {
            log.Printf("Couldn't %v", me.actionString)
        }
    }
}

/* Pushes a neigbouring entity over one */
func PushAction(me *Action) {
    me.act = func(e interfaces.Entity, g interfaces.Game) {
        boardId, x, y := e.GetPosition()
        x2, y2 := getDirectionCoords(me.actionString[1], x, y)

        if entity := g.GetEntityByLocation(boardId, x2, y2); entity != nil {
            x3, y3 := getDirectionCoords(me.actionString[1], x2, y2)
            if g.IsSpaceEmpty(x3, y3) {
                entity.SetPosition(boardId, x3, y3)
            } else {
                log.Printf("Couldn't push %v by %v", entity.GetName(),
                    me.actionString)
            }
        } else {
            log.Printf("Nothing to push at %v", me.actionString)
        }
    }
}

/* Reads a character like 'r' and changes the given x/y to the adjacent
 * tile based on the direction given. i.e. 'r' would add 1 to x */
func getDirectionCoords(direction byte, x, y int) (int, int) {
    if direction == 'r' {
        return x + 1, y
    } else if direction == 'l' {
        return x - 1, y
    } else if direction == 'u' {
        return x, y - 1
    } else if direction == 'd' {
        return x, y + 1
    } else {
        log.Fatalln("Invalid direction given", direction)
    }

    return x, y
}
