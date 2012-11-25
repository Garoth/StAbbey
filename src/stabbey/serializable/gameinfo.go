package serializable

import (
    "encoding/json"
    "log"

    "stabbey/order"
)

/* Object used for JSON serialization */
type GameInfo struct {
    Version string
    StartingActions []*Action
}

func NewGameInfo() *GameInfo {
    me := &GameInfo{}

    me.Version = "pre-alpha"
    me.StartingActions = append(me.StartingActions,
        NewAction(order.NewAction(".")))
    me.StartingActions = append(me.StartingActions,
        NewAction(order.NewAction("mu")))

    return me
}

func (me *GameInfo) Json() string {
    b, e := json.Marshal(me)

    if e != nil {
        log.Fatalf("Error serializing GameInfo: %v", e)
    }

    return string(b)
}
