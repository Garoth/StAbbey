package stabbey

import (
    "appengine/datastore"
)

type DatabaseGame struct {
    Players []int
    Boards []int
    Entities []int
    LastTick int
    GameRunning bool
}

func NewDatabaseGame(g *Game) *DatabaseGame {
    dg := &DatabaseGame{}

    dg.LastTick = g.LastTick
    dg.GameRunning = g.GameRunning
    for _, player := range(g.Players) {
        dg.Players = append(dg.Players, player)
    }
    for _, board := range(g.Boards) {
        dg.Boards = append(dg.Boards, board)
    }
    for _, entity := range(g.Entities) {
        dg.Entities = append(dg.Entities, entity)
    }

    return dg
}

func (dg *DatabaseGame) Save(c *Context) error {
    _, e := datastore.Put(c.GAEContext, GetGameKey(c), dg)

    if e != nil {
        c.GAEContext.Errorf("Error saving Game: %v", e)
    } else {
        c.GAEContext.Infof("Successfully saved game")
    }

    return e;
}

func LoadDatabaseGame(c *Context) *DatabaseGame {
    dg := &DatabaseGame{}
    e := datastore.Get(c.GAEContext, GetGameKey(c), dg)

    if e != nil {
        c.GAEContext.Errorf("Error loading Game: %v", e)
    } else {
    }

    return dg;
}
