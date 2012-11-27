package serializable

import (
    "encoding/json"
    "log"
)

/* Object used for JSON serialization */
type GameInfo struct {
    Version string
}

func NewGameInfo() *GameInfo {
    me := &GameInfo{}

    me.Version = "pre-alpha"

    return me
}

func (me *GameInfo) Json() string {
    b, e := json.Marshal(me)

    if e != nil {
        log.Fatalf("Error serializing GameInfo: %v", e)
    }

    return string(b)
}
