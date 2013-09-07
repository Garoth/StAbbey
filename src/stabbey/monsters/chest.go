package monsters

import (
    "strconv"
)

func ChestBuilder(me *Monster) {
    me.SetName("Chest " + strconv.Itoa(me.MonsterId))
    me.SetMaxArdour(10)
    me.SetArdour(10)

    me.TickFunction = func(tick int) {
    }
}
