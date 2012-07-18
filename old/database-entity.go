package stabbey

import (
    "appengine/datastore"
)

type DatabaseEntity struct {
    EntityId, BoardId, X, Y int
    Name, Type string
}

func NewDatabaseEntity(e Entity) *DatabaseEntity {
    de := &DatabaseEntity{}
    de.EntityId = e.GetEntityId();
    de.Name = e.GetName();
    de.BoardId, de.X, de.Y = e.GetPosition();
    de.Type = e.GetType();
    return de;
}

func (de *DatabaseEntity) Save(c *Context) error {
    _, e := datastore.Put(c.GAEContext, GetEntityKey(c, de.EntityId), de)

    if e != nil {
        c.GAEContext.Errorf("Error saving entity: %v", e)
    } else {
        c.GAEContext.Infof("Successfully saved entity %v (%v)", de.EntityId,
            de.Name)
    }

    return e;
}

func LoadDatabaseEntity(c *Context, id int) *DatabaseEntity {
    de := &DatabaseEntity{}
    e := datastore.Get(c.GAEContext, GetEntityKey(c, id), de)

    if e != nil {
        c.GAEContext.Errorf("Error loading Entity: %v", e)
    } else {
        c.GAEContext.Infof("Successfully loaded Entity %v (%v)", de.EntityId,
            de.Name)
    }

    return de;
}
