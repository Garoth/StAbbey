package board

import (
    "log"
    "math/rand"
    "strconv"

    "stabbey/interfaces"
)

/* How many rooms can spawn on a wall (will be 0 to max_wal_rooms-1) */
const max_wall_rooms int = 3
const max_rooms int = (max_wall_rooms - 1) * 4 + 3

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
    // TODO need to seed random generator
    log.Printf("Starting board generation")
    // TODO ensure rooms don't spawn on top of each other
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
            y = interfaces.BOAD_HEIGHT - 1
        } else if wall == 3 {
            x = 0
            y = -1
        }

        for nRooms := rand.Intn(max_wall_rooms) - 1; nRooms >= 0; nRooms-- {
            r := &room{}
            rl.list[totalRooms] = r
            totalRooms++

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
                r.startingPointY = rand.Intn(interfaces.BOAD_HEIGHT)
            }

            log.Printf("Seeded room on wall %v at %+v", wall, r)
        }
    }

    /* Now generate some extra rooms in random places */
    for nRooms := rand.Intn(max_rooms-totalRooms + 1); nRooms >= 0; nRooms-- {
        r := &room{}
        rl.list[totalRooms] = r
        totalRooms++
        r.startingPointX = rand.Intn(interfaces.BOARD_WIDTH)
        r.startingPointY = rand.Intn(interfaces.BOAD_HEIGHT)
        log.Printf("Seeded room at %+v", r)
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

    for k, room := range rl.list {
        row := layer[room.startingPointY]
        row = row[0:room.startingPointX] + strconv.Itoa(k) +
            row[room.startingPointX+1:]
        layer[room.startingPointY] = row
    }

    b.SetLayer(0, layer)
}
