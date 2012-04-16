package stabbey

import (
    "os"
    "appengine"
    "appengine/datastore"
)

type Player struct {
    Id string
}

func NewPlayer(id string) *Player {
    return &Player{id}
}

func (p *Player) GetKey(context appengine.Context, gamekey string) *datastore.Key {
    return datastore.NewKey(context, "Player" + p.Id, gamekey, 0, nil)
}

func (p *Player) Save(context appengine.Context, gamekey string) os.Error {
    if _, e := datastore.Put(context, p.GetKey(context, gamekey), p); true {
        return e
    }
    return nil;
}

func (p *Player) Load(context appengine.Context, gamekey string) os.Error {
    return datastore.Get(context, p.GetKey(context, gamekey), p)
}
