package board

import (
    "log"
    "math/rand"
    "time"

    "stabbey/interfaces"
)

/* How many rooms can spawn on a wall (will be 0 to max_wal_rooms-1) */
const max_wall_rooms int = 3
const max_rooms int = (max_wall_rooms - 1) * 4 + 1

type room struct {
    startingPointX, startingPointY, left, right, top, bottom int
    constrainedLeft, constrainedRight, constrainedTop, constainedBottom bool
}

type roomlist struct {
    list map[int]*room
}

func NewGrowingGenerator() interfaces.BoardGenerator {
    rl := &roomlist{make(map[int]*room)}
    return rl
}

func (rl *roomlist) Apply(b interfaces.Board) {
    rand.Seed(time.Now().Unix())
    totalRooms := 0

    /* Generate some room seeds at the edges of the board
     * 0 = top
     * 1 = right
     * 2 = bottom
     * 3 = right
     */
    for wall := 0; wall < 4; wall++ {
        var x, y int

        if wall == 0 {
            x = -1
            y = 0
        } else if wall == 1 {
            x = interfaces.BOARD_WIDTH - 1
            y = -1
        } else if wall == 2 {
            x = -1
            y = interfaces.BOARD_HEIGHT - 1
        } else if wall == 3 {
            x = 0
            y = -1
        }

        for nRooms := rand.Intn(max_wall_rooms) - 1; nRooms >= 0; nRooms-- {
            for {
                r := &room{}

                if x != -1 {
                    r.startingPointX = x
                    if x == 0 {
                        r.constrainedLeft = true
                    } else {
                        r.constrainedRight = true
                    }
                } else {
                    r.startingPointX = rand.Intn(interfaces.BOARD_WIDTH)
                }

                if y != -1 {
                    r.startingPointY = y
                    if y == 0 {
                        r.constrainedTop = true
                    } else {
                        r.constainedBottom = true
                    }
                } else {
                    r.startingPointY = rand.Intn(interfaces.BOARD_HEIGHT)
                }

                if rl.AlreadyRoomThere(r) == false {
                    rl.list[totalRooms] = r
                    /* expand the room by 1 in every direciton right away */
                    for i := 0; i < 4; i++ {
                        rl.TryGrowRoom(totalRooms, i, 1)
                    }
                    totalRooms++
                    break
                }
            }
        }
    }

    /* Now generate some extra rooms in random places */
    for nRooms := rand.Intn(max_rooms-totalRooms + 1); nRooms >= 0; nRooms-- {
        for {
            r := &room{}
            r.startingPointX = rand.Intn(interfaces.BOARD_WIDTH)
            r.startingPointY = rand.Intn(interfaces.BOARD_HEIGHT)

            if rl.AlreadyRoomThere(r) == false {
                rl.list[totalRooms] = r
                /* expand the room by 1 in every direciton right away */
                for i := 0; i < 4; i++ {
                    rl.TryGrowRoom(totalRooms, i, 1)
                }
                totalRooms++
                break
            }
        }
    }

    /* Randomly grow the rooms a bit, define walls */
    for i := 0; i < 5; i++ {
        for k, _ := range rl.list {
            rl.TryGrowRoom(k, rand.Intn(4), 1)
        }
    }

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

    for _, room := range rl.list {
        /* Draw the top wall */
        spY := room.startingPointY - room.top
        spX := room.startingPointX
        // Not efficient, but I don't care
        for x := spX - room.left; x <= spX + room.right; x++ {
            var e error
            // TODO magic character
            if layer, e = setTile(layer, x, spY, "-"); e != nil {
                log.Fatalf("%v", e)
            }
        }

        /* Draw the bottom wall */
        spY = room.startingPointY + room.bottom
        spX = room.startingPointX
        // Not efficient, but I don't care
        for x := spX - room.left; x <= spX + room.right; x++ {
            var e error
            // TODO magic character
            if layer, e = setTile(layer, x, spY, "-"); e != nil {
                log.Fatalf("%v", e)
            }
        }

        /* Draw the left wall */
        spY = room.startingPointY
        spX = room.startingPointX - room.left
        // Not efficient, but I don't care
        for y := spY - room.top; y <= spY + room.bottom; y++ {
            var e error
            // TODO magic character
            if layer, e = setTile(layer, spX, y, "|"); e != nil {
                log.Fatalf("%v", e)
            }
        }

        /* Draw the right wall */
        spY = room.startingPointY
        spX = room.startingPointX + room.right
        // Not efficient, but I don't care
        for y := spY - room.top; y <= spY + room.bottom; y++ {
            var e error
            // TODO magic character
            if layer, e = setTile(layer, spX, y, "|"); e != nil {
                log.Fatalf("%v", e)
            }
        }

        /* Draw the L corners */
        layer, _ = setTile(layer, room.startingPointX - room.left,
            room.startingPointY - room.top, "L")
        layer, _ = setTile(layer, room.startingPointX + room.right,
            room.startingPointY - room.top, "L")
        layer, _ = setTile(layer, room.startingPointX - room.left,
            room.startingPointY + room.bottom, "L")
        layer, _ = setTile(layer, room.startingPointX + room.right,
            room.startingPointY + room.bottom, "L")
    }

    b.SetLayer(0, layer)
    rl.DumpRooms()
}

/* Checks if a room's spawn point is close to any other room. Currently,
 * rooms must be at least 3 spaces away to count as sufficiently far away. */
func (rl *roomlist) AlreadyRoomThere(r *room) bool {
    minDist := 3

    for _, room := range rl.list {
        if r.startingPointX >= room.startingPointX - minDist &&
                r.startingPointX <= room.startingPointX + minDist &&
                r.startingPointY >= room.startingPointY - minDist &&
                r.startingPointY <= room.startingPointY + minDist {

            //log.Printf("Room generation collision for (%v, %v) with (%v, %v)",
            //    r.startingPointX, r.startingPointY,
            //    room.startingPointY, room.startingPointX)

            return true
        }
    }

    return false
}

/* Attempts to expand the given room by the given amount if possible */
func (rl *roomlist) TryGrowRoom(roomNumber, direction, amount int) bool {
    room := rl.list[roomNumber]
    //log.Printf("Trying to grow room (%v, %v) in direction %v",
    //    room.startingPointX, room.startingPointY, direction)

    /* top */
    if direction == 0 && room.startingPointY - room.top > 0 {
        room.top += 1
    /* right */
    } else if direction == 1 &&
            room.startingPointX + room.right < interfaces.BOARD_WIDTH - 1 {
        room.right += 1
    /* bottom */
    } else if direction == 2 &&
            room.startingPointY + room.bottom < interfaces.BOARD_HEIGHT - 1 {
        room.bottom += 1
    /* left */
    } else if direction == 3 && room.startingPointX - room.left > 0 {
        room.left += 1
    }

    return false
}

func (rl *roomlist) DumpRooms() {
    log.Printf("--- List of Rooms ---")
    for k, room := range rl.list {
        log.Printf("Room %v: at (%v, %v) with size (%v, %v, %v, %v)", k,
            room.startingPointX, room.startingPointY, room.top, room.right,
            room.bottom, room.left)
    }
    log.Printf("--- End List of Rooms ---")
}
