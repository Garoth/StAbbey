/* Static entrance level generator */
package board

import (
    "stabbey/interfaces"
    "stabbey/entity"
)

type centerSprawlGen struct { *Board }

func NewCenterSprawlGenerator(b *Board) interfaces.BoardGenerator {
    return &centerSprawlGen{b}
}

func (me *centerSprawlGen) Apply() {
    /* Player stating location */
    me.PlayerStartX = 6
    me.PlayerStartY = 6

    /* StartX, StartY, Left, Right, Top, Bottom, ... constraints */
    me.RoomList[0] = &Room{0, 0, 1, 3, 1, 5, false, false, false, false}
    me.RoomList[1] = &Room{4, 0, 1, 2, 1, 2, false, false, false, false}
    me.RoomList[2] = &Room{7, 0, 1, 2, 1, 4, false, false, false, false}
    me.RoomList[3] = &Room{10, 2, 1, 4, 1, 2, false, false, false, false}
    me.RoomList[4] = &Room{0, 6, 1, 3, 1, 3, false, false, false, false}
    me.RoomList[5] = &Room{0, 10, 1, 6, 1, 2, false, false, false, false}
    me.RoomList[6] = &Room{6, 5, 1, 4, 1, 2, false, false, false, false}
    me.RoomList[7] = &Room{11, 5, 1, 2, 1, 2, false, false, false, false}
    me.RoomList[8] = &Room{11, 5, 1, 6, 1, 2, false, false, false, false}
    me.RoomList[9] = &Room{7, 8, 1, 2, 1, 4, false, false, false, false}
    me.RoomList[10] = &Room{10, 8, 1, 1, 1, 1, false, false, false, false}
    me.RoomList[11] = &Room{10, 10, 1, 1, 1, 2, false, false, false, false}
    me.RoomList[12] = &Room{12, 8, 1, 4, 1, 4, false, false, false, false}

    /* LocX, LocY */
    me.DoorList[0] = &Tile{6, 0}
    me.DoorList[1] = &Tile{9, 0}
    me.DoorList[2] = &Tile{6, 3}
    me.DoorList[3] = &Tile{9, 3}
    me.DoorList[4] = &Tile{3, 4}
    me.DoorList[5] = &Tile{7, 4}
    me.DoorList[6] = &Tile{8, 4}
    me.DoorList[7] = &Tile{13, 6}
    me.DoorList[8] = &Tile{3, 7}
    me.DoorList[9] = &Tile{7, 7}
    me.DoorList[10] = &Tile{8, 7}
    me.DoorList[11] = &Tile{6, 8}
    me.DoorList[12] = &Tile{9, 8}
    me.DoorList[13] = &Tile{11, 8}
    me.DoorList[14] = &Tile{6, 11}
    me.DoorList[15] = &Tile{9, 11}
    me.DoorList[16] = &Tile{15, 4}

    /* LocX, LocY */
    me.WaterList[len(me.WaterList)] = &Tile{4, 5}
    me.WaterList[len(me.WaterList)] = &Tile{4, 6}

    /* LocX, LocY, Type, Map Char */
    me.GroundDecorList[0]   =  &GroundDecor{6,   5,   "carpet", "c"}
    me.GroundDecorList[1]   =  &GroundDecor{6,   6,   "carpet", "c"}
    me.GroundDecorList[2]   =  &GroundDecor{9,   5,   "carpet", "c"}
    me.GroundDecorList[3]   =  &GroundDecor{9,   6,   "carpet", "c"}
}

func (me *centerSprawlGen) LoadEntities(game interfaces.Game) {
    setEntity(game, entity.NewInertStatue(game), me.Id, 1, 6)
    setEntity(game, entity.NewInertStatue(game), me.Id, 1, 8)
    setEntity(game, entity.NewInertStatue(game), me.Id, 2, 10)
    setEntity(game, entity.NewInertStatue(game), me.Id, 2, 11)
    setEntity(game, entity.NewInertStatue(game), me.Id, 7, 10)
    setEntity(game, entity.NewInertStatue(game), me.Id, 13, 9)

    setEntity(game, entity.NewTeleportTrap(game, 11, 5), me.Id, 10, 10)
    setEntity(game, entity.NewTeleportTrap(game, 4, 10), me.Id, 9, 5)
    setEntity(game, entity.NewTeleportTrap(game, 2, 0), me.Id, 5, 0)

    if (me.Id != game.GetNumBoards() - 1) {
        setEntity(game, entity.NewStairsUp(game), me.Id, 1, 7)
    }
}
