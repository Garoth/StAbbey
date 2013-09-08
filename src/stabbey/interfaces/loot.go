package interfaces

const (
    LOOT_TYPE_ABILITY_PUSH = "pu"
)

type Loot interface {
    Entity
    SetGameFunctions(gameFns LootModifyGame)
}
