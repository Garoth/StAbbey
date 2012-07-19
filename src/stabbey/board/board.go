package board

import (
)

type Board struct {
    /* Unique game level number (i.e. floor number) */
    Level int
    /* Rendering layers of the map, going up in z-index */
    Layers map[int] []string
}

/* Creates a brand new board, for level -- 0, 1, 2, etc */
func New(level int) *Board {
    b := &Board{}
    b.Level = level
    b.Layers = make(map[int] []string, 10)
    b.MakeTestBoard()
    return b
}

func (b *Board) GetLevel() int {
    return b.Level
}

func (b *Board) SetLevel(level int) {
    b.Level = level
}

func (b *Board) GetLayers() map[int] []string {
    return b.Layers
}

func (b *Board) SetLayer(layer int, layout []string) {
    b.Layers[layer] = layout
}


/* Creates a static board for testing */
func (b *Board) MakeTestBoard() {
    b.Layers[0] = []string {"L--------------L",
                            "|..|...........|",
                            "|..|...........|",
                            "|..|...........|",
                            "|..|.----------|",
                            "|..............|",
                            "|..............|",
                            "|-----------L..|",
                            "|...........|..|",
                            "|..............|",
                            "|...........|..|",
                            "L--------------L"}
}
