package interfaces

type MonsterModifyGame interface {
    /* Lets the monster drop some loot */
    DropLoot(boardId, x, y int)
}
