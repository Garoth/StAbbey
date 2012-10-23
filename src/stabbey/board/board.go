package board

import (
    "log"
    "math/rand"
    "time"
    "stabbey/interfaces"
)

type Room struct {
    StartingPointX, StartingPointY, Left, Right, Top, Bottom int
    ConstrainedLeft, ConstrainedRight, ConstrainedTop, ConstrainedBottom bool
}

type Board struct {
    /* Unique game level number (i.e. floor number) */
    Level, Width, Height int
    /* List of rooms */
    RoomList map[int]*Room
}

/* Creates a brand new board, for level -- 0, 1, 2, etc */
func New(level int) *Board {
    b := &Board{}
    b.Level = level
    b.Width = interfaces.BOARD_WIDTH
    b.Height = interfaces.BOARD_HEIGHT
    b.RoomList = make(map[int]*Room)
    NewGrowingGenerator(b).Apply()
    DumpRooms(b)
    rand.Seed(time.Now().Unix())
    return b
}

/* Picks a random spawn point. TODO: should be in Game so that things can't
 * spawn over entities */
func (b *Board) GetRandomSpawnPoint() (int, int) {
    maxAttempts := 1000

    for x := 0; x < maxAttempts; x++ {
        x := rand.Intn(interfaces.BOARD_WIDTH)
        y := rand.Intn(interfaces.BOARD_HEIGHT)

        // TODO Bad way to do this
        if b.GetRender()[y][x] == '.' {
            return x, y
        }
    }

    /* TODO should never happen */
    return 0, 0
}

func (b *Board) GetRender() []string {
    /* And finally, write the rooms to the board */
    layer := []string {"L--------------L",
                       "|..............|",
                       "|..............|",
                       "|..............|",
                       "|..............|",
                       "|..............|",
                       "|..............|",
                       "|..............|",
                       "|..............|",
                       "|..............|",
                       "|..............|",
                       "L--------------L"}

    for _, room := range b.RoomList {
        /* Draw the top wall */
        spY := room.StartingPointY - room.Top
        spX := room.StartingPointX
        // Not efficient, but I don't care
        for x := spX - room.Left; x <= spX + room.Right; x++ {
            var e error
            // TODO magic character
            if layer, e = setTile(layer, x, spY, "-"); e != nil {
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
            if layer, e = setTile(layer, x, spY, "-"); e != nil {
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
            if layer, e = setTile(layer, spX, y, "|"); e != nil {
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
            if layer, e = setTile(layer, spX, y, "|"); e != nil {
                log.Fatalf("%v", e)
            }
        }

        /* Draw the L corners */
        layer, _ = setTile(layer, room.StartingPointX - room.Left,
            room.StartingPointY - room.Top, "L")
        layer, _ = setTile(layer, room.StartingPointX + room.Right,
            room.StartingPointY - room.Top, "L")
        layer, _ = setTile(layer, room.StartingPointX - room.Left,
            room.StartingPointY + room.Bottom, "L")
        layer, _ = setTile(layer, room.StartingPointX + room.Right,
            room.StartingPointY + room.Bottom, "L")
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
