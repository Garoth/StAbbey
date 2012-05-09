package stabbey

type Position struct {
    BoardId, X, Y int
}

/* A monster, player, or some special thing of that sort */
type Entity interface {
    SetPosition(Position)
    GetPosition() Position
    SetName(name string)
    GetName() string
    /* Entity Ids should be unique */
    SetEntityId(id int)
    GetEntityId() int
}
