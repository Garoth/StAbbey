package monsters

import (
    "strconv"
)

func GargoyleBuilder(me *Monster) {
    me.SetName("Gargoyle " + strconv.Itoa(me.MonsterId))
    me.TickFunction = func(tick int) {
    }
}
