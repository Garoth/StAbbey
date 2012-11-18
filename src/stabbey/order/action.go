package order

import (
    "log"

    "stabbey/interfaces"
)

var ACTIONS = map[byte] func(partialAction *Action) {
    'm' :  MoveAction,
}

type Action struct {
    actionString string
    count int
    act func(e interfaces.Entity, g interfaces.Game)
}

func NewAction(at string) interfaces.Action {

    me := &Action{}
    me.actionString = at
    me.count = 0
    ACTIONS[at[0]](me)

    return me
}

func (a *Action) ActionString() string {
    return a.actionString
}

func (a *Action) Act(e interfaces.Entity, g interfaces.Game) {
    a.act(e, g)
}

func MoveAction(me *Action) {
    me.act = func(e interfaces.Entity, g interfaces.Game) {
        boardId, x, y := e.GetPosition()

        if me.actionString == "mr" && g.IsSpaceEmpty(x + 1, y) {
            e.SetPosition(boardId, x + 1, y)
        } else if me.actionString == "ml" && g.IsSpaceEmpty(x - 1, y) {
            e.SetPosition(boardId, x - 1, y)
        } else if me.actionString == "mu" && g.IsSpaceEmpty(x, y - 1) {
            e.SetPosition(boardId, x, y - 1)
        } else if me.actionString == "md" && g.IsSpaceEmpty(x, y + 1) {
            e.SetPosition(boardId, x, y + 1)
        } else {
            log.Printf("Order %v did nothing.", me.actionString)
        }
    }
}
