package stabbey

import (
    "appengine/datastore"
    "time"
)

type DatabasePlayer struct {
    /* Player Data */
    Id string
    LastTick int
    LastTickTime time.Time
    /* Entity Data */
    EntityId, BoardId, X, Y int
    Name string
}

func NewDatabasePlayer(p *Player) *DatabasePlayer {
    dp := &DatabasePlayer{}

    dp.Id = p.Id
    dp.LastTick = p.LastTick
    dp.LastTickTime = p.LastTickTime
    dp.EntityId = p.EntityId
    dp.BoardId = p.BoardId
    dp.X = p.X
    dp.Y = p.Y
    dp.Name = p.Name

    return dp
}

func (dp *DatabasePlayer) Save(c *Context) error {
    _, e := datastore.Put(c.GAEContext, GetPlayerKey(c, dp.Id), dp)

    if e != nil {
        c.GAEContext.Errorf("Error saving Player: %v", e)
    } else {
        c.GAEContext.Infof("Successfully saved player %v", dp.Id)
    }

    return e;
}

func LoadDatabasePlayer(c *Context, id string) *DatabasePlayer {
    dp := &DatabasePlayer{}
    e := datastore.Get(c.GAEContext, GetPlayerKey(c, id), dp)

    if e != nil {
        c.GAEContext.Errorf("Error loading Player: %v", e)
    } else {
        //c.GAEContext.Infof("Successfully loaded player %v", id)
    }

    return dp;
}
