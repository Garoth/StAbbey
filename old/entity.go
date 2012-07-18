package stabbey

import (
    "appengine/datastore"
    "strconv"
)

type EntityStruct struct {
    EntityId, BoardId, X, Y int
    Name, Type string
}

/* A monster, player, or some special thing of that sort */
type Entity interface {
    /* Entity Ids must be unique */
    SetEntityId(id int)
    GetEntityId() int
    SetPosition(boardid, x, y int)
    GetPosition() (boardid, x, y int)
    SetName(name string)
    GetName() string
    SetType(t string)
    GetType() string
    Save(c *Context) error
}

func NewEntity(c *Context, entid int) Entity {
    es := &EntityStruct{}
    es.SetEntityId(entid)
    es.Save(c)
    return es
}

func NewEntityFromDatabase(de *DatabaseEntity) Entity {
    e := &EntityStruct{}
    e.SetName(de.Name)
    e.SetEntityId(de.EntityId)
    e.SetPosition(de.BoardId, de.X, de.Y)
    e.SetType(de.Type)
    return e
}

func (ep *EntityStruct) SetEntityId(id int) {
    ep.EntityId = id
}

func (ep *EntityStruct) GetEntityId() int {
    return ep.EntityId
}

func (ep *EntityStruct) GetPosition() (boardid, x, y int) {
    return ep.BoardId, ep.X, ep.Y
}

func (ep *EntityStruct) SetPosition(boardid, x, y int) {
    ep.BoardId = boardid
    ep.X = x
    ep.Y = y
}

func (ep *EntityStruct) SetName(name string) {
    ep.Name = name
}

func (ep *EntityStruct) GetName() string {
    return ep.Name
}

func (ep *EntityStruct) SetType(t string) {
    ep.Type = t
}

func (ep *EntityStruct) GetType() string {
    return ep.Type
}

func (ep *EntityStruct) Save (c *Context) error {
    return NewDatabaseEntity(ep).Save(c)
}

func LoadEntity(c *Context, id int) Entity {
    return NewEntityFromDatabase(LoadDatabaseEntity(c, id))
}

func GetEntityKey(c *Context, id int) *datastore.Key {
    return datastore.NewKey(c.GAEContext, "Entity", strconv.Itoa(id), 0, nil)
}
