package board

import (
    "math/rand"
    "time"

    "stabbey/interfaces"
)

/* How many rooms can spawn on a wall (will be 0 to max_wal_rooms-1) */
const max_wall_rooms int = 3
const max_rooms int = (max_wall_rooms - 1) * 4 + 1

type growingGen struct {
    board *Board
}

/* Idea of this generator is to pick semi-random starting seeds for
 * a bunch of rooms, and then randomly grow them out until they collide */
func NewGrowingGenerator(b *Board) interfaces.BoardGenerator {
    return &growingGen{b}
}

func (me *growingGen) LoadEntities(game interfaces.Game) {
}

func (me *growingGen) Apply() {
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
            x = me.board.GetWidth() - 1
            y = -1
        } else if wall == 2 {
            x = -1
            y = me.board.GetHeight() - 1
        } else if wall == 3 {
            x = 0
            y = -1
        }

        for nRooms := rand.Intn(max_wall_rooms) - 1; nRooms >= 0; nRooms-- {
            for {
                r := &Room{}

                if x != -1 {
                    r.StartingPointX = x
                    if x == 0 {
                        r.ConstrainedLeft = true
                    } else {
                        r.ConstrainedRight = true
                    }
                } else {
                    r.StartingPointX = rand.Intn(me.board.GetWidth())
                }

                if y != -1 {
                    r.StartingPointY = y
                    if y == 0 {
                        r.ConstrainedTop = true
                    } else {
                        r.ConstrainedBottom = true
                    }
                } else {
                    r.StartingPointY = rand.Intn(me.board.GetHeight())
                }

                if me.AlreadyRoomThere(r) == false {
                    me.board.RoomList[totalRooms] = r
                    /* expand the room by 1 in every direciton right away */
                    for i := 0; i < 4; i++ {
                        me.TryGrowRoom(totalRooms, i, 1)
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
            r := &Room{}
            r.StartingPointX = rand.Intn(me.board.GetWidth())
            r.StartingPointY = rand.Intn(me.board.GetHeight())

            if me.AlreadyRoomThere(r) == false {
                me.board.RoomList[totalRooms] = r
                /* expand the room by 1 in every direciton right away */
                for i := 0; i < 4; i++ {
                    me.TryGrowRoom(totalRooms, i, 1)
                }
                totalRooms++
                break
            }
        }
    }

    /* Randomly grow the rooms a bit, define walls */
    for i := 0; i < 5; i++ {
        for k, _ := range me.board.RoomList {
            me.TryGrowRoom(k, rand.Intn(4), 1)
        }
    }
}

/* Checks if a room's spawn point is close to any other room. Currently,
 * rooms must be at least 3 spaces away to count as sufficiently far away. */
func (me *growingGen) AlreadyRoomThere(r *Room) bool {
    minDist := 3

    for _, room := range me.board.RoomList {
        if r.StartingPointX >= room.StartingPointX - minDist &&
                r.StartingPointX <= room.StartingPointX + minDist &&
                r.StartingPointY >= room.StartingPointY - minDist &&
                r.StartingPointY <= room.StartingPointY + minDist {

            return true
        }
    }

    return false
}

/* Attempts to expand the given room by the given amount if possible */
func (me *growingGen) TryGrowRoom(roomNumber, direction, amount int) bool {
    room := me.board.RoomList[roomNumber]

    /* top */
    if direction == 0 && room.StartingPointY - room.Top > 0 {
        room.Top += 1
    /* right */
    } else if direction == 1 &&
            room.StartingPointX + room.Right < me.board.GetWidth() - 1 {
        room.Right += 1
    /* bottom */
    } else if direction == 2 &&
            room.StartingPointY + room.Bottom < me.board.GetHeight() - 1 {
        room.Bottom += 1
    /* left */
    } else if direction == 3 && room.StartingPointX - room.Left > 0 {
        room.Left += 1
    }

    return false
}
