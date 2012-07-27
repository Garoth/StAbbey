package order

import (
)

var ACTIONS = map[string] string {
    "MoveRight" : "mr",
    "MoveLeft"  : "ml",
    "MoveUp"    : "mu",
    "MoveDown"  : "md" }

type Action struct {
    actionType string
}

func NewAction(ActionType string) *Action {
    return &Action{ActionType}
}

func (a *Action) ActionType() string {
    return a.actionType
}
