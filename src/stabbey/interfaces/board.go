package interfaces

import (
)

const BOARD_WIDTH int = 16
const BOARD_HEIGHT int = 12

type Board interface {
    LoadStartingEntities()
    WarpPlayersToStart()
    GetId() int
    SetId(int)
    GetWidth() int
    SetWidth(int)
    GetHeight() int
    SetHeight(int)
    GetRender() []string
}

type BoardGenerator interface {
    Apply()
    LoadEntities(Game)
}
