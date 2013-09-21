package board

import (
    "stabbey/interfaces"
    "stabbey/entity"
)

type piecesGen struct { *Board }

/* Idea of this board generator is to use pre-made pieces and combine them
 * in a functional but random way. */
func NewPiecesGenerator(b *Board) interfaces.BoardGenerator {
    return &piecesGen{b}
}

func setEntity(game interfaces.Game, entity interfaces.Entity, x, y int) {
    entity.SetPosition(game.GetCurrentBoard(), x, y)
    game.AddEntity(entity)
}

func (me *piecesGen) Apply() {
    /* Player stating location */
    me.PlayerStartX = 2
    me.PlayerStartY = 11

    /* StartX, StartY, Left, Right, Top, Bottom, ... constraints */
    me.RoomList[0] = &Room{0, 0, 1, 6, 1, 4, false, false, false, false}
    me.RoomList[1] = &Room{7, 0, 1, 6, 1, 4, false, false, false, false}
    me.RoomList[2] = &Room{14, 0, 1, 2, 1, 4, false, false, false, false}
    me.RoomList[3] = &Room{0, 5, 1, 6, 1, 7, false, false, false, false}
    me.RoomList[4] = &Room{7, 5, 1, 9, 1, 7, false, false, false, false}

    me.RoomList[5] = &Room{7, 5, 0, 0, 0, 0, false, false, false, false}
    me.RoomList[6] = &Room{15, 5, 0, 0, 0, 0, false, false, false, false}
    me.RoomList[7] = &Room{7, 11, 0, 0, 0, 0, false, false, false, false}
    me.RoomList[8] = &Room{15, 11, 0, 0, 0, 0, false, false, false, false}

    me.RoomList[9] = &Room{9, 6, 0, 0, 0, 0, false, false, false, false}
    me.RoomList[10] = &Room{13, 6, 0, 0, 0, 0, false, false, false, false}
    me.RoomList[11] = &Room{9, 10, 0, 0, 0, 0, false, false, false, false}
    me.RoomList[12] = &Room{13, 10, 0, 0, 0, 0, false, false, false, false}

    /* LocX, LocY */
    me.DoorList[0] = &Tile{2, 4}
    me.DoorList[1] = &Tile{3, 4}
    me.DoorList[2] = &Tile{6, 1}
    me.DoorList[3] = &Tile{6, 2}
    me.DoorList[4] = &Tile{13, 0}
    me.DoorList[5] = &Tile{10, 4}
    me.DoorList[6] = &Tile{11, 4}
    me.DoorList[7] = &Tile{12, 4}

    /* LocX, LocY */
    me.WaterList[0] = &Tile{0, 8}
    me.WaterList[1] = &Tile{1, 8}
    me.WaterList[2] = &Tile{1, 9}
    me.WaterList[3] = &Tile{1, 10}
    me.WaterList[4] = &Tile{1, 11}

    me.WaterList[5] = &Tile{5, 8}
    me.WaterList[6] = &Tile{4, 8}
    me.WaterList[7] = &Tile{4, 9}
    me.WaterList[8] = &Tile{4, 10}
    me.WaterList[9] = &Tile{4, 11}

    me.WaterList[12] = &Tile{11, 3}
    me.WaterList[13] = &Tile{11, 4}
    me.WaterList[14] = &Tile{11, 5}
    me.WaterList[15] = &Tile{11, 6}
    me.WaterList[16] = &Tile{11, 7}
    me.WaterList[17] = &Tile{9, 8}
    me.WaterList[18] = &Tile{10, 8}
    me.WaterList[19] = &Tile{11, 8}
    me.WaterList[20] = &Tile{12, 8}
    me.WaterList[21] = &Tile{13, 8}
    me.WaterList[22] = &Tile{11, 9}
    me.WaterList[23] = &Tile{11, 10}

    /* LocX, LocY, Type, Map Char */
    me.GroundDecorList[0]   =  &GroundDecor{2,   1,   "carpet", "c"}
    me.GroundDecorList[1]   =  &GroundDecor{3,   1,   "carpet", "c"}
    me.GroundDecorList[2]   =  &GroundDecor{4,   1,   "carpet", "c"}
    me.GroundDecorList[3]   =  &GroundDecor{5,   1,   "carpet", "c"}
    me.GroundDecorList[4]   =  &GroundDecor{6,   1,   "carpet", "c"}
    me.GroundDecorList[5]   =  &GroundDecor{7,   1,   "carpet", "c"}
    me.GroundDecorList[6]   =  &GroundDecor{8,   1,   "carpet", "c"}
    me.GroundDecorList[7]   =  &GroundDecor{9,   1,   "carpet", "c"}
    me.GroundDecorList[8]   =  &GroundDecor{10,  1,   "carpet", "c"}
    me.GroundDecorList[9]   =  &GroundDecor{11,  1,   "carpet", "c"}
    me.GroundDecorList[10]  =  &GroundDecor{12,  1,   "carpet", "c"}

    me.GroundDecorList[11]   = &GroundDecor{2,   2,   "carpet", "c"}
    me.GroundDecorList[12]   = &GroundDecor{3,   2,   "carpet", "c"}
    me.GroundDecorList[13]   = &GroundDecor{4,   2,   "carpet", "c"}
    me.GroundDecorList[14]   = &GroundDecor{5,   2,   "carpet", "c"}
    me.GroundDecorList[15]   = &GroundDecor{6,   2,   "carpet", "c"}
    me.GroundDecorList[16]   = &GroundDecor{7,   2,   "carpet", "c"}
    me.GroundDecorList[17]   = &GroundDecor{8,   2,   "carpet", "c"}
    me.GroundDecorList[18]   = &GroundDecor{9,   2,   "carpet", "c"}
    me.GroundDecorList[19]   = &GroundDecor{10,  2,   "carpet", "c"}
    me.GroundDecorList[20]   = &GroundDecor{11,  2,   "carpet", "c"}
    me.GroundDecorList[21]   = &GroundDecor{12,  2,   "carpet", "c"}

    me.GroundDecorList[22]   = &GroundDecor{2,   3,   "carpet", "c"}
    me.GroundDecorList[23]   = &GroundDecor{3,   3,   "carpet", "c"}
    me.GroundDecorList[24]   = &GroundDecor{10,  3,   "carpet", "c"}
    me.GroundDecorList[25]   = &GroundDecor{12,  3,   "carpet", "c"}

    me.GroundDecorList[26]   = &GroundDecor{2,   4,   "carpet", "c"}
    me.GroundDecorList[27]   = &GroundDecor{3,   4,   "carpet", "c"}
    me.GroundDecorList[28]   = &GroundDecor{10,  4,   "carpet", "c"}
    me.GroundDecorList[29]   = &GroundDecor{12,  4,   "carpet", "c"}

    me.GroundDecorList[30]   = &GroundDecor{10,  5,   "carpet", "c"}
    me.GroundDecorList[31]   = &GroundDecor{12,  5,   "carpet", "c"}

    me.GroundDecorList[32]   = &GroundDecor{10,  5,   "carpet", "c"}
    me.GroundDecorList[33]   = &GroundDecor{12,  5,   "carpet", "c"}

    me.GroundDecorList[34]   = &GroundDecor{10,  6,   "carpet", "c"}
    me.GroundDecorList[35]   = &GroundDecor{12,  6,   "carpet", "c"}

    me.GroundDecorList[36]   = &GroundDecor{8,   7,   "carpet", "c"}
    me.GroundDecorList[37]   = &GroundDecor{9,   7,   "carpet", "c"}
    me.GroundDecorList[38]   = &GroundDecor{10,  7,   "carpet", "c"}
    me.GroundDecorList[39]   = &GroundDecor{12,  7,   "carpet", "c"}
    me.GroundDecorList[40]   = &GroundDecor{13,  7,   "carpet", "c"}
    me.GroundDecorList[41]   = &GroundDecor{14,  7,   "carpet", "c"}
    me.GroundDecorList[42]   = &GroundDecor{15,  7,   "carpet", "c"}

    me.GroundDecorList[43]   = &GroundDecor{8,   8,   "carpet", "c"}
    me.GroundDecorList[44]   = &GroundDecor{14,  8,   "carpet", "c"}
    me.GroundDecorList[45]   = &GroundDecor{15,  8,   "carpet", "c"}

    me.GroundDecorList[46]   = &GroundDecor{8,   9,   "carpet", "c"}
    me.GroundDecorList[47]   = &GroundDecor{9,   9,   "carpet", "c"}
    me.GroundDecorList[48]   = &GroundDecor{10,  9,   "carpet", "c"}
    me.GroundDecorList[49]   = &GroundDecor{12,  9,   "carpet", "c"}
    me.GroundDecorList[50]   = &GroundDecor{13,  9,   "carpet", "c"}
    me.GroundDecorList[51]   = &GroundDecor{14,  9,   "carpet", "c"}
    me.GroundDecorList[52]   = &GroundDecor{15,  9,   "carpet", "c"}

    me.GroundDecorList[53]   = &GroundDecor{10,  10,  "carpet", "c"}
    me.GroundDecorList[54]   = &GroundDecor{12,  10,  "carpet", "c"}

    me.GroundDecorList[55]   = &GroundDecor{10,  11,  "carpet", "c"}
    me.GroundDecorList[56]   = &GroundDecor{11,  11,  "carpet", "c"}
    me.GroundDecorList[57]   = &GroundDecor{12,  11,  "carpet", "c"}

    me.GroundDecorList[58] = &GroundDecor{0,  5,   "grass",  "g"}
    me.GroundDecorList[59] = &GroundDecor{1,  5,   "grass",  "g"}
    me.GroundDecorList[60] = &GroundDecor{2,  5,   "grass",  "g"}
    me.GroundDecorList[61] = &GroundDecor{3,  5,   "flower", "f"}
    me.GroundDecorList[62] = &GroundDecor{4,  5,   "grass",  "g"}
    me.GroundDecorList[63] = &GroundDecor{5,  5,   "grass",  "g"}
    me.GroundDecorList[64] = &GroundDecor{0,  6,   "grass",  "g"}
    me.GroundDecorList[65] = &GroundDecor{1,  6,   "grass",  "g"}
    me.GroundDecorList[66] = &GroundDecor{2,  6,   "grass",  "g"}
    me.GroundDecorList[67] = &GroundDecor{3,  6,   "grass",  "g"}
    me.GroundDecorList[68] = &GroundDecor{4,  6,   "grass",  "g"}
    me.GroundDecorList[69] = &GroundDecor{5,  6,   "grass",  "g"}
    me.GroundDecorList[70] = &GroundDecor{0,  7,   "grass",  "g"}
    me.GroundDecorList[71] = &GroundDecor{1,  7,   "grass",  "g"}
    me.GroundDecorList[72] = &GroundDecor{2,  7,   "flower", "f"}
    me.GroundDecorList[73] = &GroundDecor{3,  7,   "grass",  "g"}
    me.GroundDecorList[74] = &GroundDecor{4,  7,   "grass",  "g"}
    me.GroundDecorList[75] = &GroundDecor{5,  7,   "grass",  "g"}
    me.GroundDecorList[76] = &GroundDecor{2,  8,   "grass",  "g"}
    me.GroundDecorList[77] = &GroundDecor{3,  8,   "flower", "f"}
    me.GroundDecorList[78] = &GroundDecor{0,  9,   "grass",  "g"}
    me.GroundDecorList[79] = &GroundDecor{2,  9,   "grass",  "g"}
    me.GroundDecorList[80] = &GroundDecor{3,  9,   "grass",  "g"}
    me.GroundDecorList[81] = &GroundDecor{5,  9,   "grass",  "g"}
    me.GroundDecorList[82] = &GroundDecor{0,  10,  "grass",  "g"}
    me.GroundDecorList[83] = &GroundDecor{2,  10,  "grass",  "f"}
    me.GroundDecorList[84] = &GroundDecor{3,  10,  "grass",  "g"}
    me.GroundDecorList[85] = &GroundDecor{5,  10,  "grass",  "g"}
    me.GroundDecorList[86] = &GroundDecor{0,  11,  "grass",  "g"}
    me.GroundDecorList[87] = &GroundDecor{2,  11,  "flower", "f"}
    me.GroundDecorList[88] = &GroundDecor{3,  11,  "grass",  "g"}
    me.GroundDecorList[89] = &GroundDecor{5,  11,  "grass",  "g"}
}

func (me *piecesGen) LoadEntities(game interfaces.Game) {
    setEntity(game, entity.NewInertStatue(game), 1, 5)
    setEntity(game, entity.NewInertStatue(game), 4, 5)
    setEntity(game, entity.NewInertStatue(game), 8, 5)
    setEntity(game, entity.NewInertStatue(game), 14, 5)
    setEntity(game, entity.NewInertStatue(game), 9, 11)
    setEntity(game, entity.NewInertStatue(game), 13, 11)

    setEntity(game, entity.NewChest(game), 15, 3)

    setEntity(game, entity.NewTree(game), 0, 5)
    setEntity(game, entity.NewTree(game), 5, 5)
    setEntity(game, entity.NewTree(game), 0, 7)
    setEntity(game, entity.NewTree(game), 5, 7)
    setEntity(game, entity.NewTree(game), 0, 9)
    setEntity(game, entity.NewTree(game), 5, 9)
    setEntity(game, entity.NewTree(game), 0, 11)
    setEntity(game, entity.NewTree(game), 5, 11)
}
