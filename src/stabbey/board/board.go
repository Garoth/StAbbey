package board

import (
    "log"
    "stabbey/interfaces"
)

type Room struct {
    StartingPointX, StartingPointY, Left, Right, Top, Bottom int
    ConstrainedLeft, ConstrainedRight, ConstrainedTop, ConstrainedBottom bool
}

type Tile struct {
    LocX, LocY int
}

type Board struct {
    /* Unique game level number (i.e. floor number) */
    Level, Width, Height int
    /* List of rooms */
    RoomList map[int]*Room
    /* List of "doors" -- tiles that override walls */
    DoorList map[int]*Tile
}

/* Creates a brand new board, for level -- 0, 1, 2, etc */
func New(level int) *Board {
    b := &Board{}
    b.Level = level
    b.Width = interfaces.BOARD_WIDTH
    b.Height = interfaces.BOARD_HEIGHT
    b.RoomList = make(map[int]*Room)
    b.DoorList = make(map[int]*Tile)
    NewPiecesGenerator(b).Apply()
    DumpRooms(b)
    return b
}

func (b *Board) GetRender() []string {
    /* And finally, write the rooms to the board */
    layer := []string {"................",
                       "................",
                       "................",
                       "................",
                       "................",
                       "................",
                       "................",
                       "................",
                       "................",
                       "................",
                       "................",
                       "................"}

    for _, room := range b.RoomList {
        /* Draw the top wall */
        spY := room.StartingPointY - room.Top
        spX := room.StartingPointX
        // Not efficient, but I don't care
        for x := spX - room.Left; x <= spX + room.Right; x++ {
            var e error
            // TODO magic character
            if layer, e = setTile(layer, x, spY, "#"); e != nil {
                log.Fatalf("%v", e)
            }
        }

        /* Draw the bottom wall */
        spY = room.StartingPointY + room.Bottom
        spX = room.StartingPointX
        // Not efficient, but I don't care
        for x := spX - room.Left; x <= spX + room.Right; x++ {
            var e error
            // TODO magic character
            if layer, e = setTile(layer, x, spY, "#"); e != nil {
                log.Fatalf("%v", e)
            }
        }

        /* Draw the left wall */
        spY = room.StartingPointY
        spX = room.StartingPointX - room.Left
        // Not efficient, but I don't care
        for y := spY - room.Top; y <= spY + room.Bottom; y++ {
            var e error
            // TODO magic character
            if layer, e = setTile(layer, spX, y, "#"); e != nil {
                log.Fatalf("%v", e)
            }
        }

        /* Draw the right wall */
        spY = room.StartingPointY
        spX = room.StartingPointX + room.Right
        // Not efficient, but I don't care
        for y := spY - room.Top; y <= spY + room.Bottom; y++ {
            var e error
            // TODO magic character
            if layer, e = setTile(layer, spX, y, "#"); e != nil {
                log.Fatalf("%v", e)
            }
        }
    }

    for _, door := range b.DoorList {
        var e error
        if layer, e = setTile(layer, door.LocX, door.LocY, "|"); e != nil {
            log.Fatalln(e)
        }
    }

    return layer
}

func (b *Board) GetLevel() int {
    return b.Level
}

func (b *Board) SetLevel(level int) {
    b.Level = level
}

func (b *Board) GetWidth() int {
    return b.Width
}

func (b *Board) SetWidth(w int) {
    b.Width = w
}

func (b *Board) GetHeight() int {
    return b.Height
}

func (b *Board) SetHeight(h int) {
    b.Height = h
}
