package interfaces

type Order interface {
    GetCommandCode() int
    GetTickNumber() int
    GetActions() []string
    GetPlayer() Player
}
