package layout_generation

import (
	"cyclicdungeongenerator/generator/layout_generation/pathfinder"
	"fmt"
)

//ACTION_PLACE_NODE_AT_PATH     = iota
//ACTION_PLACE_NODE_NEAR_PATH   = iota

func (step *patternStep) execPatternStep(layout *LayoutMap) bool {
	switch step.actionType {
	case ACTION_PLACE_NODE_AT_EMPTY:
		return step.execPlaceNodeAtEmpty(layout)
	case ACTION_PLACE_OBSTACLE_IN_CENTER:
		return step.execPlaceObstacleInCenter(layout)
	case ACTION_PLACE_OBSTACLE_AT_COORDS:
		return step.execPlaceObstacleAtCoords(layout)
	//case ACTION_PLACE_RANDOM_OBSTACLES:
	//	return execPlaceRandomObstacles(layout)
	case ACTION_PLACE_PATH_FROM_TO:
		return step.execPlacePathFromTo(layout)
	case ACTION_CLEAR_OBSTACLES:
		return step.execClearObstacles(layout)
	case ACTION_PLACE_NODE_NEAR_PATH:
		return step.execPlaceNodeNearPath(layout)
	case ACTION_PLACE_RANDOM_CONNECTED_NODES:
		return step.execPlaceRandomConnectedNodes(layout)
	case ACTION_FILL_WITH_RANDOM_CONNECTED_NODES:
		return step.execFillWithRandomConnectedNodes(layout)
	case ACTION_SET_NODE_TAGS:
		return step.execSetNodeTags(layout)
	case ACTION_LOCK_PATH:
		return step.execLockPath(layout)
	case ACTION_SET_NODE_CONNECTIONS_LOCKED:
		return step.execSetNodeConnectionsLocked(layout)
	case ACTION_PLACE_NODE_AT_PATH:
		return step.execPlaceNodeAtPath(layout)
	case ACTION_GROW_ALL_NODES:
		layout.growAllNodesToFillSpace(step.maxNodeSize)
	default:
		panic("No implementation for action!")
	}
	return true
}

func (step *patternStep) execPlaceNodeAtEmpty(layout *LayoutMap) bool {
	minEmpties := step.minEmptyCellsNear
	totalConnections := step.pattern.getTotalConnectionsForNodeWithName(step.nameOfNode)
	if minEmpties == 0 {
		minEmpties = totalConnections
	}
	var x, y int
	fx, fy, tx, ty := step.getAbsoluteCoordsForStep(layout)
	if fx == 0 && fy == 0 && tx == 0 && ty == 0 { // the coords were not set, so we can use absolutely any ones
		// Don't place the node in the corner if there should be more than 2 connections for it
		cornerAllowed := totalConnections <= 2
		// Don't place the node at the edge if there should be more than 3 connections for it
		edgeAllowed := totalConnections <= 3
		dists := step.pattern.getAllMinDistancesForNode(step.nameOfNode)
		x, y = layout.getRandomEmptyCellCoords(minEmpties, cornerAllowed, edgeAllowed, dists)
		//fmt.Printf("%s: %v\n", step.nameOfNode, [2]int{x,y})
	} else {
		x, y = layout.getRandomEmptyCellCoordsInRange(fx, fy, tx, ty, minEmpties)
	}
	if x != -1 && y != -1 {
		layout.placeNodeAtCoords(x, y, step.nameOfNode)
		layout.elements[x][y].nodeInfo.setTags(step.tags)
		if step.maxNodeSize > 1 {
			for growStep := 0; growStep < layout.rnd.Rand(step.maxNodeSize); growStep++ {
				layout.tryGrowingNodeByName(step.nameOfNode)
			}
		}
		return true
	}
	return false
	panic("execPlaceNodeAtEmpty: Node " + step.nameOfNode + " refuses to be placed!")
}

