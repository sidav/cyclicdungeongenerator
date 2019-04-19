package layout_to_tiled

import (
	"DoomSlayeRL/routines"
	"strings"
)

func getRandomTilemapFromArray(arr *[][]string) *[]string {
	return &((*arr)[routines.Random(len(*arr))]) // ow that's quite of some pointer magic!
}

func getDirectionsByConnsArray(conns *[][]int) (n, e, s, w bool) {
	for i := 0; i < len(*conns); i++ {
		con := (*conns)[i]
		if con[0] == 0 {
			if con[1] == 1 {
				s = true
			}
			if con[1] == -1 {
				n = true
			}
		}
		if con[1] == 0 {
			if con[0] == 1 {
				e = true
			}
			if con[0] == -1 {
				w = true
			}
		}
	}
	return
}

func reverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func getRotatedStringArray(arr *[]string) *[]string { // rotates 90 degrees clockwise
	newArr := make([]string, 0)
	for i := 0; i < len((*arr)[0]); i++ {
		str := ""
		for j := 0; j < len(*arr); j++ {
			str += string((*arr)[j][i])
		}
		newArr = append(newArr, str)
	}
	return &newArr
}

func getMirroredStringArray(arr *[]string, v, h bool) *[]string {
	newArr := make([]string, 0)
	if v && h {
		for i := len(*arr) - 1; i >= 0; i-- {
			newArr = append(newArr, reverseString((*arr)[i]))
		}
		return &newArr
	}
	if v {
		for i := len(*arr) - 1; i >= 0; i-- {
			newArr = append(newArr, (*arr)[i])
		}
		return &newArr
	}
	if h {
		for i := 0; i < len(*arr); i++ {
			newArr = append(newArr, reverseString((*arr)[i]))
		}
		return &newArr
	}
	return nil
}

func getSingleConnTilemap(conn []int) *[]string {
	var tilemap = getRandomTilemapFromArray(&single_entrance_rooms)
	if conn[0] == 0 && conn[1] == 1 { // south-faced door
		return tilemap
	}
	if conn[0] == 0 && conn[1] == -1 { // north-faced door
		return getMirroredStringArray(tilemap, true, false)
	}
	if conn[0] == 1 && conn[1] == 0 { // east-faced door
		return getRotatedStringArray(tilemap)
	}
	if conn[0] == -1 && conn[1] == 0 { // west-faced door
		return getMirroredStringArray(getRotatedStringArray(tilemap), false, true)
	}
	return nil
}

func getTwoConnTilemap(conn [][]int, isRoom bool) *[]string {
	north, east, south, west := getDirectionsByConnsArray(&conn)
	// first, determine whether the connections are symmetric
	if north && south || east && west {
		// they are symmetric, well yeah
		tilemap := getRandomTilemapFromArray(&ns_entrance_corrs)
		if isRoom {
			tilemap = getRandomTilemapFromArray(&ns_entrance_rooms)
		}

		if north && south { // north and south
			return tilemap
		} else { // west-east connection
			return getRotatedStringArray(tilemap)
		}
	} else {
		room := getRandomTilemapFromArray(&se_entrance_corrs)
		if isRoom {
			room = getRandomTilemapFromArray(&se_entrance_rooms)
		}
		if south && east { // south-east
			return room
		}
		if north && east { // north-east
			return getMirroredStringArray(room, true, false)
		}
		if north && west { // north-west
			return getMirroredStringArray(room, true, true)
		}
		if south && west { // south-west
			return getMirroredStringArray(room, false, true)
		}
	}
	return nil
}

func getThreeConnTilemap(conn [][]int, isRoom bool) *[]string {
	tilemap := getRandomTilemapFromArray(&nes_entrance_corrs)
	if isRoom {
		tilemap = getRandomTilemapFromArray(&nes_entrance_rooms)
	}
	north, east, south, west := getDirectionsByConnsArray(&conn)
	if !west {
		return tilemap
	}
	if !east {
		return getMirroredStringArray(tilemap, false, true)
	}
	if !north {
		return getRotatedStringArray(tilemap)
	}
	if !south {
		return getRotatedStringArray(getMirroredStringArray(tilemap, false, true))
	}
	return nil
}

func placePassagesIntoRoomByConnection(room *[]string, conns *[][]int, placeDoors bool) {
	doorChar := "."
	if placeDoors {
		doorChar = "+"
	}
	w, h := len(*room), len((*room)[0])
	for _, con := range *conns {
		cy, cx := con[0], con[1]
		doorx, doory := w/2+cx*w/2, h/2+cy*h/2
		if doorx == w {
			doorx--
		}
		if doory == h {
			doory--
		}
		newStr := (*room)[doorx]
		newStr = newStr[:doory] + doorChar + newStr[doory+1:]
		(*room)[doorx] = newStr
	}
}

func createOutlinedRoomFromRoom(room *[]string) *[]string {
	h := len((*room)[0])
	wall_row := strings.Repeat("#", h+2)
	var outlined_room []string
	outlined_room = append(outlined_room, wall_row)
	for i := range *room {
		outlined_room = append(outlined_room, "#"+(*room)[i]+"#")
	}
	outlined_room = append(outlined_room, wall_row)
	return &outlined_room
}

func getRoomByNodeConnections(conns *[][]int) *[]string {
	var roomTemplate []string
	switch len(*conns) {
	case 0:
		roomTemplate = *getRandomTilemapFromArray(&no_entrance_rooms)
	case 1:
		roomTemplate = *getSingleConnTilemap((*conns)[0])
	case 2:
		roomTemplate = *getTwoConnTilemap(*conns, true)
	case 3:
		roomTemplate = *getThreeConnTilemap(*conns, true)
	case 4:
		roomTemplate = *getRandomTilemapFromArray(&all_entrance_rooms)
	}
	// now we should outline the room with walls
	outlinedRoom := createOutlinedRoomFromRoom(&roomTemplate)
	placePassagesIntoRoomByConnection(outlinedRoom, conns, true)
	return outlinedRoom
}

func getCorridorByNodeConnections(conns *[][]int) *[]string {
	var corrTemplate []string
	switch len(*conns) {
	case 0:
		corrTemplate = *getRandomTilemapFromArray(&no_entrance_corrs)
	case 1:
		corrTemplate = *getSingleConnTilemap((*conns)[0])
	case 2:
		corrTemplate = *getTwoConnTilemap(*conns, false)
	case 3:
		corrTemplate = *getThreeConnTilemap(*conns, false)
	case 4:
		corrTemplate = *getRandomTilemapFromArray(&all_entrance_corrs)
	}
	return &corrTemplate
}


func getTilemapByNodeConnections(conns *[][]int, isRoom bool) *[]string {
	if isRoom {
		return getRoomByNodeConnections(conns)
	}
	return getCorridorByNodeConnections(conns)
}
