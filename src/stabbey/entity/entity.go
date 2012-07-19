package entity

import (
)

type Entity struct {
    EntityId, BoardId, X, Y int
    Name, Type string
}

func New(entid int) *Entity {
    e := &Entity{}
    e.SetEntityId(entid)
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
