package board

import (
    "stabbey/interfaces"
    "stabbey/util"
)

func setTile(layer []string, posX, posY int, char string) ([]string, error) {

    if posX < 0 || posX > interfaces.BOARD_WIDTH - 1 ||
            posY < 0 || posY > interfaces.BOARD_HEIGHT - 1 {

        message := "Board Generator: Can't set character out of bounds"
        return layer, util.NewError(message)
    }

    row := layer[posY]
    row = row[0:posX] + char + row[posX + 1:]
    layer[posY] = row

    return layer, nil
}
