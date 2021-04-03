package layout_to_generated

import (
	"CyclicDungeonGenerator/layout_generation"
	"fmt"
	"CyclicDungeonGenerator/random/additive_random"
)

type Generator struct {
	Level                       [][]rune
	rnd                         additive_random.FibRandom
	DesiredWidth, DesiredHeight int
	MinRoomXY, MaxRoomXY        int
}

func (g *Generator) ProcessLayout(a *layout_generation.LayoutMap) {
	g.rnd.InitDefault()
	layoutw, layouth := a.GetSize()
	g.Level = make([][]rune, g.DesiredWidth)
	for i := range g.Level {
		g.Level[i] = make([]rune, g.DesiredHeight)
	}
	for lx := 0; lx < layoutw; lx++ {
		for ly := 0; ly < layouth; ly++ {
			node := a.GetElement(lx, ly)
			conns := node.GetAllConnectionsCoords()
			if len(conns) > 0 {
				g.PlaceCorridors(lx, ly, &conns)
				g.PlaceRoom(lx, ly)
			}
		}
	}
}

func (g *Generator) PlaceCorridors(nx, ny int, conns *[][]int) {
	for i := range *conns {
		cx := (*conns)[i][0]
		cy := (*conns)[i][1]
		for x := nx*g.MaxRoomXY + g.MaxRoomXY/2; x <= (nx+cx)*g.MaxRoomXY+g.MaxRoomXY/2; x++ {
			for y := ny*g.MaxRoomXY + g.MaxRoomXY/2; y <= (ny+cy)*g.MaxRoomXY+g.MaxRoomXY/2; y++ {
				g.Level[x][y] = '%'
			}
		}
	}
}

func (g *Generator) PlaceRoom(nx, ny int) {
	roomSize := 0
	var xmin, xmax, ymin, ymax int
	for roomSize == 0 || xmin < 0 || xmax >= g.DesiredWidth || ymin < 0 || ymax >= g.DesiredHeight {
		roomSize = g.rnd.RandInRange(g.MinRoomXY, g.MaxRoomXY)
		xmin = nx*g.MaxRoomXY + g.MaxRoomXY/2-roomSize/2
		xmax = nx*g.MaxRoomXY+g.MaxRoomXY/2+roomSize/2
		ymin = ny*g.MaxRoomXY + g.MaxRoomXY/2
		ymax = ny*g.MaxRoomXY+g.MaxRoomXY/2+roomSize/2
	}
	for x := xmin; x <= xmax; x++ {
		for y := ymin; y <= ymax; y++ {
			if x < 0 || x > g.DesiredWidth || y < 0 || y > g.DesiredHeight {
				panic(fmt.Sprint("%d, %d, %d", roomSize, x, y))
			}
			g.Level[x][y] = '.'

		}
	}
}
