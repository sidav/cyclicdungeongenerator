package layout_to_tiled

import "DoomSlayeRL/routines"

func getRandomRoomFromArray(arr *[][]string) *[]string {
	return &((*arr)[routines.Random(len(*arr))]) // ow that's quite of some pointer magic!
}

func reverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func getRotatedStringArray(arr *[]string) *[]string { // rotates 90 degrees (counter)clockwise TODO: find it out
	newArr := make([]string, 0)
	for i:=0; i < len((*arr)[0]);i++ {
		str := ""
		for j:=0; j < len(*arr); j++ {
			str += string((*arr)[j][i])
		}
		newArr = append(newArr, str)
	}
	return &newArr
}

func getMirroredStringArray(arr *[]string, v, h bool) *[]string {
	newArr := make([]string, 0)
	if v && h {
		for i := len(*arr) - 1; i >= 0; i--{
			newArr = append(newArr, reverseString((*arr)[i]))
		}
		return &newArr
	}
	if v {
		for i := len(*arr) - 1; i >= 0; i--{
			newArr = append(newArr, (*arr)[i])
		}
		return &newArr
	}
	if h {
		for i := 0; i < len(*arr) ; i++{
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

func getTwoConnRoom(conn[][]int) *[]string {
	// first, determine whether the connections are symmetric
	if conn[0][0] == 0 && conn[1][0] == 0 || conn[0][1] == 0 && conn[1][1] == 0 {
		// they are symmetric, well yeah
		room := getRandomRoomFromArray(&ns_entrance_rooms)
		if conn[0][0] == 0 && conn[1][0] == 0 { // north and south
			return room
		} else { // west-east connection
			return getRotatedStringArray(room)
		}
	} else {
		room := getRandomRoomFromArray(&se_entrance_rooms)
		if conn[0][1] == 1 && conn[1][0] == 1 { // south-east
			return room
		}
		if conn[0][1] == -1 && conn[1][0] == 1 { // north-east
			return getMirroredStringArray(room, true, false)
		}
		if conn[0][1] == -1 && conn[1][0] == -1 { // south-east
			return getRotatedStringArray(getMirroredStringArray(room, true, false))
		}
		if conn[0][1] == 1 && conn[1][0] == -1 { // south-east
			return getMirroredStringArray(room, false, true)
		}
	}
	return nil
}

func GetRoomByNodeConnections(conns *[][]int) *[]string {
	switch len(*conns) {
	case 1:
		return getSingleConnRoom((*conns)[0])
	case 2:
		return getTwoConnRoom(*conns)
	}
	return getRandomRoomFromArray(&no_entrance_rooms)
}

