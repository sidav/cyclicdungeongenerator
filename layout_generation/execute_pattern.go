package layout_generation

import (
	rnd "CyclicDungeonGenerator/random"
)
import rpath "CyclicDungeonGenerator/layout_generation/pathfinder"

//ACTION_PLACE_NODE_AT_PATH     = iota
//ACTION_PLACE_NODE_NEAR_PATH   = iota

func execPatternStep(step *patternStep) bool {
	switch step.actionType {
	case ACTION_PLACE_NODE_AT_EMPTY:
		return execPlaceNodeAtEmpty(step)
	case ACTION_PLACE_OBSTACLE_IN_CENTER:
		return execPlaceObstacleInCenter(step)
	//case ACTION_PLACE_RANDOM_OBSTACLES:
	//	return execPlaceRandomObstacles(step)
	case ACTION_PLACE_PATH_FROM_TO:
		return execPlacePathFromTo(step)
	case ACTION_CLEAR_OBSTACLES:
		return execClearObstacles()
	case ACTION_PLACE_NODE_NEAR_PATH:
		return execPlaceNodeNearPath(step)
	case ACTION_PLACE_RANDOM_CONNECTED_NODES:
		return execPlaceRandomConnectedNodes(step)
	case ACTION_FILL_WITH_RANDOM_CONNECTED_NODES:
		return execFillWithRandomConnectedNodes(step)
	case ACTION_SET_NODE_STATUS:
		return execSetNodeStatus(step)
	case ACTION_SET_NODE_CONNECTION_LOCKED_FROM_PATH:
		return execSetNodeConnectionsLockedFromPath(step)
	case ACTION_PLACE_NODE_AT_PATH:
		return execPlaceNodeAtPath(step)
	default:
		panic("No implementation for action!")
	}
	return true
}

func execPlaceNodeAtEmpty(step *patternStep) bool {
	minEmpties := step.minEmptyCellsNear
	var x, y int
	fx, fy, tx, ty := getAbsoluteCoordsForStep(step)
	if fx == 0 && fy == 0 && tx == 0 && ty == 0 { // the coords were not set, so we can use absolutely any ones
		x, y = layout.getRandomEmptyCellCoords(minEmpties)
	} else {
		x, y = layout.getRandomEmptyCellCoordsInRange(fx, fy, tx, ty, minEmpties)
	}
	if x != -1 && y != -1 {
		layout.placeNodeAtCoords(x, y, step.nameOfNode)
		return true
	}
	return false
	panic("execPlaceNodeAtEmpty: Node " + step.nameOfNode + " refuses to be placed!")
}

func execPlaceNodeNearPath(step *patternStep) bool {
	num := step.pathNumber
	px, py, x, y := layout.getRandomPathCoordsAndRandomCellNearPath(num, step.allowPlaceNearNode)
	if px == -1 || py == -1 || x == -1 || y == -1 {
		return false // no cell was returned, step failed...
	}
	layout.placeNodeAtCoords(x, y, step.nameOfNode)
	layout.elements[x][y].setConnectionByCoords(&connection{pathNum: num}, px-x, py-y)
	layout.elements[px][py].setConnectionByCoords(&connection{pathNum: num}, x-px, y-py)
	return true
}

func execPlaceNodeAtPath(step *patternStep) bool {
	num := step.pathNumber
	x, y := layout.getRandomPathCellCoords(num, false)
	if x != -1 && y != -1 {
		layout.placeNodeAtCoords(x, y, step.nameOfNode)
		return true
	}
	return false
}

func execPlaceRandomConnectedNodes(step *patternStep) bool {
	nodesToAdd := rnd.RandInRange(step.countFrom, step.countTo)
	for currNodeNum := 1; currNodeNum <= nodesToAdd; currNodeNum++ {
		px, py, x, y := layout.getRandomNonEmptyCoordsAndRandomCellNearIt()
		if px == -1 || py == -1 || x == -1 || y == -1 {
			if currNodeNum > step.countFrom {
				return true // minimum number of nodes was added anyway, return true.
			}
			return false // no cell was returned, step failed...
		}
		layout.placeNodeAtCoords(x, y, step.nameOfNode)
		layout.elements[x][y].setConnectionByCoords(&connection{}, px-x, py-y)
		layout.elements[px][py].setConnectionByCoords(&connection{}, x-px, y-py)
	}
	return true
}

