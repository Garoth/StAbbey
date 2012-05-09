package stabbey

import (
    "fmt"
    "appengine"
    "appengine/datastore"
    "appengine/channel"
)

type Player struct {
    Id, Name string
    BoardId, X, Y, EntityId int
}

func NewPlayer(id string) *Player {
    return &Player{id, "NONAME", 0, 0, 0, -1}
}

/* Returns the database key for the player */
func (p *Player) GetKey(context appengine.Context,
        gamekey string) *datastore.Key {

    return datastore.NewKey(context, "Player" + p.Id, gamekey, 0, nil)
}

/* Save the player to the database */
func (p *Player) Save(context appengine.Context, gamekey string) error {
    _, e := datastore.Put(context, p.GetKey(context, gamekey), p)

    if e != nil {
        context.Errorf("Error saving Player: %v", e)
    }

    return e;
}

/* Load a player from the database */
func (p *Player) Load(context appengine.Context, gamekey string) error {
    e := datastore.Get(context, p.GetKey(context, gamekey), p)

    if e != nil {
        context.Errorf("Error loading Player: %v", e)
    }

    return e
}

/* Returns the key for the JS communication channel */
func (p *Player) getChannelKey(gamekey string) string {
    return p.Id + "/" + gamekey
}

/* Opens the communications channel to the the JS client */
func (p *Player) OpenChannel(context appengine.Context,
        gamekey string) (string, error) {

    fmt.Println("Making channel of:", p.getChannelKey(gamekey))
    tok, e := channel.Create(context, p.getChannelKey(gamekey))

    if e != nil {
        context.Errorf("Error opening channel: %v", e)
    }

    return tok, e
}

/* Send a JSON message to the player */
func (p *Player) SendGamestate(context appengine.Context, gamekey string,
        game *Game) error {

    e := channel.Send(context, p.getChannelKey(gamekey),
        game.JSONGamestate(context, gamekey, p))

    if e != nil {
        context.Errorf("Error sending JSON: %v", e)
    }

    return e
}

/*** Implement Entity ***/
func (p *Player) GetPosition() Position {
    return Position{p.BoardId, p.X, p.Y}
}

func (p *Player) SetPosition(pos Position) {
    p.BoardId = pos.BoardId
    p.X = pos.X
    p.Y = pos.Y
}

func (p *Player) SetName(name string) {
    p.Name = name
}

func (p *Player) GetName() string {
    return p.Name
}

func (p *Player) SetEntityID(id int) {
    p.EntityId = id
}

func (p *Player) GetEntityID() int {
    return p.EntityId
}
/*** End Implement Entity ***/
