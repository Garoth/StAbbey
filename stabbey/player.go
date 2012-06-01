package stabbey

import (
    "appengine"
    "appengine/channel"
    "appengine/datastore"
    "strconv"
    "time"
)

type Player struct {
    EntityPosition
    Id int
    LastTick int
    LastTickTime time.Time
}

func NewPlayer(c *Context, id int) *Player {
    p := &Player{}
    p.Id = id
    p.Name = "NONAME"
    p.LastTick = -1
    p.LastTickTime = time.Unix(0, 0)
    p.BoardId = -1
    p.Save(c)
    return p
}

/* Creates a player object based on a database player object */
func NewPlayerFromDatabase(dp *DatabasePlayer) *Player {
    p := &Player{}

    p.Id = dp.Id
    p.LastTick = dp.LastTick
    p.LastTickTime = dp.LastTickTime
    p.EntityId = dp.EntityId
    p.BoardId = dp.BoardId
    p.X = dp.X
    p.Y = dp.Y
    p.Name = dp.Name

    return p
}

/* Loads the given player from the database */
func LoadPlayer(c *Context, id int) *Player {
    return NewPlayerFromDatabase(LoadDatabasePlayer(c, id));
}

/* Returns the database key for the player */
func GetPlayerKey(c *Context, id int) *datastore.Key {
    return datastore.NewKey(c.GAEContext, "Player" + strconv.Itoa(id),
        c.Gamekey, 0, nil)
}

/* Save the player to the database */
func (p *Player) Save(c *Context) error {
    dp := NewDatabasePlayer(p)
    return dp.Save(c)
}

/* Checks to see if two players are equal */
func (p *Player) Equals(p2 *Player) bool {
    return p.Id == p2.Id || p.LastTick == p2.LastTick ||
        p.EntityId == p2.EntityId || p.BoardId == p2.BoardId ||
        p.X == p2.X || p.Y == p2.Y || p.Name == p2.Name
}

/* Updates last tick */
func PlayerUpdateLastTick(c *Context, id int, newtick int) {
    datastore.RunInTransaction(c.GAEContext, func(x appengine.Context) error {
        p := LoadPlayer(c, id)
        p.LastTick = newtick
        p.LastTickTime = time.Now()
        p.Save(c)
        return nil // TODO
    }, nil)
}

/******************************************************************************/
/* CHANNEL COMMUNICATIONS FUNCTIONS                                           */
/******************************************************************************/

/* Returns the key for the JS communication channel */
func (p *Player) getChannelKey(c *Context) string {
    return strconv.Itoa(p.Id) + "/" + c.Gamekey
}

/* Opens the communications channel to the the JS client */
func (p *Player) ChannelOpen(c *Context) (string, error) {

    tok, e := channel.Create(c.GAEContext, p.getChannelKey(c))

    if e != nil {
        c.GAEContext.Errorf("Error opening channel: %v", e)
    }

    return tok, e
}

/* Send a string to the player (should be a JSON string) */
func (p *Player) ChannelSend(c *Context, str string) error {
    if e := channel.Send(c.GAEContext, p.getChannelKey(c), str); e != nil {
        c.GAEContext.Errorf("Error sending JSON: %v", e)
        return e
    }
    return nil
}

/* Send the game state to the player */
func (p *Player) ChannelSendGame(c *Context, game *Game) error {
    return p.ChannelSend(c, game.JSON(c, p))
}
