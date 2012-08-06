package interfaces

import (
)

const BOARD_WIDTH int = 16
const BOARD_HEIGHT int = 12

type Board interface {
    GetLevel() int
    SetLevel(level int)
    GetLayers() map[int] []string
    SetLayer(layer int, layout []string)
}

type BoardGenerator interface {
    Apply(board Board)
}
