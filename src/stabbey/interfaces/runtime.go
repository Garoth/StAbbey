package interfaces

type MonsterModifyGame interface {
    /* Lets the monster drop some loot */
    DropLoot(boardId, x, y int, loot Loot)
}

type LootModifyGame interface {
    /* Lets the loot look up a player by entity */
    GetPlayerByEntity(entity Entity) Player
}
