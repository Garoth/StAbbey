package order

import (
    "stabbey/interfaces"
)

type order struct {
    commandcode int
    ticknum int
    actions []string
    player interfaces.Player
}

func NewOrder(commandcode, ticknum int, actions []string,
        player interfaces.Player) *order {
    return &order{commandcode, ticknum, actions, player}
}

func (o *order) GetCommandCode() int {
    return o.commandcode
}

func (o *order) GetTickNumber() int {
    return o.ticknum
}

func (o *order) GetActions() []string {
    return o.actions
}

func (o *order) GetPlayer() interfaces.Player {
    return o.player
}
