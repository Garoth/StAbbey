package stabbey

import (
    "appengine"
    "appengine/datastore"
)

type Game struct {
    Players []string
    Boards []int
}

func NewGame() *Game {
    return &Game{}
}

/* Adds a player to the current game. Call once per player per game */
func (game *Game) AddPlayer(p *Player) {
    game.Players = append(game.Players, p.Id)
}

/* Adds a board to the current game. Call once per board per game */
func (game *Game) AddBoard(b *Board) {
    game.Boards = append(game.Boards, b.Id)
}

/* Gets the database key for the game */
func (game *Game) GetKey(context appengine.Context,
        gamekey string) *datastore.Key {

    return datastore.NewKey(context, "Game", gamekey, 0, nil)
}

/* Saves the game to the database */
func (game *Game) Save(context appengine.Context, gamekey string) error {
    _, e := datastore.Put(context, game.GetKey(context, gamekey), game)

    if e != nil {
        context.Errorf("Error saving Game: %v", e)
        return e
    }

    return nil
}

/* Loads the game from the database */
func (game *Game) Load(context appengine.Context, gamekey string) error {
    e := datastore.Get(context, game.GetKey(context, gamekey), game)

    if e != nil {
        context.Errorf("Error loading game: %v", e)
    }

    return e
}
