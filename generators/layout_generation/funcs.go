package layout_generation

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func euclideanDistance(fx, fy, tx, ty int) int {
	return abs(fx-tx)+ abs(fy-ty)
}
