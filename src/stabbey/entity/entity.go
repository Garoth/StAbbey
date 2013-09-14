package entity

import (
    "log"

    "stabbey/interfaces"
    "stabbey/uidgenerator"
)

var UIDG = uidgenerator.New()

type Entity struct {
    EntityId, BoardId, X, Y, Ardour, MaxArdour int
    Game interfaces.Game
    Name, Type string
    Tangible, Dead bool
    DeathFunction func()
    TroddenFunction func(interfaces.Entity)
    TickFunction func(int)
    ActionQueue []interfaces.Action
}

func New(entid int, g interfaces.Game) *Entity {
    e := &Entity{}
    e.SetEntityId(entid)
    e.BoardId = 0
    e.X = 0
    e.Y = 0
    e.Game = g
    e.MaxArdour = 50
    e.Ardour = 50
    e.Dead = false
    e.Tangible = true
    e.Name = "Some unspecified entity"
    e.DeathFunction = func() {
        log.Println(e.GetName(), "has died")
    }
    e.TroddenFunction = func(by interfaces.Entity) {
        log.Println(e.GetName(), "was stepped on")
    }
    e.TickFunction = func(tick int) {
    }
    e.ActionQueue = make([]interfaces.Action, 0, 10)
    return e
}

/* TODO should force entity to be unique */
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
        e.Ardour = 0
        e.Dead = true;
        e.Die()
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
    if e.Ardour == 0 {
        e.Die()
    }
}

func (e *Entity) GetArdour() int {
    return e.Ardour
}

func (e *Entity) IsTangible() bool {
    return e.Tangible
}

func (e *Entity) SetTangible(tangible bool) {
    e.Tangible = tangible
}

func (e *Entity) IsDead() bool {
    return e.Dead
}

func (e *Entity) Die() {
    e.SetTangible(false)
    e.DeathFunction()
}

func (e *Entity) Trodden(by interfaces.Entity) {
    e.TroddenFunction(by)
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

func (me *Entity) WorldTick(tick int) {
    me.TickFunction(tick)
}
