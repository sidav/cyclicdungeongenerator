package layout_to_tiled

import (
	"DoomSlayeRL/routines"
	"strings"
)

func getRandomRoomFromArray(arr *[][]string) *[]string {
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

func getSingleConnRoom(conn []int) *[]string {
	var room = getRandomRoomFromArray(&single_entrance_rooms)
	if conn[0] == 0 && conn[1] == 1 { // south-faced door
		return room
	}
	if conn[0] == 0 && conn[1] == -1 { // north-faced door
		return getMirroredStringArray(room, true, false)
	}
	if conn[0] == 1 && conn[1] == 0 { // east-faced door
		return getRotatedStringArray(room)
	}
	if conn[0] == -1 && conn[1] == 0 { // west-faced door
		return getMirroredStringArray(getRotatedStringArray(room), false, true)
	}
	return nil
}

func getTwoConnRoom(conn [][]int) *[]string {
	north, east, south, west := getDirectionsByConnsArray(&conn)
	// first, determine whether the connections are symmetric
	if north && south || east && west {
		// they are symmetric, well yeah
		room := getRandomRoomFromArray(&ns_entrance_rooms)
		if north && south { // north and south
			return room
		} else { // west-east connection
			return getRotatedStringArray(room)
		}
	} else {
		room := getRandomRoomFromArray(&se_entrance_rooms)
		if south && east { // south-east
			return room
		}
		if north && east { // north-east
			return getMirroredStringArray(room, true, false)
		}
		if north && west { // north-west
			return getMirroredStringArray(getRotatedStringArray(room), false, true)
		}
		if south && west { // south-west
			return getMirroredStringArray(room, false, true)
		}
	}
	return nil
}

func placeDoorsToRoomByConnection(room *[]string, conns *[][]int) {
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
		newStr = newStr[:doory] + "+" + newStr[doory+1:]
		(*room)[doorx] = newStr
	}
}

func GetRoomByNodeConnections(conns *[][]int) *[]string {
	var room []string
	switch len(*conns) {
	case 0:
		room = *getRandomRoomFromArray(&no_entrance_rooms)
	case 1:
		room = *getSingleConnRoom((*conns)[0])
	case 2:
		room = *getTwoConnRoom(*conns)
	case 4:
		room = *getRandomRoomFromArray(&all_entrance_rooms)
	}
	// now we should outline the room with walls
	h := len((room)[0])
	wall_row := strings.Repeat("#", h+2)
	var outlined_room []string
	outlined_room = append(outlined_room, wall_row)
	for i := range room {
		outlined_room = append(outlined_room, "#"+room[i]+"#")
	}
	outlined_room = append(outlined_room, wall_row)
	placeDoorsToRoomByConnection(&outlined_room, conns)
	return &outlined_room
}
