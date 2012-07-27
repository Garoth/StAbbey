package interfaces

type Order interface {
    GetCommandCode() int
    GetTickNumber() int
    GetActions() []Action
    GetPlayer() Player
}
