package ai

import (
	"log"
	"stabbey/interfaces"

	"github.com/oleiade/lane"
)

type tile struct {
	X, Y int
}

var nilTile = tile{-1, -1}

// Costs of different kinds of tiles we may encounter
const (
	groundCost     = 1
	waterCost      = 3
	fireCost       = 6
	impassableCost = 999999
)

// Guide to understanding A*
// http://www.redblobgames.com/pathfinding/a-star/introduction.html
type AStar struct {
	frontier    *lane.PQueue
	cameFrom    map[tile]tile
	costSoFar   map[tile]int
	start, goal tile
	game        interfaces.Game
}

// Creates a new A* object from one point to another, which can be queried to
// find routing information.
func NewAStar(fromX, fromY, toX, toY int, g interfaces.Game) *AStar {
	me := &AStar{}

	// TODO handle costs of start/goal tiles OR unreachable goal
	me.start = tile{fromX, fromY}
	me.goal = tile{toX, toY}

	me.frontier = lane.NewPQueue(lane.MINPQ)
	me.cameFrom = make(map[tile]tile)
	me.costSoFar = make(map[tile]int)
	me.game = g

	me.frontier.Push(me.start, 0)
	me.cameFrom[me.start] = nilTile
	me.costSoFar[me.start] = 0

	for me.frontier.Size() != 0 {
		currentInterface, _ := me.frontier.Pop()
		current := currentInterface.(tile)

		if current == me.goal {
			break
		}

		for _, next := range me.getNeighbors(current.X, current.Y) {
			newCost := me.costSoFar[current] + me.getTileCost(next)

			if me.costSoFar[next] == 0 || newCost < me.costSoFar[next] {
				me.costSoFar[next] = newCost
				priority := newCost + me.heuristic(me.goal, next)
				me.frontier.Push(next, priority)
				me.cameFrom[next] = current
			}
		}
	}

	return me
}

func (me *AStar) Dump() {
	log.Println("Starting A* Dump")
	log.Println("----------------")
	log.Println("start:", me.start.X, me.start.Y)
	log.Println("end:", me.goal.X, me.goal.Y)
	log.Println("cameFrom:")
	for k, v := range me.cameFrom {
		log.Println("  ", k, v)
	}
	log.Println("costSofar:")
	for k, v := range me.costSoFar {
		log.Println("  ", k, v)
	}
}

// Returns the Manhattan Distance between two tiles. Used as the 'heuristic'
// for A* generally.
func (me *AStar) heuristic(goal, next tile) int {
	diffX := goal.X - next.X
	diffY := goal.Y - next.Y

	// Need absolute value
	if diffX < 0 {
		diffX = diffX * -1
	}

	if diffY < 0 {
		diffY = diffY * -1
	}

	return diffX + diffY
}

// Assigns the movement cost of a given tile, and then returns it
func (me *AStar) getTileCost(t tile) int {
	currentBoard := me.game.GetCurrentBoard()

	if me.game.CanMoveToSpace(currentBoard, t.X, t.Y) == false {
		return impassableCost
	}

	if me.game.IsFire(currentBoard, t.X, t.Y) {
		return fireCost
	}

	if me.game.IsWater(currentBoard, t.X, t.Y) {
		return waterCost
	}

	return groundCost
}

// Returns an unvisited neighbor of the given location. Excludes tiles that are
// off the map.
func (me *AStar) getNeighbors(X, Y int) []tile {
	final := []tile{}

	if Y-1 >= 0 {
		final = append(final, tile{X, Y - 1})
	}
	if Y+1 < interfaces.BOARD_HEIGHT {
		final = append(final, tile{X, Y + 1})
	}
	if X-1 >= 0 {
		final = append(final, tile{X - 1, Y})
	}
	if X+1 < interfaces.BOARD_WIDTH {
		final = append(final, tile{X + 1, Y})
	}

	return final
}

// Returns an optimal path array of X,Y coords from the start position
// to the goal. Start position is excluded.
func (me *AStar) GetOptimalPath() [][2]int {
	array := [][2]int{}
	reversedArray := [][2]int{}
	current := me.goal
	var ok bool

	for true {
		array = append(array, [2]int{current.X, current.Y})

		current, ok = me.cameFrom[current]
		if !ok || current == me.start {
			break
		}
	}

	for i := len(array) - 1; i >= 0; i-- {
		reversedArray = append(reversedArray, array[i])
	}

	return reversedArray
}
