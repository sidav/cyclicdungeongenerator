package pathfinder

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// DEPRECATED
// global func for compatibility
func FindPath(costMap *[][]int, fromx, fromy, tox, toy int, diagonalMoveAllowed bool, forceGetPath, forceIncludeFinish bool) *Cell {
	pf := AStarPathfinder{
		diagonalMoveAllowed,
		forceGetPath,
		forceIncludeFinish,
		false,
	}
	return pf.FindPath(costMap, fromx, fromy, tox, toy)
}
