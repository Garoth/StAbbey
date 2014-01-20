package entity

import (
    "log"
    "runtime/debug"

    "stabbey/interfaces"
    "stabbey/uidgenerator"
)

var UIDG = uidgenerator.New()

type Entity struct {
    EntityId, BoardId, X, Y, Ardour, MaxArdour int
    Game interfaces.Game
    Name, Type, Subtype string
    Tangible, Dead bool
    DeathFunction func()
    TroddenFunction func(interfaces.Entity)
    RepositionFunction func(fromBId, fromX, fromY, toBId, toX, toY int)
    TurnFunction func(int) bool
    ActionQueue []interfaces.Action
}

func New(entid int, g interfaces.Game) *Entity {
    e := &Entity{}
    e.SetEntityId(entid)
    e.BoardId = 0
    e.X = -1
    e.Y = -1
    e.Game = g
    e.MaxArdour = 50
    e.Ardour = 50
    e.Dead = false
    e.Tangible = true
    e.Name = "Some unspecified entity"
    e.Type = interfaces.ENTITY_TYPE_UNKNOWN
    e.Subtype = interfaces.ENTITY_SUBTYPE_UNKNOWN
    e.DeathFunction = func() {
        log.Println(e.GetName(), "has died")
    }
    e.TroddenFunction = func(by interfaces.Entity) {
        log.Println(e.GetName(), "was stepped on")
    }
    e.TurnFunction = func(tick int) bool {
        return false
    }
    e.RepositionFunction = func(fromBId, fromX, fromY, toBId, toX, toY int) {
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

func (e *Entity) SwapPositionWith(other interfaces.Entity) {
    if !e.IsTangible() || !other.IsTangible() {
        log.Fatalf("Swap only makes sense for tangible entities. Got %v & %v",
            e.GetName(), other.GetName())
    }

    /* It's crucial that two tagnible entities don't end up on same tile */
    boardId, x, y := e.GetPosition()
    boardId2, x2, y2 := other.GetPosition()
    e.BoardId, e.X, e.Y = boardId2, x2, y2
    /* Can work since both entities are technically in the same place */
    other.SetPosition(boardId, x, y)

    /* Manually trigger trodden functions since we didn't SetPosition */
    for _, entity := range e.Game.GetEntitiesAtSpace(e.GetPosition()) {
        entity.Trodden(e)
    }

    /* other's reposition function will be called by SetPosition */
    e.RepositionFunction(boardId, x, y, boardId2, x2, y2)
}

func (e *Entity) SetPosition(boardId, x, y int) {
    if e.IsTangible() && !e.Game.CanMoveToSpace(boardId, x, y) {
        debug.PrintStack()
        log.Fatalln("Can't move", e.GetName(), "to impossible place", x, y)
    }

    /* Ordering of operations matters here, since we want to see what
     * entities are on the target tile before we're on that tile
     * (to exclude ourselves). Also, we can't trigger the trodden
     * on those entities we're landing on before moving, since
     * it could be that landing on one repositions us right away */
    entitiesAtTarget := e.Game.GetEntitiesAtSpace(boardId, x, y)

    oldBId, oldX, oldY := e.GetPosition()
    e.BoardId, e.X, e.Y = boardId, x, y

    if e.IsTangible() {
        for _, entity := range entitiesAtTarget {
            entity.Trodden(e);
        }
    }

    e.RepositionFunction(oldBId, oldX, oldY, boardId, x, y)
}

func (e *Entity) GetPosition() (boardid, x, y int) {
    return e.BoardId, e.X, e.Y
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

func (e *Entity) SetSubtype(t string) {
    e.Subtype = t
}

func (e *Entity) GetSubtype() string {
    return e.Subtype
}

func (e *Entity) SetMaxArdour(ardour int) {
    e.MaxArdour = ardour

    if e.Ardour > e.MaxArdour {
        e.Ardour = e.MaxArdour
    }

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
        e.Die()
    }

    log.Println(e.Name, "changed ardour to", e.Ardour)
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
    e.Ardour = 0
    e.Dead = true;
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

/* Runs custom code for the entity's turn. Useful for programming monsters
 * and the like.
 *
 * Returns true if it did something that should be announced to the players.
 */
func (me *Entity) RunTurn(tick int) bool {
    return me.TurnFunction(tick)
}
