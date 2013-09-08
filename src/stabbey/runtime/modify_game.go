package runtime

import (
    "log"
    "time"

    "stabbey/interfaces"
    "stabbey/monsters"
)

/* TODO move *ModifyGame to different files */

/* Holder for functions that Loot may use to modify the game */
type LootModifyGame struct {
}

func NewLootModifyGame() *LootModifyGame {
    me := &LootModifyGame{}
    return me
}

func (me *LootModifyGame) GetPlayerByEntity(
        entity interfaces.Entity) interfaces.Player {
    return GAME.GetPlayerByEntity(entity)
}

/* Holder for functions that Monsters may use to modify the game */
type MonsterModifyGame struct {
}

func NewMonsterModifyGame() *MonsterModifyGame {
    me := &MonsterModifyGame{}
    return me
}

/* Lets an entity drop loot */
func (me *MonsterModifyGame) DropLoot(boardId, x, y int,
        loot interfaces.Loot) {

    loot.SetGameFunctions(NewLootModifyGame())
    loot.SetPosition(boardId, x, y)
    GAME.AddEntity(loot)
    log.Println("loot dropped")
}

/* Initializes game-logic stuff for a particular game level */
func initLevel(levelId int) {
    monsterFunctions := NewMonsterModifyGame();

    /* Spawn some starting monsters */
    for i := 0; i < 3; i++ {
        m := monsters.New(monsters.GargoyleBuilder, monsterFunctions)
        x, y := GAME.GetRandomEmptySpace()
        m.SetPosition(levelId, x, y)
        GAME.AddMonster(m)
    }

    /* TODO Spawn a chest with a skill in */
    c := monsters.New(monsters.ChestBuilder, monsterFunctions)
    x, y := GAME.GetRandomEmptySpace()
    c.SetPosition(levelId, x, y)
    GAME.AddMonster(c)
}

/* Ticks when the players are done their round and other stuff can change */
func worldTick() {
    /* Inform non-player entities */
    for _, ent := range GAME.GetEntities() {
        if (ent.GetType() == interfaces.ENTITY_TYPE_MONSTER) {
            monster := GAME.GetMonsterByEntityId(ent.GetEntityId())
            monster.WorldTick(GAME.GetLastTick())
        }
    }
}

/* Checks whether all players are ready (done queueing) */
func allPlayersReady() bool {
    for _, player := range GAME.GetPlayers() {
        /* Ready players are 1 ahead of the game's tick */
        if player.GetLastTick() <= GAME.GetLastTick() {
            return false
        }
    }

    return true
}

/* Sends the gamestate to everyone */
func broadcastGamestate() {
    for _, player := range GAME.GetPlayers() {
        player.SendMessage(GAME.Json(player))
    }
}

/* Generic action handler for any entity */
func act(entity interfaces.Entity, action interfaces.Action) {
    action.Act(entity, GAME)
}

/* Updates players' ticks and send out gamestate when everyone's ready */
func updateTick(order interfaces.Order) {
    p := order.GetPlayer()
    p.SetLastTick(order.GetTickNumber())
    p.SetLastTickTime(time.Now())
}
