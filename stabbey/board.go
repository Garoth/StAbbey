package stabbey

import (
    "appengine"
    "appengine/datastore"
)

type Board struct {
    /* Unique game level number (i.e. floor number) */
    Id string
    /* Rendering layers of the map, going up in z-index */
    Layer0, Layer1, Layer2, Layer3, Layer4, Layer5, Layer6, Layer7 string
}

/* Creates a brand new board, for level "id" -- 0, 1, 2, etc */
func NewBoard(c *Context, id string) *Board {
    b := &Board{}
    b.Id = id
    NewDatabaseBoard(b).Save(c)
    b.MakeTestBoard(c, string(id))
    return b
}

func NewBoardFromDatabase(db *DatabaseBoard) *Board {
    b := &Board{}

    b.Id = db.Id
    b.Layer0 = db.Layer0
    b.Layer1 = db.Layer1
    b.Layer2 = db.Layer2
    b.Layer3 = db.Layer3
    b.Layer4 = db.Layer4
    b.Layer5 = db.Layer5
    b.Layer6 = db.Layer6
    b.Layer7 = db.Layer7

    return b
}

/* Returns the database key for the board */
func GetBoardKey(c *Context, boardId string) *datastore.Key {
    return datastore.NewKey(c.GAEContext, "Board" + boardId,
        c.Gamekey, 0, nil)
}

/* Save the board to the database */
func (b *Board) Save(c *Context) error {
    return NewDatabaseBoard(b).Save(c)
}

/* Load a board from the database -- by level id: 0, 1, 2, 3, etc */
func LoadBoard(c *Context, id string) *Board {
    return NewBoardFromDatabase(LoadDatabaseBoard(c, id))
}

func (b *Board) MakeTestBoard(c *Context, id string) {
    datastore.RunInTransaction(c.GAEContext, func(x appengine.Context) error {
        b := LoadBoard(c, id)
        b.Layer0 = "L--------------L" +
                   "|  |           |" +
                   "|  |           |" +
                   "|  |           |" +
                   "|  | ----------|" +
                   "|              |" +
                   "|              |" +
                   "|-----------L  |" +
                   "|           |  |" +
                   "|              |" +
                   "|           |  |" +
                   "L--------------L"
        b.Save(c)
        return nil // TODO
    }, nil)
}
