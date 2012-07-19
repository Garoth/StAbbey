package interfaces

import (
)

type Board interface {
    GetLevel() int
    SetLevel(level int)
    GetLayers() map[int] []string
    SetLayer(layer int, layout []string)
}
