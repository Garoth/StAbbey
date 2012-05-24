package stabbey

type EntityPosition struct {
    EntityId, BoardId, X, Y int
    Name string
}

/* A monster, player, or some special thing of that sort */
type Entity interface {
    SetPosition(boardid, x, y int)
    GetPosition() (boardid, x, y int)
    SetName(name string)
    GetName() string
    /* Entity Ids should be unique */
    SetEntityId(id int)
    GetEntityId() int
}

func (ep *EntityPosition) GetPosition() (boardid, x, y int) {
    return ep.BoardId, ep.X, ep.Y
}

func (ep *EntityPosition) SetPosition(boardid, x, y int) {
    ep.BoardId = boardid
    ep.X = x
    ep.Y = y
}

func (ep *EntityPosition) SetName(name string) {
    ep.Name = name
}

func (ep *EntityPosition) GetName() string {
    return ep.Name
}

func (ep *EntityPosition) SetEntityID(id int) {
    ep.EntityId = id
}

func (ep *EntityPosition) GetEntityID() int {
    return ep.EntityId
}
