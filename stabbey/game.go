package stabbey

import (
    "os"
    "appengine"
    "appengine/datastore"
)

type Game struct {
    Players []string
}

func NewGame() *Game {
    return &Game{}
}

/* Adds a player to the current game. Call once per player per game */
func (game *Game) AddPlayer(p *Player) {
    game.Players = append(game.Players, p.Id)
}

/* Gets the database key for the game */
func (game *Game) GetKey(context appengine.Context,
        gamekey string) *datastore.Key {

    return datastore.NewKey(context, "Game", gamekey, 0, nil)
}

/* Saves the game to the database */
func (game *Game) Save(context appengine.Context, gamekey string) os.Error {
    _, e := datastore.Put(context, game.GetKey(context, gamekey), game)

    if e != nil{
        context.Errorf("Error saving Game: %v", e)
        return e
    }

    for _, Id := range game.Players {
        p := NewPlayer(Id)
        p.Save(context, gamekey)
    }

    return nil
}

/* Loads the game from the database */
func (game *Game) Load(context appengine.Context, gamekey string) os.Error {
    e := datastore.Get(context, game.GetKey(context, gamekey), game)

    if e != nil {
        context.Errorf("Error loading game: %v", e)
    }

    return e
}
