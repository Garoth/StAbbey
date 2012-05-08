package stabbey

import (
    "fmt"
    "appengine"
    "appengine/datastore"
)

type Board struct {
    /* Unique game level number (i.e. floor number) */
    Id int
    /* Number of x/y tiles on the board */
    Width, Height int
    /* Rendering layers of the map, going up in z-index */
    Layers [][]string
}

func NewBoard(level int) *Board {
    b := &Board{}
    b.Id = level
    b.Width = 16
    b.Height = 12
    return b
}

func (b *Board) MakeTestBoard() {
    fmt.Println("Making test board")
    b.Layers = append(b.Layers, []string{"xxxxxxxxxxxxxxxx",
                                         "xxxxxxxxxxxxxxxx",
                                         "xxxxxxxxxxxxxxxx",
                                         "xxxxxxxxxxxxxxxx",
                                         "xxxxxxxxxxxxxxxx",
                                         "xxxxxxxxxxxxxxxx",
                                         "xxxxxxxxxxxxxxxx",
                                         "xxxxxxxxxxxxxxxx",
                                         "xxxxxxxxxxxxxxxx",
                                         "xxxxxxxxxxxxxxxx",
                                         "xxxxxxxxxxxxxxxx",
                                         "xxxxxxxxxxxxxxxx"})
}

/* Returns the database key for the board */
func (b *Board) GetKey(context appengine.Context,
        gamekey string) *datastore.Key {

    return datastore.NewKey(context, "Board" + string(b.Id), gamekey, 0, nil)
}

/* Save the board to the database */
func (b *Board) Save(context appengine.Context, gamekey string) error {
    _, e := datastore.Put(context, b.GetKey(context, gamekey), b)

    if e != nil {
        context.Errorf("Error saving Board: %v", e)
    }

    return e;
}

/* Load a board from the database */
func (b *Board) Load(context appengine.Context, gamekey string) error {
    e := datastore.Get(context, b.GetKey(context, gamekey), b)

    if e != nil {
        context.Errorf("Error loading Board: %v", e)
    }

    return e
}
