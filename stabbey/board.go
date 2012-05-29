package stabbey

import (
    "appengine/datastore"
)

type Board struct {
    /* Unique game level number (i.e. floor number) */
    Id int
    /* Rendering layers of the map, going up in z-index */
    Layer0, Layer1, Layer2, Layer3, Layer4, Layer5, Layer6, Layer7 string
}

func NewBoard(level int) *Board {
    b := &Board{}
    b.Id = level
    b.MakeTestBoard()
    return b
}

/* Returns the database key for the board */
func GetBoardKey(c *Context, boardId string) *datastore.Key {
    return datastore.NewKey(c.GAEContext, "Board" + boardId,
        c.Gamekey, 0, nil)
}

/* Load a board from the database */
func LoadBoard(c *Context, level string) *Board {
    b := &Board{}
    e := datastore.Get(c.GAEContext, GetBoardKey(c, level), b)

    if e != nil {
        c.GAEContext.Errorf("Error loading Board: %v", e)
    }

    return b
}

func (b *Board) MakeTestBoard() {
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
}

/* Save the board to the database */
func (b *Board) Save(c *Context) error {
    _, e := datastore.Put(c.GAEContext, GetBoardKey(c, string(b.Id)), b)

    if e != nil {
        c.GAEContext.Errorf("Error saving Board: %v", e)
    }

    return e
}
