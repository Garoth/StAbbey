package board

import (
    "log"
    "stabbey/interfaces"
)

func setTile(layer []string, posX, posY int, char string) ([]string, error) {

    if posX < 0 || posX > interfaces.BOARD_WIDTH - 1 ||
            posY < 0 || posY > interfaces.BOARD_HEIGHT - 1 {

        return layer, nil
    }

    row := layer[posY]
    row = row[0:posX] + char + row[posX + 1:]
    layer[posY] = row

    return layer, nil
}

func PrintBoardInfo(b *Board) {
    log.Printf("--- Board Rooms ---")
    for k, room := range b.RoomList {
        log.Printf("Room %v: at (%v, %v) with size (%v, %v, %v, %v)", k,
            room.StartingPointX, room.StartingPointY, room.Top, room.Right,
            room.Bottom, room.Left)
    }

    decorline := ""
    for i := 0; i < interfaces.BOARD_WIDTH + 4; i++ {
        decorline = decorline + "-"
    }

    log.Println(decorline)
    for _, row := range b.GetRender() {
        log.Println("|", row, "|")
    }
    log.Println(decorline)

    log.Printf("--- End Board Info ---")
}
