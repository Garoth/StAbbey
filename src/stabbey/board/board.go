package board

import (
    "log"
    "stabbey/interfaces"
)

type Room struct {
    StartingPointX, StartingPointY, Left, Right, Top, Bottom int
    ConstrainedLeft, ConstrainedRight, ConstrainedTop, ConstrainedBottom bool
}

type Tile struct {
    LocX, LocY int
}

type GroundDecor struct {
    LocX, LocY int
    DecorType, TextMapChar string
}

type Board struct {
    PlayerStartX, PlayerStartY, Id, Width, Height int
    /* Game object (for adding entities) */
    Game interfaces.Game
    /* Board generator */
    Generator interfaces.BoardGenerator
    /* List of rooms */
    RoomList map[int]*Room
    /* List of "doors" -- tiles that override walls */
    DoorList map[int]*Tile
    /* List of water tiles */
    WaterList map[int]*Tile
    /* List of ground decoration tiles */
    GroundDecorList map[int]*GroundDecor
}

/* Creates a brand new board, for id -- 0, 1, 2, etc */
func New(id int, game interfaces.Game) *Board {
    b := &Board{}
    b.Game = game
    b.Id = id
    b.Width = interfaces.BOARD_WIDTH
    b.Height = interfaces.BOARD_HEIGHT
    b.RoomList = make(map[int]*Room)
    b.DoorList = make(map[int]*Tile)
    b.WaterList = make(map[int]*Tile)
    b.GroundDecorList = make(map[int]*GroundDecor)
    b.pickGenerator()
    return b
}

func (b *Board) pickGenerator() {
    if b.Id == 0 {
        b.Generator = NewShadowsLevelOneGenerator(b)
    } else if b.Id == 1 {
        b.Generator = NewEntranceGenerator(b)
    } else {
        b.Generator = NewGrowingGenerator(b)
    }

    b.Generator.Apply()
    PrintBoardInfo(b)
}

func (b *Board) LoadStartingEntities() {
    b.Generator.LoadEntities(b.Game)
}

func (b *Board) WarpPlayersToStart() {
    for _, player := range b.Game.GetPlayers() {
        b.Game.PlaceAtNearestTile(player, b.Id,
            b.PlayerStartX, b.PlayerStartY)
    }
}

func (b *Board) GetRender() []string {
    /* And finally, write the rooms to the board */
    layer := []string {"................",
                       "................",
                       "................",
                       "................",
                       "................",
                       "................",
                       "................",
                       "................",
                       "................",
                       "................",
                       "................",
                       "................"}

    for _, room := range b.RoomList {
        /* Draw the top wall */
        spY := room.StartingPointY - room.Top
        spX := room.StartingPointX
        // Not efficient, but I don't care
        for x := spX - room.Left; x <= spX + room.Right; x++ {
            var e error
            // TODO magic character
            if layer, e = setTile(layer, x, spY, "#"); e != nil {
                log.Fatalf("%v", e)
            }
        }

        /* Draw the bottom wall */
        spY = room.StartingPointY + room.Bottom
        spX = room.StartingPointX
        // Not efficient, but I don't care
        for x := spX - room.Left; x <= spX + room.Right; x++ {
            var e error
            // TODO magic character
            if layer, e = setTile(layer, x, spY, "#"); e != nil {
                log.Fatalf("%v", e)
            }
        }

        /* Draw the left wall */
        spY = room.StartingPointY
        spX = room.StartingPointX - room.Left
        // Not efficient, but I don't care
        for y := spY - room.Top; y <= spY + room.Bottom; y++ {
            var e error
            // TODO magic character
            if layer, e = setTile(layer, spX, y, "#"); e != nil {
                log.Fatalf("%v", e)
            }
        }

        /* Draw the right wall */
        spY = room.StartingPointY
        spX = room.StartingPointX + room.Right
        // Not efficient, but I don't care
        for y := spY - room.Top; y <= spY + room.Bottom; y++ {
            var e error
            // TODO magic character
            if layer, e = setTile(layer, spX, y, "#"); e != nil {
                log.Fatalf("%v", e)
            }
        }
    }

    for _, door := range b.DoorList {
        var e error
        if layer, e = setTile(layer, door.LocX, door.LocY, "|"); e != nil {
            log.Fatalln(e)
        }
    }

    for _, water := range b.WaterList {
        var e error
        if layer, e = setTile(layer, water.LocX, water.LocY, "~"); e != nil {
            log.Fatalln(e)
        }
    }

    for _, decor := range b.GroundDecorList {
        var e error
        if layer, e = setTile(layer, decor.LocX, decor.LocY,
                decor.TextMapChar); e != nil {

            log.Fatalln(e)
        }
    }

    return layer
}

func (b *Board) GetId() int {
    return b.Id
}

func (b *Board) SetId(id int) {
    b.Id = id
}

func (b *Board) GetWidth() int {
    return b.Width
}

func (b *Board) SetWidth(w int) {
    b.Width = w
}

func (b *Board) GetHeight() int {
    return b.Height
}

func (b *Board) SetHeight(h int) {
    b.Height = h
}


func (b *Board) simpleAddRoom(startX, startY, width, height int) {
    /* StartX, StartY, Left, Right, Top, Bottom, ... constraints */
    b.RoomList[len(b.RoomList)] = &Room{startX, startY, 1, width, 1, height,
        false, false, false, false}
}

func (b *Board) addDoor(x, y int) {
    b.DoorList[len(b.DoorList)] = &Tile{x, y}
}

func (b *Board) addWater(x, y int) {
    b.WaterList[len(b.WaterList)] = &Tile{x, y}
}

func (b *Board) addDecoration(x, y int, name string) {
    var tileChar string

    switch name {
    case "grass": tileChar = "g"
    case "carpet": tileChar = "c"
    }

    b.GroundDecorList[len(b.GroundDecorList)] = &GroundDecor{x, y, name, tileChar}
}
