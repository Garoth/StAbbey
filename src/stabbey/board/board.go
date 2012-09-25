package board

import (
    "math/rand"
    "time"
    "stabbey/interfaces"
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
    NewGrowingGenerator().Apply(b)
    return b
}

/* Picks a random spawn point. TODO: should be in Game so that things can't
 * spawn over entities */
func (b *Board) GetRandomSpawnPoint() (int, int) {
    maxAttempts := 1000
    rand.Seed(time.Now().Unix())

    for x := 0; x < maxAttempts; x++ {
        x := rand.Intn(interfaces.BOARD_WIDTH);
        y := rand.Intn(interfaces.BOARD_HEIGHT);

        if b.Layers[0][y][x] == '.' {
            return x, y;
        }
    }

    /* TODO should never happen */
    return 0, 0
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
