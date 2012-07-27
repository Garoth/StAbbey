package order

import (
    "stabbey/interfaces"
)

type order struct {
    commandcode int
    ticknum int
    actions []interfaces.Action
    player interfaces.Player
}

func NewOrder(commandcode, ticknum int, actions []string,
        player interfaces.Player) *order {

    order := &order{commandcode, ticknum,
        make([]interfaces.Action, len(actions)), player}
    for k, a := range actions {
        order.actions[k] = NewAction(a)
    }

    return order
}

func (o *order) GetCommandCode() int {
    return o.commandcode
}

func (o *order) GetTickNumber() int {
    return o.ticknum
}

func (o *order) GetActions() []interfaces.Action {
    return o.actions
}

func (o *order) GetPlayer() interfaces.Player {
    return o.player
}
