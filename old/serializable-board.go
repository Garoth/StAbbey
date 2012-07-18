package stabbey

import (
)

type SerializableBoard struct {
    /* Unique game level number (i.e. floor number) */
    Id int
    /* Rendering layers of the map, going up in z-index */
    Layers [BOARD_NUM_LAYERS][BOARD_HEIGHT]string
}

/* Converts a layer represented as a big long string to an array of strings */
func layerStringToStringArray(layerStr string) [BOARD_HEIGHT]string {
    var layerArrays [BOARD_HEIGHT]string

    if len(layerStr) != BOARD_HEIGHT * BOARD_WIDTH {
        return layerArrays
    }

    for x := 0; x < BOARD_HEIGHT; x++ {
        layerArrays[x] = layerStr[x*BOARD_WIDTH:x * BOARD_WIDTH + BOARD_WIDTH]
    }

    return layerArrays
}

func NewSerializableBoard(b *Board) *SerializableBoard {
    sb := &SerializableBoard{}

    sb.Id = b.Id
    sb.Layers[0] = layerStringToStringArray(b.Layer0)
    sb.Layers[1] = layerStringToStringArray(b.Layer1)
    sb.Layers[2] = layerStringToStringArray(b.Layer2)
    sb.Layers[3] = layerStringToStringArray(b.Layer3)
    sb.Layers[4] = layerStringToStringArray(b.Layer4)
    sb.Layers[5] = layerStringToStringArray(b.Layer5)
    sb.Layers[6] = layerStringToStringArray(b.Layer6)
    sb.Layers[7] = layerStringToStringArray(b.Layer7)

    return sb
}
