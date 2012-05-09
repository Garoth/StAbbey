package stabbey

import (
    "fmt"
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

func LoadPlayer(c *Context, id string) *Player {
    p := &Player{}
    e := datastore.Get(c.GAEContext, GetPlayerKey(c, id), p)

    if e != nil {
        c.GAEContext.Errorf("Error loading Player: %v", e)
    }

    return p
}

/* Returns the database key for the player */
func GetPlayerKey(c *Context, id string) *datastore.Key {
    return datastore.NewKey(c.GAEContext, "Player" + id, c.Gamekey, 0, nil)
}

/* Save the player to the database */
func (p *Player) Save(c *Context) error {
    _, e := datastore.Put(c.GAEContext, GetPlayerKey(c, p.Id), p)

    if e != nil {
        c.GAEContext.Errorf("Error saving Player: %v", e)
    }

    return e;
}

/* Returns the key for the JS communication channel */
func (p *Player) getChannelKey(c *Context) string {
    return p.Id + "/" + c.Gamekey
}

/* Opens the communications channel to the the JS client */
func (p *Player) OpenChannel(c *Context) (string, error) {

    fmt.Println("Making channel of:", p.getChannelKey(c))
    tok, e := channel.Create(c.GAEContext, p.getChannelKey(c))

    if e != nil {
        c.GAEContext.Errorf("Error opening channel: %v", e)
    }

    return tok, e
}

/* Send a JSON message to the player */
func (p *Player) SendGamestate(c *Context, game *Game) error {

    e := channel.Send(c.GAEContext, p.getChannelKey(c),
        game.JSONGamestate(c, p))

    if e != nil {
        c.GAEContext.Errorf("Error sending JSON: %v", e)
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
