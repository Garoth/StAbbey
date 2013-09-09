package board

import (
    "stabbey/interfaces"
)

type piecesGen struct { *Board }

/* Idea of this board generator is to use pre-made pieces and combine them
 * in a functional but random way. */
func NewPiecesGenerator(b *Board) interfaces.BoardGenerator {
    return &piecesGen{b}
}

func (me *piecesGen) Apply() {
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
    me.RoomList[11] = &Room{9, 9, 0, 0, 0, 0, false, false, false, false}
    me.RoomList[12] = &Room{13, 9, 0, 0, 0, 0, false, false, false, false}

    /* LocX, LocY */
    me.DoorList[0] = &Tile{2, 4}
    me.DoorList[1] = &Tile{3, 4}
    me.DoorList[2] = &Tile{6, 1}
    me.DoorList[3] = &Tile{6, 2}
    me.DoorList[4] = &Tile{13, 0}
    me.DoorList[5] = &Tile{10, 4}
    me.DoorList[6] = &Tile{11, 4}
    me.DoorList[7] = &Tile{12, 4}
}