func execFillWithRandomConnectedNodes(step *patternStep) bool {
	for {
		px, py, x, y := layout.getRandomNonEmptyCoordsAndRandomCellNearIt()
		if px == -1 || py == -1 || x == -1 || y == -1 {
			return true // no more empty spaces to fill
		}
		layout.placeNodeAtCoords(x, y, step.nameOfNode)
		layout.elements[x][y].setConnectionByCoords(&connection{}, px-x, py-y)
		layout.elements[px][py].setConnectionByCoords(&connection{}, x-px, y-py)
	}
}

func execPlaceObstacleInCenter(step *patternStep) bool {
	obstSize := step.obstacleRadius
	cx, cy := layoutWidth/2, layoutHeight/2
	//if size % 2 == 1 {
	//	cx++
	//	cy++
	//}
	for i := -obstSize; i < obstSize+1; i++ {
		for j := -obstSize; j < obstSize+1; j++ {
			if i*i+j*j <= obstSize*obstSize {
				layout.placeObstacleAtCoords(cx+i, cy+j)
			}
		}
	}
	return true
}

//func execPlaceRandomObstacles(step *patternStep) bool {
//	count := getRandomCountForStep(step)
//	for i := 0; i < count; i++ {
//		x, y := layout.getRandomEmptyCellCoords(0)
//		if !(x*y == 0 || x == layoutWidth-1 || y == layoutHeight-1) {
//			layout.placeObstacleAtCoords(x, y)
//		}
//	}
//	return true
//}

func execPlacePathFromTo(step *patternStep) bool {
	pmap := layout.getPassabilityMapForPathfinder()
	fx, fy := layout.getCoordsOfNode(step.nameFrom)
	tx, ty := layout.getCoordsOfNode(step.nameTo)
	path := rpath.FindPath(pmap, fx, fy, tx, ty, false, false, true)
	if path == nil {
		return false
	}
	for path.Child != nil {
		x, y := path.GetCoords()
		vx, vy := path.GetNextStepVector()
		layout.elements[x][y].setConnectionByCoords(&connection{pathNum: step.pathNumber}, vx, vy) // place connection
		path = path.Child
		x, y = path.GetCoords()
		layout.placePathAtCoords(x, y, step.pathNumber)
		layout.elements[x][y].setConnectionByCoords(&connection{pathNum: step.pathNumber}, -vx, -vy) // place reverse connection
	}
	return true
}

func execClearObstacles() bool {
	layout.removeAllObstacles()
	return true
}

func execSetNodeStatus(step *patternStep) bool {
	nname := step.nameOfNode
	status := step.status
	nx, ny := layout.getCoordsOfNode(nname)
	if nx == -1 && ny == -1 {
		return false
	}
	layout.elements[nx][ny].nodeInfo.AddStatus(status)
	return true
}

func execSetNodeConnectionsLockedFromPath(step *patternStep) bool {
	nname := step.nameOfNode
	nx, ny := layout.getCoordsOfNode(nname)
	if nx == -1 && ny == -1 {
		return false
	}
	layout.elements[nx][ny].setAllConnectionsLockedForPath(step.pathNumber, step.lockNumber)
	return true
}

// technical shit below

func getRandomCoordsForStep(step *patternStep) (int, int) {
	fx, fy, tx, ty := getAbsoluteCoordsForStep(step)
	if fx == 0 && fy == 0 && tx == 0 && ty == 0 { // the coords were not set, so we can use absolutely any ones
		// WARNING: may (and will) cause problems if you specially want a cell to be placed at (0,0) and manually set the coords range in step accordingly!
		// TODO: think about tle previous line.
		w, h := layout.GetSize()
		tx = w - 1
		ty = h - 1
	}
	x, y := rnd.RandInRange(fx, tx), rnd.RandInRange(fy, ty)
	return x, y
}

func getRandomCountForStep(step *patternStep) (int) {
	return rnd.RandInRange(step.countFrom, step.countTo)
}

func getAbsoluteCoordsForStep(step *patternStep) (int, int, int, int) {
	fx, fy, tx, ty := step.fx, step.fy, step.tx, step.ty
	if fx < 0 {
		fx = layoutWidth + fx
	}
	if fy < 0 {
		fy = layoutHeight + fy
	}
	if tx < 0 {
		tx = layoutWidth + tx
	}
	if ty < 0 {
		ty = layoutHeight + ty
	}
	return fx, fy, tx, ty
}
