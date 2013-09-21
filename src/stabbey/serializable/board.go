package serializable

import (
    "stabbey/interfaces"
)

type Board struct {
    /* Unique game level number (i.e. floor number) */
    Level int
    /* Rendering layers of the map, going up in z-index */
    Layers [][]string
}

func NewBoard(b interfaces.Board) *Board {
    sb := &Board{}

    sb.Level = b.GetId()
    sb.Layers = make([][]string, 1)
    sb.Layers[0] = b.GetRender()

    return sb
}
