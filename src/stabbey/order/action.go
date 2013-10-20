package order

/* Representation of player/monster abilities in the game */

import (
    "log"
    "math/rand"
    "stabbey/interfaces"
)

/* Map of the base command string 'verbs' to which functions handle them */
var ACTIONS = map[byte] func(partialAction *Action) {
    '.' : IdleAction,
    'm' : MoveAction,
    'p' : PushAction,
    '*' : PunchAction,
    'l' : LeapAction,
    'd' : DrunkAction,
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
        x2, y2 := getDirectionCoords(me.actionString[1], x, y, 1)

        if g.CanMoveToSpace(boardId, x2, y2) {
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
        x2, y2 := getDirectionCoords(me.actionString[1], x, y, 1)

        if entity := getAliveTangibleEntity(boardId, x2, y2, g); entity != nil {
            x3, y3 := getDirectionCoords(me.actionString[1], x2, y2, 1)
            if g.CanMoveToSpace(boardId, x3, y3) {
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
        x2, y2 := getDirectionCoords(me.actionString[1], x, y, 1)

        if entity := getAliveTangibleEntity(boardId, x2, y2, g); entity != nil {
            entity.ChangeArdour(-10)
        } else {
            log.Printf("Nothing to punch with %v", me.actionString)
        }
    }
}

/* Leaps in a direction, over other entities */
func LeapAction(me *Action) {
    me.shortDesc = "Leap"
    me.longDesc = "You fall, but miss the ground for a while."

    me.act = func(e interfaces.Entity, g interfaces.Game) {
        leapLength := 3

        boardId, xOrig, yOrig := e.GetPosition()
        x, y := xOrig, yOrig
        xDest, yDest := -1, -1
        iLeft := leapLength

        /* Find landing spot */
        for i := 1; i <= leapLength; i++ {
            x, y = getDirectionCoords(me.actionString[1], xOrig, yOrig, i)

            if g.IsWall(boardId, x, y) {
                iLeft = i - 1
                break
            }
        }

        for i := iLeft; i >= 0; i-- {
            x, y = getDirectionCoords(me.actionString[1], xOrig, yOrig, i)

            if g.CanMoveToSpace(boardId, x, y) {
                xDest, yDest = x, y
                iLeft = i
                break
            }
        }

        /* Land there if possible */
        if xDest == -1 && yDest == -1 {
            log.Println("Leap by", e.GetName(), "failed -- no available space")
            return
        } else {
            log.Println(e.GetName(), "lept to", xDest, yDest)
            e.SetPosition(boardId, xDest, yDest)
        }

        /* Trigger trodden on all entities along the way manually */
        for iLeft = iLeft - 1; iLeft > 0; iLeft-- {
            x, y := getDirectionCoords(me.actionString[1], xOrig, yOrig, iLeft)

            for _, entity := range g.GetEntitiesAtSpace(boardId, x, y) {
                entity.Trodden(e);
            }
        }
    }
}

/* Drunk movement.  Whiskey!!! */
func DrunkAction(me *Action) {
    me.shortDesc = "Durnked "
    me.longDesc  = "You confidently move with the grace of a hippo."
    me.availableDirections = [5]bool{false, false, false, false, false}
  
    xOffset := rand.Intn(3) - 1;
    yOffset := rand.Intn(3) - 1;
      
    me.act = func(e interfaces.Entity, g interfaces.Game) {
        boardId, x, y := e.GetPosition()
        x = x + xOffset;
        y = y + yOffset;
  
        if g.CanMoveToSpace(boardId, x, y) {
            e.SetPosition(boardId, x, y)
        } else {
            log.Printf("Couldn't %v", me.actionString)
        }
    }
}

/* Reads a character like 'r' and changes the given x/y to the adjacent
 * tile based on the direction given. i.e. 'r' would add 1 to x.
 * The repeat setting how many tiles in that direction to go */
func getDirectionCoords(direction byte, x, y, repeat int) (int, int) {
    if direction == 'r' {
        return x + repeat, y
    } else if direction == 'l' {
        return x - repeat, y
    } else if direction == 'u' {
        return x, y - repeat
    } else if direction == 'd' {
        return x, y + repeat
    } else {
        log.Fatalln("Invalid direction given", direction)
    }

    return x, y
}

/* Checks entities on a space and returns the first non-dead, tangible one */
func getAliveTangibleEntity(boardId, x, y int,
        g interfaces.Game) interfaces.Entity {

    for _, entity := range g.GetEntitiesAtSpace(boardId, x, y) {
        if !entity.IsDead() && entity.IsTangible() {
            return entity
        }
    }

    return nil
}
