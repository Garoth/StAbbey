package util

import (
    "math"
)

/* Returns the rounded integer distance between two points */
func Distance(fromX, fromY, toX, toY int) int {
    dist := math.Sqrt(
        math.Pow(float64(toX - fromX), 2) +
        math.Pow(float64(toY - fromY), 2))
    return int(dist + 0.5)
}

/* Returns the primary direction of travel of a vector */
func PrimaryDirection(fromX, fromY, toX, toY int) byte {
    deltaX, deltaY := toX - fromX, toY - fromY

    /* If tied, then we arbitrate to horizontal axis */
    if math.Abs(float64(deltaX)) >= math.Abs(float64(deltaY)) {
        /* Primary direction is horizontal */
        if deltaX >= 0 {
            return 'E'
        } else {
            return 'W'
        }
    } else {
        /* Primary direction is vertical */
        if deltaY >= 0 {
            return 'S'
        } else {
            return 'N'
        }
    }

    /* Impossible to reach statement */
    return 'X'
}
