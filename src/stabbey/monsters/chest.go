package monsters

import (
    "log"
    "strconv"
    "stabbey/loot"
)

func ChestBuilder(me *Monster) {
    me.SetName("Chest " + strconv.Itoa(me.MonsterId))
    me.SetMaxArdour(10)
    me.SetArdour(10)

    me.TickFunction = func(tick int) {
    }

    me.DeathFunction = func() {
        log.Println(me.GetName(), "drops loot")
        loot := loot.New()
        boardId, x, y := me.GetPosition()
        me.GameFunctions.DropLoot(boardId, x, y, loot)
    }
}