func (step *patternStep) execPlaceNodeNearPath(layout *LayoutMap) bool {
	num := step.pathNumber
	px, py, x, y := layout.getRandomPathCoordsAndRandomCellNearPath(num, step.allowPlaceNearNode)
	if px == -1 || py == -1 || x == -1 || y == -1 {
		return false // no cell was returned, step failed...
	}
	layout.placeNodeAtCoords(x, y, step.nameOfNode)
	layout.elements[x][y].nodeInfo.setTags(step.tags)
	layout.setConnectionsBetweenTwoCoords(&connection{pathNum: num}, x, y, px, py)
	if step.maxNodeSize > 1 {
		for growStep := 0; growStep < layout.rnd.Rand(step.maxNodeSize); growStep++ {
			layout.tryGrowingNodeByName(step.nameOfNode)
		}
	}
	return true
}

func (step *patternStep) execPlaceNodeAtPath(layout *LayoutMap) bool {
	num := step.pathNumber
	x, y := layout.getRandomPathCellCoords(num, false)
	if x != -1 && y != -1 {
		layout.placeNodeAtCoords(x, y, step.nameOfNode)
		layout.elements[x][y].nodeInfo.setTags(step.tags)
		if step.maxNodeSize > 1 {
			for growStep := 0; growStep < layout.rnd.Rand(step.maxNodeSize); growStep++ {
				layout.tryGrowingNodeByName(step.nameOfNode)
			}
		}
		return true
	}
	return false
}

func (step *patternStep) execPlaceRandomConnectedNodes(layout *LayoutMap) bool {
	nodesToAdd := layout.rnd.RandInRange(step.countFrom, step.countTo)
	for currNodeNum := 1; currNodeNum <= nodesToAdd; currNodeNum++ {
		px, py, x, y := layout.getRandomNonEmptyCoordsAndRandomCellNearIt()
		if px == -1 || py == -1 || x == -1 || y == -1 {
			if currNodeNum > step.countFrom {
				return true // minimum number of nodes was added anyway, return true.
			}
			return false // no cell was returned, step failed...
		}
		layout.placeNodeAtCoords(x, y, step.nameOfNode)
		layout.setConnectionsBetweenTwoCoords(&connection{}, x, y, px, py)
	}
	return true
}

func (step *patternStep) execFillWithRandomConnectedNodes(layout *LayoutMap) bool {
	for {
		px, py, x, y := layout.getRandomNonEmptyCoordsAndRandomCellNearIt()
		if px == -1 || py == -1 || x == -1 || y == -1 {
			return true // no more empty spaces to fill
		}
		layout.placeNodeAtCoords(x, y, step.nameOfNode)
		layout.setConnectionsBetweenTwoCoords(&connection{}, x, y, px, py)
	}
}

