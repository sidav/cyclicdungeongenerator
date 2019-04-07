package layout_to_tiled

import "DoomSlayeRL/routines"

// This contains rooms maps for the CDG's nodes to be translated into tiled rooms.

var single_entrance_rooms = [][]string{ // all entrances should be at bottom
	{
		"##########",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"####+#####",
	},
	{
		"##########",
		"#........#",
		"#........#",
		"#....#####",
		"#........#",
		"#####....#",
		"#........#",
		"#....#####",
		"#........#",
		"####+#####",
	},
}

var ns_entrance_rooms = [][]string {
	{
		"####+#####",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"####+#####",
	},
	{
		"####+#####",
		"#........#",
		"#........#",
		"#....#####",
		"#........#",
		"######...#",
		"#........#",
		"#....#####",
		"#........#",
		"####+#####",
	},
}

var se_entrance_rooms = [][]string {
	{
		"##########",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"#........+",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"####+#####",
	},
	{
		"##########",
		"#........#",
		"#........#",
		"#....#####",
		"#........#",
		"######...+",
		"#........#",
		"#....#####",
		"#........#",
		"#........#",
		"####+#####",
	},
}

var nes_entrance_rooms = [][]string {
	{
		"####+#####",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"#........+",
		"#........#",
		"#........#",
		"#........#",
		"####+#####",
	},
	{
		"####+#####",
		"#........#",
		"#........#",
		"#....#####",
		"#........#",
		"######...+",
		"#........#",
		"#....#####",
		"#........#",
		"####+#####",
	},
}

var all_entrance_rooms = [][]string {
	{
		"####+#####",
		"#........#",
		"#........#",
		"#........#",
		"#........#",
		"+........+",
		"#........#",
		"#........#",
		"#........#",
		"####+#####",
	},
	{
		"####+#####",
		"#........#",
		"#....#####",
		"#........#",
		"#####....#",
		"+........+",
		"####.....#",
		"#.....####",
		"#........#",
		"####+#####",
	},
}

func reverseString(s string) (result string) {
	for _,v := range s {
		result = string(v) + result
	}
	return
}

func getRotatedStringArray(arr *[]string) *[]string {
	return nil 
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
		for i := 0; i < len(*arr) ; i--{
			newArr = append(newArr, reverseString((*arr)[i]))
		}
		return &newArr
	}
	return nil
}

func getSingleConnRoom(conn []int) *[]string {
	var room = &single_entrance_rooms[routines.Random(len(single_entrance_rooms))]
	if conn[0] == 0 && conn[1] == 1 { // south-faced door
		return room
	}
	if conn[0] == 0 && conn[1] == -1 { // north-faced door
		return getMirroredStringArray(room, true, false)
	}
	if conn[0] == 1 && conn[1] == 0 { // east-faced door
		return room
	}
	if conn[0] == -1 && conn[1] == 0 { // west-faced door
		return room
	}
	return nil
}

func GetRoomByNodeConnections(conns *[][]int) *[]string {
	var room *[]string
	if len(*conns) == 1 {
		room = getSingleConnRoom((*conns)[0])
	}
	return
}

