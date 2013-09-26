package util

import (
    "log"
    "math"

    "stabbey/interfaces"
)

type Direction struct {
    code byte
}

func NewDirection(code byte) *Direction {
    me := &Direction{code}

    if me.code != interfaces.UTIL_DIRECTION_NORTH &&
            me.code != interfaces.UTIL_DIRECTION_EAST &&
            me.code != interfaces.UTIL_DIRECTION_SOUTH &&
            me.code != interfaces.UTIL_DIRECTION_WEST {

        log.Fatalf("Attempted to make invalid direction")
    }

    return me
}

func (me *Direction) Code() byte {
    return me.code
}

func (me *Direction) Name() string {
    if me.code == interfaces.UTIL_DIRECTION_NORTH {
        return interfaces.UTIL_DIRECTION_NORTH_NAME
    } else if me.code == interfaces.UTIL_DIRECTION_SOUTH {
        return interfaces.UTIL_DIRECTION_SOUTH_NAME
    } else if me.code == interfaces.UTIL_DIRECTION_WEST {
        return interfaces.UTIL_DIRECTION_WEST_NAME
    } else if me.code == interfaces.UTIL_DIRECTION_EAST {
        return interfaces.UTIL_DIRECTION_EAST_NAME
    } else {
        log.Fatalf("Direction somehow has impossible value")
    }

    /* Impossible */
    return ""
}

func (me *Direction) XYDelta() (int, int) {
    x, y := 0, 0

    if me.Code() == interfaces.UTIL_DIRECTION_NORTH {
        y = -1
    } else if me.Code() == interfaces.UTIL_DIRECTION_SOUTH {
        y = 1
    } else if me.Code() == interfaces.UTIL_DIRECTION_WEST {
        x = -1
    } else if me.Code() == interfaces.UTIL_DIRECTION_EAST {
        x = 1
    } else {
        log.Fatalf("Invalid direction set")
    }

    return x, y
}

func (me *Direction) Opposite() interfaces.Direction {

    if me.Code() == interfaces.UTIL_DIRECTION_NORTH {
        return NewDirection(interfaces.UTIL_DIRECTION_SOUTH)
    } else if me.Code() == interfaces.UTIL_DIRECTION_SOUTH {
        return NewDirection(interfaces.UTIL_DIRECTION_NORTH)
    } else if me.Code() == interfaces.UTIL_DIRECTION_WEST {
        return NewDirection(interfaces.UTIL_DIRECTION_EAST)
    } else if me.Code() == interfaces.UTIL_DIRECTION_EAST {
        return NewDirection(interfaces.UTIL_DIRECTION_WEST)
    } else {
        log.Fatalf("Invalid direction set")
    }

    /* Impossible */
    return &Direction{'X'}
}

/* Returns the rounded integer distance between two points */
func Distance(fromX, fromY, toX, toY int) int {
    dist := math.Sqrt(
        math.Pow(float64(toX - fromX), 2) +
        math.Pow(float64(toY - fromY), 2))
    return int(dist + 0.5)
}

/* Returns the primary direction of travel of a vector */
func PrimaryDirection(fromX, fromY, toX, toY int) interfaces.Direction {
    deltaX, deltaY := toX - fromX, toY - fromY

    /* If tied, then we arbitrate to horizontal axis */
    if math.Abs(float64(deltaX)) >= math.Abs(float64(deltaY)) {
        /* Primary direction is horizontal */
        if deltaX >= 0 {
            return NewDirection(interfaces.UTIL_DIRECTION_EAST)
        } else {
            return NewDirection(interfaces.UTIL_DIRECTION_WEST)
        }
    } else {
        /* Primary direction is vertical */
        if deltaY >= 0 {
            return NewDirection(interfaces.UTIL_DIRECTION_SOUTH)
        } else {
            return NewDirection(interfaces.UTIL_DIRECTION_NORTH)
        }
    }

    /* Impossible to reach statement */
    return &Direction{'X'}
}
