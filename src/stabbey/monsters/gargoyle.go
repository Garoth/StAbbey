package monsters

import (
    "strconv"
)

func GargoyleBuilder(me *Monster) {
    me.SetName("Gargoyle " + strconv.Itoa(me.MonsterId))
    me.SetMaxArdour(30)
    me.SetArdour(30)

    me.TickFunction = func(tick int) {
    }
}
