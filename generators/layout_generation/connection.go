package layout_generation

type connection struct {
	// This is an allowed (walkable) transition between the cells. May or may not be a door.
	// Each cell has up to 4 transitions (north, east, etc).
	IsNodeExtension bool // if true, it will be "wall-less" connection.
	IsLocked        bool
	LockNum         int
	pathNum         int
}
