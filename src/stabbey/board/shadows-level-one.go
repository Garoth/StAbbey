/* Static entrance level generator */
package board

import (
    "stabbey/interfaces"
    "stabbey/entity"
)

type shadowsOneGen struct { *Board }

func NewShadowsLevelOneGenerator (b *Board) interfaces.BoardGenerator {
    return &shadowsOneGen{b}
}

func (me *shadowsOneGen) Apply() {
    /* Player stating location */
    me.PlayerStartX = 7
    me.PlayerStartY = 11

    /* x, y, width, height */
    /* left side */
    me.simpleAddRoom(0, 0, 4, 4)
    me.simpleAddRoom(0, 5, 4, 3)
    me.simpleAddRoom(0, 9, 4, 3)
    /* right side */
    me.simpleAddRoom(11, 0, 5, 2)
    me.simpleAddRoom(11, 3, 5, 3)
    me.simpleAddRoom(11, 9, 5, 3)

    /* x, y */
    /* left side doors */
    me.addDoor(4, 3)
    me.addDoor(4, 7)
    me.addDoor(4, 11)
    /* right side doors */
    me.addDoor(10, 1)
    me.addDoor(10, 11)
    me.addDoor(15, 6)

    /* x, y */
    me.addWater(11, 3)
    me.addWater(11, 4)
    me.addWater(12, 5)
    me.addWater(13, 3)
    me.addWater(13, 4)
    me.addWater(14, 3)
    me.addWater(15, 4)

    /* x, y, decoration name */
    for i := 0; i < 12; i++ {
        me.addDecoration(6, i, "carpet")
        me.addDecoration(8, i, "carpet")
    }
    me.addDecoration(7, 11, "carpet")
}

func (me *shadowsOneGen) LoadEntities(game interfaces.Game) {
    /* statue army */
    setEntity(game, entity.NewInertStatue(game), me.Id, 5, 0)
    setEntity(game, entity.NewInertStatue(game), me.Id, 9, 0)
    setEntity(game, entity.NewInertStatue(game), me.Id, 5, 2)
    setEntity(game, entity.NewInertStatue(game), me.Id, 7, 2)
    setEntity(game, entity.NewInertStatue(game), me.Id, 9, 2)
    setEntity(game, entity.NewInertStatue(game), me.Id, 5, 4)
    setEntity(game, entity.NewInertStatue(game), me.Id, 7, 4)
    setEntity(game, entity.NewInertStatue(game), me.Id, 9, 4)
    setEntity(game, entity.NewInertStatue(game), me.Id, 5, 6)
    setEntity(game, entity.NewInertStatue(game), me.Id, 7, 6)
    setEntity(game, entity.NewInertStatue(game), me.Id, 9, 6)
    setEntity(game, entity.NewInertStatue(game), me.Id, 5, 8)
    setEntity(game, entity.NewInertStatue(game), me.Id, 7, 8)
    setEntity(game, entity.NewInertStatue(game), me.Id, 9, 8)

    /* random card chest / burning room */
    setEntity(game, entity.NewChest(game), me.Id, 12, 3)

    /* antivenom necklace chest with shades */
    setEntity(game, entity.NewChest(game), me.Id, 1, 6)
    setEntity(game, entity.NewGargoyle(game), me.Id, 2, 6)
    setEntity(game, entity.NewGargoyle(game), me.Id, 2, 7)

    /* troll chest with teleport traps */
    setEntity(game, entity.NewChest(game), me.Id, 14, 10)
    setEntity(game, entity.NewTeleportTrap(game, 1, 2), me.Id, 15, 11)
    setEntity(game, entity.NewTeleportTrap(game, 2, 2), me.Id, 14, 11)
    setEntity(game, entity.NewTeleportTrap(game, 3, 2), me.Id, 13, 10)
    setEntity(game, entity.NewTeleportTrap(game, 1, 1), me.Id, 13, 9)
    setEntity(game, entity.NewTeleportTrap(game, 2, 1), me.Id, 12, 10)
    setEntity(game, entity.NewTeleportTrap(game, 3, 1), me.Id, 11, 9)
    setEntity(game, entity.NewTeleportTrap(game, 3, 1), me.Id, 10, 11)

    if (me.Id != game.GetNumBoards() - 1) {
        setEntity(game, entity.NewStairsUp(game), me.Id, 7, 0)
    }
}
