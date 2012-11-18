package entity

import (
    "stabbey/interfaces"
    "stabbey/uidgenerator"
)

var UIDG = uidgenerator.New()

type Entity struct {
    EntityId, BoardId, X, Y int
    Name, Type string
    ActionQueue []interfaces.Action
}

func New(entid int) *Entity {
    e := &Entity{}
    e.SetEntityId(entid)
    e.ActionQueue = make([]interfaces.Action, 0, 10)
    return e
}

func (e *Entity) SetEntityId(id int) {
    e.EntityId = id
}

func (e *Entity) GetEntityId() int {
    return e.EntityId
}

func (e *Entity) GetPosition() (boardid, x, y int) {
    return e.BoardId, e.X, e.Y
}

func (e *Entity) SetPosition(boardid, x, y int) {
    e.BoardId = boardid
    e.X = x
    e.Y = y
}

func (e *Entity) SetName(name string) {
    e.Name = name
}

func (e *Entity) GetName() string {
    return e.Name
}

func (e *Entity) SetType(t string) {
    e.Type = t
}

func (e *Entity) GetType() string {
    return e.Type
}

func (e *Entity) GetActionQueue() []interfaces.Action {
    return e.ActionQueue
}

func (e *Entity) GetStringActionQueue() []string {
    q := make([]string, len(e.GetActionQueue()))

    for i := 0; i < len(e.GetActionQueue()); i++ {
        q[i] = e.GetActionQueue()[i].ActionString()
    }

    return q
}

func (e *Entity) SetActionQueue(aq []interfaces.Action) {
    e.ActionQueue = aq
}

func (e *Entity) PopAction() interfaces.Action {
    if len(e.ActionQueue) > 0 {
        a := e.ActionQueue[0]
        e.ActionQueue = e.ActionQueue[1:]
        return a
    }
    return nil
}

