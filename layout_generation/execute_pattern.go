package layout_generation

import (
	rnd "github.com/sidav/golibrl/random"
)
import "github.com/sidav/golibrl/astar"

//ACTION_PLACE_NODE_AT_PATH     = iota
//ACTION_PLACE_NODE_NEAR_PATH   = iota

func execPatternStep(step *patternStep) bool {
	switch step.actionType {
	case ACTION_PLACE_NODE_AT_EMPTY:
		return execPlaceNodeAtEmpty(step)
	case ACTION_PLACE_OBSTACLE_IN_CENTER:
		return execPlaceObstacleInCenter(step)
	case ACTION_PLACE_RANDOM_OBSTACLES:
		return execPlaceRandomObstacles(step)
	case ACTION_PLACE_PATH_FROM_TO:
		return execPlacePathFromTo(step)
	case ACTION_CLEAR_OBSTACLES:
		return execClearObstacles()
	case ACTION_PLACE_NODE_NEAR_PATH:
		return execPlaceNodeNearPath(step)
	}
	return true
}

func execPlaceNodeAtEmpty(step *patternStep) bool {
	const tries = 25
	for try := 0; try < tries; try++ {
		x, y := getRandomCoordsForStep(step)
		if layout.areCoordsEmpty(x, y) {
			layout.placeNodeAtCoords(x, y, step.nameOfNode)
			return true
		}
	}
	return false
	panic("execPlaceNodeAtEmpty: Node " + step.nameOfNode + " refuses to be placed!")
}

func execPlaceNodeNearPath(step *patternStep) bool {
	num := step.pathNumber
	px, py, x, y :=  layout.getRandomPathCoordsAndRandomCellNearPath(num, step.allowPlaceNearNode)
	layout.placeNodeAtCoords(x, y, step.nameOfNode)
	layout.elements[x][y].setConnectionByCoords(&connection{pathNum:num},px-x, py-y)
	layout.elements[px][py].setConnectionByCoords(&connection{pathNum:num},x-px, y-py)
	return true
}

func execPlaceObstacleInCenter(step *patternStep) bool {
	obstSize := step.obstacleRadius
	cx, cy := size/2, size/2
	//if size % 2 == 1 {
	//	cx++
	//	cy++
	//}
	for i := -obstSize; i < obstSize+1; i++ {
		for j := -obstSize; j < obstSize+1; j++ {
			if i*i + j*j <= obstSize*obstSize {
				layout.placeObstacleAtCoords(cx+i, cy+j)
			}
		}
	}
	return true
}

func execPlaceRandomObstacles(step *patternStep) bool {
	count := getRandomCountForStep(step)
	for i := 0; i < count; i++ {
		x, y := layout.getRandomEmptyCellCoords()
		if !(x*y == 0 || x == size-1 || y == size-1) {
			layout.placeObstacleAtCoords(x, y)
		}
	}
	return true
}

func execPlacePathFromTo(step *patternStep) bool {
	pmap := layout.getPassabilityMapForPathfinder()
	fx, fy := layout.getCoordsOfNode(step.nameFrom)
	tx, ty := layout.getCoordsOfNode(step.nameTo)
	path := astar.FindPath(pmap, fx, fy, tx, ty, false, false, true)
	if path == nil {
		return false
	}
	for path.Child != nil {
		x, y := path.GetCoords()
		vx, vy := path.GetNextStepVector()
		layout.elements[x][y].setConnectionByCoords(&connection{pathNum: step.pathNumber}, vx, vy)// place connection
		path = path.Child
		x, y = path.GetCoords()
		layout.placePathAtCoords(x, y, step.pathNumber)
		layout.elements[x][y].setConnectionByCoords(&connection{pathNum: step.pathNumber}, -vx, -vy)// place reverse connection
	}
	return true
}

func execClearObstacles() bool {
	layout.removeAllObstacles()
	return true
}

// technical shit below

func getRandomCoordsForStep(step *patternStep) (int, int) {
	fx, fy, tx, ty := getAbsoluteCoordsForStep(step)
	x, y :=  rnd.RandInRange(fx, tx), rnd.RandInRange(fy, ty)
	return x, y
}

func getRandomCountForStep(step *patternStep) (int) {
	return rnd.RandInRange(step.countFrom, step.countTo)
}

func getAbsoluteCoordsForStep(step *patternStep) (int, int, int, int) {
	fx, fy, tx, ty := step.fx, step.fy, step.tx, step.ty
	if fx < 0 {
		fx = size + fx
	}
	if fy < 0 {
		fy = size + fy
	}
	if tx < 0 {
		tx = size + tx
	}
	if ty < 0 {
		ty = size + ty
	}
	return fx, fy, tx, ty
}

