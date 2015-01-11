package entity

import (
	"log"
	"stabbey/ai"
	"stabbey/interfaces"
	"strconv"
)

func NewShade(g interfaces.Game) interfaces.Entity {
	me := newBasicMonster(g)
	me.SetSubtype(interfaces.ENTITY_MONSTER_SUBTYPE_SHADE)
	me.SetName("Shade " + strconv.Itoa(me.GetEntityId()))
	me.SetMaxArdour(30)
	me.SetArdour(30)

	me.TurnFunction = func(tick int) bool {
		myBoard, myX, myY := me.GetPosition()
		curBoard := g.GetCurrentBoard()

		// Decide if I'm going to move at all
		if myBoard != curBoard {
			return false
		}

		// TODO should check all players and find nearest one to chase
		// players := g.GetPlayers()
		// for k, v := players {
		// }

		// Compute optimal path to nearest player
		player := g.GetPlayer(0)
		_, playerX, playerY := player.GetPosition()
		pather := ai.NewAStar(myX, myY, playerX, playerY, g)
		path := pather.GetOptimalPath()

		// Printing path
		// log.Println("Path is:")
		// for k, v := range path {
		// 	log.Println("  ", k, v)
		// }

		// Move towards them
		if len(path) <= 0 {
			return false
		}
		targetX, targetY := path[0][0], path[0][1]
		if g.CanMoveToSpace(myBoard, targetX, targetY) {
			me.SetPosition(myBoard, targetX, targetY)
			return true
		} else {
			log.Println("Pather gave invalid space for",
				me.GetName()+":", targetX, targetY)
		}

		return false
	}

	return me
}
