package interfaces

import (
)

const BOARD_WIDTH int = 16
const BOARD_HEIGHT int = 12

type Board interface {
    GetLevel() int
    SetLevel(int)
    GetWidth() int
    SetWidth(int)
    GetHeight() int
    SetHeight(int)
    GetRender() []string
}

type BoardGenerator interface {
    Apply()
}
