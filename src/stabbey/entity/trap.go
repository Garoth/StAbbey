package entity

import (
    "log"
    "strconv"
    "stabbey/interfaces"
)

func NewTeleportTrap(g interfaces.Game, x, y int) *Entity {
    me := newBasicTrigger(g)
    me.SetSubtype(interfaces.ENTITY_TRIGGER_SUBTYPE_TELEPORT_TRAP)
    me.SetName("Teleport trap " + strconv.Itoa(me.GetEntityId()) +  " to " +
        strconv.Itoa(x) + ", " + strconv.Itoa(y))

    me.TroddenFunction = func(by interfaces.Entity) {
        /* I've already triggered */
        if me.IsDead() {
            return
        }

        if me.Game.CanMoveToSpace(g.GetCurrentBoard(), x, y) {
            log.Println("Teleporting", by.GetName(), "to", x, y)
            by.SetPosition(me.Game.GetCurrentBoard(), x, y)
        } else {
            log.Println(me.GetName() + " failed, since destination is blocked")
        }
        me.Die()

        sprungTrap := NewSprungTrap(g)
        sprungTrap.SetPosition(me.GetPosition())
        me.Game.AddEntity(sprungTrap)
    }

    return me
}

func NewCaltropTrap(g interfaces.Game) *Entity {
    damage := 20
    me := newBasicTrigger(g)
    me.SetSubtype(interfaces.ENTITY_TRIGGER_SUBTYPE_CALTROP_TRAP)
    me.SetName("Caltrop trap " + strconv.Itoa(me.GetEntityId()))

    me.TroddenFunction = func(by interfaces.Entity) {
        if by.IsTangible() {
            log.Println(me.GetName(), "hurt", by.GetName(), "for", damage)
            by.ChangeArdour(-damage)
        }
    }

    return me
}

/* Makes a boulder trap that throws a boulder from the nearest wall
 * in the giver startWallDir, which is either N, E, S, or W */
func NewBoulderTrap(g interfaces.Game, startWallDir byte) *Entity {
    me := newBasicTrigger(g)
    me.SetSubtype(interfaces.ENTITY_TRIGGER_SUBTYPE_BOULDER_TRAP)
    me.SetName("Boulder trap " + strconv.Itoa(me.GetEntityId()) +
        " from " + string(startWallDir) + " wall")

    me.TroddenFunction = func(by interfaces.Entity) {
        if by.IsTangible() == false || me.IsDead() {
            return
        }

        /* TODO this is real trashy code */
        var boulderDir byte
        x, y := 0, 0
        if startWallDir == 'N' {
            y = -1
            boulderDir = 'S'
        } else if startWallDir == 'S' {
            y = 1
            boulderDir = 'N'
        } else if startWallDir == 'W' {
            x = -1
            boulderDir = 'E'
        } else if startWallDir == 'E' {
            x = 1
            boulderDir = 'W'
        } else {
            log.Fatalf("Direction must be one of N, E, S, W")
        }

        boardId, boulderX, boulderY := me.GetPosition()
        for ; me.Game.IsWall(boardId, boulderX, boulderY) == false; {
            boulderX += x
            boulderY += y
        }
        boulderX -= x
        boulderY -= y

        boulder := NewBoulder(me.Game, boulderDir)
        me.Game.PlaceAtNearestTile(boulder, boardId, boulderX, boulderY)
        me.Game.AddEntity(boulder)

        me.Die()
        sprungTrap := NewSprungTrap(g)
        sprungTrap.SetPosition(me.GetPosition())
        me.Game.AddEntity(sprungTrap)
    }

    return me
}

/* Spawns a rolling boulder in a given direction, which is one of
 * N, E, S, or W */
func NewBoulder(g interfaces.Game, travelDir byte) *Entity {
    damage := 70

    me := newBasicMonster(g)
    me.SetSubtype(interfaces.ENTITY_MONSTER_SUBTYPE_BOULDER)
    me.SetName("Boulder " + strconv.Itoa(me.GetEntityId()))

    me.TickFunction = func(tick int) {
        if me.IsDead() {
            return
        }

        /* TODO this is real trashy code */
        x, y := 0, 0
        if travelDir == 'N' {
            y = -1
        } else if travelDir == 'S' {
            y = 1
        } else if travelDir == 'W' {
            x = -1
        } else if travelDir == 'E' {
            x = 1
        } else {
            log.Fatalf("Direction must be one of N, E, S, W")
        }

        boardId, myX, myY := me.GetPosition()
        myX += x
        myY += y

        entity := me.Game.GetTangibleEntityAtSpace(boardId, myX, myY)
        if entity != nil {
            me.SwapPositionWith(entity)
            entity.ChangeArdour(-damage)
            log.Println(me.GetName(), "ran over", entity.GetName())
        } else {
            if me.Game.IsWall(boardId, myX, myY) == false {
                me.SetPosition(boardId, myX, myY)
            } else {
                me.Die()
            }
        }
    }

    return me
}

func NewSprungTrap(g interfaces.Game) *Entity {
    me := newBasicInert(g)
    me.SetSubtype(interfaces.ENTITY_INERT_SUBTYPE_TRAP)
    me.SetTangible(false)
    me.SetName("Sprung trap " + strconv.Itoa(me.GetEntityId()))
    return me
}
