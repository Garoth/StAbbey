package interfaces

import (
)

const (
    UTIL_DIRECTION_NORTH  = 'N'
    UTIL_DIRECTION_SOUTH  = 'S'
    UTIL_DIRECTION_WEST   = 'W'
    UTIL_DIRECTION_EAST   = 'E'

    UTIL_DIRECTION_NORTH_NAME  = "north"
    UTIL_DIRECTION_SOUTH_NAME  = "south"
    UTIL_DIRECTION_WEST_NAME   = "west"
    UTIL_DIRECTION_EAST_NAME   = "east"
)

type Direction interface {
    /* The internal code of the direction */
    Code() byte
    /* The "long name" of the direction */
    Name() string
    /* X, Y differences for the given direction. Ex. N is (0, -1) */
    XYDelta() (int, int)
    /* Returns the (new) opposite direction of this direction */
    Opposite() Direction
}
