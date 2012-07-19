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

    sb.Level = b.GetLevel()
    sb.Layers = make([][]string, len(b.GetLayers()))
    for k, layer := range b.GetLayers() {
        sb.Layers[k] = layer
    }

    return sb
}
