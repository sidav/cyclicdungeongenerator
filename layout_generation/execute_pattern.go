package layout_generation

import (
	rnd "github.com/sidav/golibrl/random"
)
import "github.com/sidav/golibrl/astar"

//ACTION_PLACE_NODE_AT_PATH     = iota
//ACTION_PLACE_NODE_NEAR_PATH   = iota

var (
	currPathNumber = 1
)


func execPatternStep(step *patternStep) {
	switch step.actionType {
	case ACTION_PLACE_NODE_AT_EMPTY:
		execPlaceNodeAtEmpty(step)
	case ACTION_PLACE_OBSTACLE_IN_CENTER:
		execPlaceObstacleInCenter(step)
	case ACTION_PLACE_RANDOM_OBSTACLES:
		execPlaceRandomObstacles(step)
	case ACTION_PLACE_PATH_FROM_TO:
		execPlacePathFromTo(step)
	case ACTION_CLEAR_OBSTACLES:
		execClearObstacles()
	case ACTION_PLACE_NODE_NEAR_PATH:
		execPlaceNodeNearPath(step)
	}
}

func execPlaceNodeAtEmpty(step *patternStep) {
	x, y :=  getRandomCoordsForStep(step)
	layout.placeNodeAtCoords(x, y, step.nameOfNode)
}

func execPlaceNodeNearPath(step *patternStep) {
	num := step.pathNumber
	x, y :=  layout.getRandomCellNearPath(num)
	layout.placeNodeAtCoords(x, y, step.nameOfNode)
}

func execPlaceObstacleInCenter(step *patternStep) {
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
}

func execPlaceRandomObstacles(step *patternStep) {
	count := getRandomCountForStep(step)
	for i := 0; i < count; i++ {
		x, y := layout.getRandomEmptyCellCoords()
		if !(x*y == 0 || x == size-1 || y == size-1) {
			layout.placeObstacleAtCoords(x, y)
		}
	}
}

func execPlacePathFromTo(step *patternStep) {
	pmap := layout.getPassabilityMapForPathfinder()
	fx, fy := layout.getCoordsOfNode(step.nameFrom)
	tx, ty := layout.getCoordsOfNode(step.nameTo)
	path := astar.FindPath(pmap, fx, fy, tx, ty, false, true)
	if path == nil {
		return
	}
	for path.Child != nil {
		path = path.Child
		x, y := path.GetCoords()
		layout.placePathAtCoords(x, y, currPathNumber)
	}
	currPathNumber++
}

func execClearObstacles() {
	layout.removeAllObstacles()
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

