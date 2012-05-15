package stabbey

import (
    "appengine"
    "appengine/datastore"
    "encoding/json"
)

/* Internal, GAE struct for database storing */
type Game struct {
    Players []string
    Boards []int
    LastTick int
    GameRunning bool
}

/* Object used for JSON serialization */
type jsonGame struct {
    Players []*Player
    Boards []*Board
    LastTick int
}

func NewGame() *Game {
    g := &Game{}
    g.LastTick = 0
    g.GameRunning = false
    return g
}

/* Loads the game from the database */
func LoadGame(c *Context) *Game {
    g := &Game{}
    e := datastore.Get(c.GAEContext, GetGameKey(c), g)

    if e != nil {
        c.GAEContext.Errorf("Error loading game: %v", e)
    }

    return g
}

/* Gets the database key for the game */
func GetGameKey(c *Context) *datastore.Key {
    return datastore.NewKey(c.GAEContext, "Game", c.Gamekey, 0, nil)
}

/* Adds a player to the current game. Call once per player per game */
func (game *Game) AddPlayer(p *Player) {
    game.Players = append(game.Players, p.Id)
}

/* Adds a board to the current game. Call once per board per game */
func (game *Game) AddBoard(b *Board) {
    game.Boards = append(game.Boards, b.Id)
}

/* Saves the game to the database */
func (game *Game) Save(c *Context) error {
    _, e := datastore.Put(c.GAEContext, GetGameKey(c), game)

    if e != nil {
        c.GAEContext.Errorf("Error saving Game: %v", e)
        return e
    }

    return nil
}

/* Gets the JSON gamestate for the given player's perspective */
func (game *Game) JSONGamestate(c *Context, p *Player) string {
    jg := jsonGame{};

    for _, ID := range game.Players {
        jg.Players = append(jg.Players, LoadPlayer(c, ID))
    }

    for _, ID := range game.Boards {
        jg.Boards = append(jg.Boards, LoadBoard(c, string(ID)))
    }

    jg.LastTick = game.LastTick

    b, _ := json.Marshal(jg)

    return string(b)
}

func GameUpdateLastTick(c *Context) {
    datastore.RunInTransaction(c.GAEContext, func(x appengine.Context) error {
        g := LoadGame(c)
        g.LastTick += 1
        g.Save(c)
        return nil // TODO
    }, nil)
}

func GameSetRunning(c *Context, running bool) {
    datastore.RunInTransaction(c.GAEContext, func(x appengine.Context) error {
        g := LoadGame(c)
        g.GameRunning = running;
        g.Save(c)
        return nil // TODO
    }, nil)
}
