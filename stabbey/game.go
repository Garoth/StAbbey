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

func (game *Game) AddPlayer(p *Player) {
    game.Players = append(game.Players, p.Id)
}

func (game *Game) GetKey(context appengine.Context, gamekey string) *datastore.Key {
    return datastore.NewKey(context, "Game", gamekey, 0, nil)
}

func (game *Game) Save(context appengine.Context, gamekey string) os.Error {
    if _, e := datastore.Put(context, game.GetKey(context, gamekey), game); true {
        return e
    }

    for _, Id := range game.Players {
        p := NewPlayer(Id)
        p.Save(context, gamekey)
    }

    return nil
}

func (game *Game) Load(context appengine.Context, gamekey string) os.Error {
    return datastore.Get(context, game.GetKey(context, gamekey), game)
}
