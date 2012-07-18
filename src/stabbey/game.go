package main

import (
    "log"
    "stabbey/uidgenerator"
)

type Game struct {
    //Players []*Player
    //Boards []*Board
    //Entities []Entity
    LastTick int
    GameRunning bool
    Gamekey string
}

/* Makes a brand new game and saves it */
func NewGame(gamekey string) *Game {
    log.Printf("Starting new game, %v", gamekey)
    g := &Game{}
    g.GameRunning = false
    g.Gamekey = gamekey
    return g
}

func (game *Game) Run() {
    log.Printf("Game %v running", game.Gamekey)
    uidg := uidgenerator.New(uidgenerator.PlayerType())
    log.Printf("Player %v", uidg.NextUid());
    log.Printf("Player %v", uidg.NextUid());
    log.Printf("Player %v", uidg.NextUid());
    log.Printf("Player %v", uidg.NextUid());
}
