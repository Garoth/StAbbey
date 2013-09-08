package order

/* Representation of player/monster abilities in the game */

import (
    "log"

    "stabbey/interfaces"
)

/* Map of the base command string 'verbs' to which functions handle them */
var ACTIONS = map[byte] func(partialAction *Action) {
    '.' : IdleAction,
    'm' : MoveAction,
    'p' : PushAction,
    '*' : PunchAction,
}

type Action struct {
    /* Short encodded string representing the action */
    actionString string
    /* Descriptions, for UI */
    shortDesc, longDesc string
    /* Which direction(s) this action is usable - left, right, up, down, self */
    availableDirections [5]bool
    /* How many times the action is to be repeated */
    count int
    /* Function that'll perform the action on the world state */
    act func(e interfaces.Entity, g interfaces.Game)
}

/* Smartly creates a new action based on the given action string */
func NewAction(at string) interfaces.Action {

    me := &Action{}
    me.actionString = at
    me.shortDesc = "MISSING SHORT DESC"
    me.longDesc = "MISSING LONG DESC"
    me.availableDirections = [5]bool{true, true, true, true, false}
    me.count = 0
    me.act = func(e interfaces.Entity, g interfaces.Game) { }
    ACTIONS[at[0]](me)

    return me
}

func (a *Action) ActionString() string {
    return a.actionString
}

func (a *Action) ShortDescription() string {
    return a.shortDesc
}

func (a *Action) LongDescription() string {
    return a.longDesc
}

func (a *Action) AvailableDirections() [5]bool {
    return a.availableDirections
}

/* Wrapper around the act member to work with interfaces */
func (a *Action) Act(e interfaces.Entity, g interfaces.Game) {
    a.act(e, g)
}

/* Makes you do nothing for one turn */
func IdleAction(me *Action) {
    me.shortDesc = "Do Nothing"
    me.longDesc = "You stand still and contemplate reality."
    me.availableDirections = [5]bool{false, false, false, false, false}
}

/* Moves your entity over one */
func MoveAction(me *Action) {
    me.shortDesc = "Move"
    me.longDesc = "You bravely advance."

    me.act = func(e interfaces.Entity, g interfaces.Game) {
        boardId, x, y := e.GetPosition()

        x2, y2 := getDirectionCoords(me.actionString[1], x, y)
        if g.CanMoveToSpace(x2, y2) {
            ents := g.GetEntitiesAtSpace(boardId, x2, y2)
            for _, entity := range ents {
                entity.Trodden(e);
            }
            e.SetPosition(boardId, x2, y2)
        } else {
            log.Printf("Couldn't %v", me.actionString)
        }
    }
}

/* Pushes a neighbouring entity over one */
func PushAction(me *Action) {
    me.shortDesc = "Push"
    me.longDesc = "You put your weight into a mighty shove."

    me.act = func(e interfaces.Entity, g interfaces.Game) {
        boardId, x, y := e.GetPosition()
        x2, y2 := getDirectionCoords(me.actionString[1], x, y)

        if entity := g.GetEntityByLocation(boardId, x2, y2); entity != nil {
            x3, y3 := getDirectionCoords(me.actionString[1], x2, y2)
            if g.CanMoveToSpace(x3, y3) {
                entity.SetPosition(boardId, x3, y3)
            } else {
                log.Printf("Couldn't push %v by %v", entity.GetName(),
                    me.actionString)
            }
        } else {
            log.Printf("Nothing to push with %v", me.actionString)
        }
    }
}

/* Punches a neighbouring entity */
func PunchAction(me *Action) {
    me.shortDesc = "Punch"
    me.longDesc = "You punch wildly."

    me.act = func(e interfaces.Entity, g interfaces.Game) {
        boardId, x, y := e.GetPosition()
        x2, y2 := getDirectionCoords(me.actionString[1], x, y)

        if entity := g.GetEntityByLocation(boardId, x2, y2); entity != nil {
            entity.ChangeArdour(-10)
        } else {
            log.Printf("Nothing to punch with %v", me.actionString)
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
