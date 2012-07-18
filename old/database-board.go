package stabbey

import (
    "appengine/datastore"
)

type DatabaseBoard struct {
    Id int
    Layer0, Layer1, Layer2, Layer3, Layer4, Layer5, Layer6, Layer7 string
}

func NewDatabaseBoard(b *Board) *DatabaseBoard {
    db := &DatabaseBoard{}

    db.Id = b.Id
    db.Layer0 = b.Layer0
    db.Layer1 = b.Layer1
    db.Layer2 = b.Layer2
    db.Layer3 = b.Layer3
    db.Layer4 = b.Layer4
    db.Layer5 = b.Layer5
    db.Layer6 = b.Layer6
    db.Layer7 = b.Layer7

    return db
}

func (db *DatabaseBoard) Save(c *Context) error {
    _, e := datastore.Put(c.GAEContext, GetBoardKey(c, db.Id), db)

    if e != nil {
        c.GAEContext.Errorf("Error saving Board: %v", e)
    } else {
        c.GAEContext.Infof("Successfully saved board %v", db.Id)
    }

    return e;
}

func LoadDatabaseBoard(c *Context, id int) *DatabaseBoard {
    db := &DatabaseBoard{}
    e := datastore.Get(c.GAEContext, GetBoardKey(c, id), db)

    if e != nil {
        c.GAEContext.Errorf("Error loading Board: %v", e)
    } else {
    }

    return db;
}
