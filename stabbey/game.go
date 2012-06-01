package stabbey

import (
    "appengine"
    "appengine/datastore"
    "encoding/json"
)

/* Internal, GAE struct for database storing */
type Game struct {
    Players []int
    Boards []int
    LastTick int
    GameRunning bool
}

/* Makes a brand new game and saves it */
func NewGame(c *Context) *Game {
    c.GAEContext.Infof("Making new game, %v", c.Gamekey)
    g := &Game{}
    g.LastTick = 0
    g.GameRunning = false
    g.Save(c);
    return g
}

/* Loads a game object from the database */
func NewGameFromDatabase(dg *DatabaseGame) *Game {
    g := &Game{}

    g.LastTick = dg.LastTick
    g.GameRunning = dg.GameRunning
    for _, player := range(dg.Players) {
        g.Players = append(g.Players, player)
    }
    for _, board := range(dg.Boards) {
        g.Boards = append(g.Boards, board)
    }

    return g
}

/* Gets the database key for the game */
func GetGameKey(c *Context) *datastore.Key {
    return datastore.NewKey(c.GAEContext, "Game", c.Gamekey, 0, nil)
}

/* Saves the game to the database */
func (game *Game) Save(c *Context) error {
    return NewDatabaseGame(game).Save(c)
}

/* Loads the game from the database */
func LoadGame(c *Context) *Game {
    return NewGameFromDatabase(LoadDatabaseGame(c))
}

/* Adds a player to the current game. Call once per player per game */
func GameAddPlayer(c *Context, p *Player) {
    datastore.RunInTransaction(c.GAEContext, func(x appengine.Context) error {
        g := LoadGame(c)
        g.Players = append(g.Players, p.Id)
        g.Save(c)
        return nil // TODO
    }, nil)
}

/* Adds a board to the current game. Call once per board per game */
func GameAddBoard(c *Context, b *Board) {
    datastore.RunInTransaction(c.GAEContext, func(x appengine.Context) error {
        g := LoadGame(c)
        g.Boards = append(g.Boards, b.Id)
        g.Save(c)
        return nil // TODO
    }, nil)
}

/* Atomically sets the last tick of the game */
func GameUpdateLastTick(c *Context) {
    datastore.RunInTransaction(c.GAEContext, func(x appengine.Context) error {
        g := LoadGame(c)
        g.LastTick += 1
        g.Save(c)
        return nil // TODO
    }, nil)
}

/* Atomically sets the running state of the game */
func GameSetRunning(c *Context, running bool) {
    datastore.RunInTransaction(c.GAEContext, func(x appengine.Context) error {
        g := LoadGame(c)
        g.GameRunning = running;
        g.Save(c)
        return nil // TODO
    }, nil)
}

/* Gets the JSON gamestate for the given player's perspective */
func (game *Game) JSON(c *Context, p *Player) string {
    sg := NewSerializableGame(c, game)
    b, _ := json.Marshal(sg)
    return string(b)
}
