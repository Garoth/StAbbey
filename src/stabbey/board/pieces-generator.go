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
    me.RoomList[0] = &Room{0, 0, 0, 5, 0, 3, false, false, false, false}
    me.RoomList[1] = &Room{3, 11, 0, 4, 5, 0, false, false, false, false}
    me.RoomList[2] = &Room{0, 11, 0, 3, 4, 0, false, false, false, false}
    me.RoomList[3] = &Room{8, 0, 0, 0, 0, 4, false, false, false, false}
    me.RoomList[4] = &Room{15, 0, 6, 0, 0, 5, false, false, false, false}
    me.RoomList[5] = &Room{15, 11, 4, 0, 6, 0, false, false, false, false}
    me.RoomList[6] = &Room{15, 4, 3, 0, 0, 0, false, false, false, false}
    me.RoomList[7] = &Room{8, 11, 0, 0, 2, 0, false, false, false, false}
    me.RoomList[8] = &Room{15, 10, 2, 0, 0, 0, false, false, false, false}

    /* LocX, LocY */
    me.DoorList[0] = &Tile{5, 2}
    me.DoorList[1] = &Tile{10, 5}
    me.DoorList[2] = &Tile{11, 10}
    me.DoorList[3] = &Tile{7, 8}
    me.DoorList[4] = &Tile{3, 10}
    me.DoorList[5] = &Tile{1, 7}
}
