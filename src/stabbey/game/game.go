package game

import (
    "log"
    "stabbey/interfaces"
    "stabbey/util"
)

type Game struct {
    Players []interfaces.Player
    //Boards []*Board
    //Entities []Entity
    LastTick int
    GameRunning bool
    Gamekey string
}

func NewGame(gamekey string) *Game {
    log.Printf("Starting new game, %v", gamekey)
    g := &Game{}
    g.GameRunning = false
    g.Gamekey = gamekey
    return g
}

func (game *Game) AddPlayer(player interfaces.Player) {
    game.Players = append(game.Players, player)
}

func (game *Game) GetPlayer(id int) interfaces.Player {
    for _, player := range game.Players {
        if (player.GetPlayerId() == id) {
            return player
        }
    }
    return nil;
}

/* Generates the gamestate for the given player's perspective */
func (game *Game) Json(playerId int) string {
    for _, player := range game.Players {
        log.Printf("Player: %v", player.GetPlayerId())
    }
    util.Stub("game.Json")
    return "";
}

func (game *Game) Run() {
    log.Printf("Game %v running", game.Gamekey)
}
