package monsters

import (
    "log"
    "strconv"
)

func GargoyleBuilder(me *Monster) {
    me.SetName("Gargoyle " + strconv.Itoa(me.MonsterId))
    me.TickFunction = func(tick int) {
        log.Println(me.GetName(), "says,", tick)
    }
}