func (step *patternStep) execPlaceObstacleInCenter(layout *LayoutMap) bool {
	obstSize := step.obstacleRadius
	layoutWidth, layoutHeight := layout.GetSize()
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

func (step *patternStep) execPlaceObstacleAtCoords(layout *LayoutMap) bool {
	fx, fy, tx, ty := step.getAbsoluteCoordsForStep(layout)
	for x := fx; x <= tx; x++ {
		for y := fy; y <= ty; y++ {
			layout.placeObstacleAtCoords(x, y)
		}
	}
	return true
}

//func (step *patternStep)execPlaceRandomObstacles(layout *LayoutMap) bool {
//	count := getRandomCountForStep(step)
//	for i := 0; i < count; i++ {
//		x, y := layout.getRandomEmptyCellCoords(0)
//		if !(x*y == 0 || x == LayoutWidth-1 || y == LayoutHeight-1) {
//			layout.placeObstacleAtCoords(x, y)
//		}
//	}
//	return true
//}

func (step *patternStep) execPlacePathFromTo(layout *LayoutMap) bool {
	pmap := layout.getPassabilityMapForPathfinder(step.allowCrossPaths)
	froms := layout.getAllCoordsOfNode(step.nameFrom)
	tos := layout.getAllCoordsOfNode(step.nameTo)
	randomFromIndex := layout.rnd.Rand(len(froms))
	randomToIndex := layout.rnd.Rand(len(tos))

	for indFrom := range froms {
		from := froms[(indFrom + randomFromIndex) % len(froms)]
		for indTo := range tos {
			to := tos[(indTo + randomToIndex) % len(tos)]
			fx, fy := from[0], from[1]
			tx, ty := to[0], to[1]
			path := pathfinder.FindPath(pmap, fx, fy, tx, ty, false, false, true)
			if path == nil {
				continue // try another coords
			}
			pathLength := 0
			for path.Child != nil {
				x, y := path.GetCoords()
				vx, vy := path.GetNextStepVector()
				layout.elements[x][y].setConnectionByCoords(&connection{pathNum: step.pathNumber}, vx, vy) // place connection
				path = path.Child
				x, y = path.GetCoords()
				layout.placePathAtCoords(x, y, step.pathNumber)
				layout.elements[x][y].setConnectionByCoords(&connection{pathNum: step.pathNumber}, -vx, -vy) // place reverse connection
				pathLength += 1
			}

			// check if the path is too short for following PLACE_ROOM_AT_PATH to ever be finished
			if pathLength >= step.pattern.getTotalNodesToBePlacedAtPath(step.pathNumber)+1 {
				return true
			} else {
				panic(fmt.Sprintf("IT DOESN'T WORK WTF %s: %d, %v, %v", step.instructionText, pathLength,
					step.pattern.getAllMinDistancesForNode(step.nameFrom), step.pattern.getAllMinDistancesForNode(step.nameTo)))
			}
		}
	}
	return false
}

func (step *patternStep) execClearObstacles(layout *LayoutMap) bool {
	layout.removeAllObstacles()
	return true
}

func (step *patternStep) execSetNodeTags(layout *LayoutMap) bool {
	nname := step.nameOfNode
	tags := step.tags
	if nname != "" {
		nx, ny := layout.getAnyOfCoordsOfNode(nname)
		if nx == -1 && ny == -1 {
			return false
		}
		layout.elements[nx][ny].nodeInfo.setTags(tags)
	} else { // set tags to random untagged node
		suitableCoords := make([][2]int, 0)
		w, h := layout.GetSize()
		for x := 0; x < w; x++ {
			for y := 0; y < h; y++ {
				elem := layout.GetElement(x, y)
				if elem.IsNode() && elem.nodeInfo.nodeTag == "" {
					suitableCoords = append(suitableCoords, [2]int{x, y})
				}
			}
		}
		if len(suitableCoords) == 0 {
			return false
		}
		randIndex := layout.rnd.Rand(len(suitableCoords))
		layout.GetElement(suitableCoords[randIndex][0], suitableCoords[randIndex][1]).nodeInfo.setTags(step.tags)
	}
	return true
}

func (step *patternStep) execLockPath(layout *LayoutMap) bool {
	w, h := layout.GetSize()
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			elem := layout.GetElement(x, y)
			if elem.IsNode() {
				layout.setAllNodeConnectionsLocked(x, y, step.pathNumber, step.lockNumber)
			}
		}
	}
	return true
}

func (step *patternStep) execSetNodeConnectionsLocked(layout *LayoutMap) bool {
	nname := step.nameOfNode
	allNodeCoords := layout.getAllCoordsOfNode(nname)
	if len(allNodeCoords) == 0 {
		return false
	}
	for _, coords := range allNodeCoords {
		nx, ny := coords[0], coords[1]
		layout.setAllNodeConnectionsLocked(nx, ny, step.pathNumber, step.lockNumber)
	}
	return true
}

// technical shit below

func (step *patternStep) getRandomCoordsForStep(layout *LayoutMap) (int, int) {
	fx, fy, tx, ty := step.getAbsoluteCoordsForStep(layout)
	if fx == 0 && fy == 0 && tx == 0 && ty == 0 { // the coords were not set, so we can use absolutely any ones
		// WARNING: may (and will) cause problems if you specially want a cell to be placed at (0,0) and manually set the coords range in step accordingly!
		// TODO: think about tle previous line.
		w, h := layout.GetSize()
		tx = w - 1
		ty = h - 1
	}
	x, y := layout.rnd.RandInRange(fx, tx), layout.rnd.RandInRange(fy, ty)
	return x, y
}

func (step *patternStep) getRandomCountForStep(layout *LayoutMap) int {
	return layout.rnd.RandInRange(step.countFrom, step.countTo)
}

func (step *patternStep) getAbsoluteCoordsForStep(layout *LayoutMap) (int, int, int, int) {
	layoutWidth, layoutHeight := layout.GetSize()
	fx, fy, tx, ty := step.fromX, step.fromY, step.toX, step.toY
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
