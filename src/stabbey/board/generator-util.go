package board

import (
    "log"
    "stabbey/interfaces"
)

func setEntity(game interfaces.Game, entity interfaces.Entity,
        boardId, x, y int) {
    entity.SetPosition(boardId, x, y)
    game.AddEntity(entity)
}


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
    log.Printf("--- Board %v Info ---", b.Id)
    log.Println("Number of rooms:", len(b.RoomList))
    log.Println("Number of doors:", len(b.DoorList))
    log.Println("Number of water tiles:", len(b.WaterList))
    log.Println("Number of decoration tiles:", len(b.GroundDecorList))

    decorline := ""
    for i := 0; i < interfaces.BOARD_WIDTH + 4; i++ {
        decorline = decorline + "-"
    }

    log.Println(decorline)
    for _, row := range b.GetRender() {
        log.Println("|", row, "|")
    }
    log.Println(decorline)
}
