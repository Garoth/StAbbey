package entity

import (
    "log"

    "stabbey/interfaces"
    "stabbey/uidgenerator"
)

var UIDG = uidgenerator.New()

type Entity struct {
    EntityId, BoardId, X, Y, Ardour, MaxArdour int
    Name, Type string
    Dead bool
    ActionQueue []interfaces.Action
}

func New(entid int) *Entity {
    e := &Entity{}
    e.SetEntityId(entid)
    e.MaxArdour = 50
    e.Ardour = 50
    e.Dead = false
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

func (e *Entity) SetMaxArdour(ardour int) {
    e.MaxArdour = ardour
    if e.MaxArdour < 0 {
        log.Println("Warning: attempt to set max ardour to below 0")
        e.MaxArdour = 1
    }
}

func (e *Entity) GetMaxArdour() int {
    return e.MaxArdour
}

func (e *Entity) ChangeArdour(difference int) int {
    e.Ardour += difference
    if e.Ardour > e.MaxArdour {
        e.Ardour = e.MaxArdour
    }

    if e.Ardour <= 0 {
        log.Println(e.Name, "has died!")
        e.Ardour = 0
        e.Dead = true;
    } else {
        e.Dead = false;
    }

    log.Println("Entity", e.Name, "changed ardour to", e.Ardour)
    return e.Ardour
}

func (e *Entity) SetArdour(ardour int) {
    e.Ardour = ardour
    if e.Ardour > e.MaxArdour {
        log.Println("Warning: attempt to set ardour to above max")
        e.Ardour = e.MaxArdour
    }
    if e.Ardour < 0 {
        log.Println("Warning: attempt to set ardour to below 0")
        e.Ardour = 0
    }
}

func (e *Entity) GetArdour() int {
    return e.Ardour
}

func (e *Entity) IsDead() bool {
    return e.Dead
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

