package layout_to_tiles2

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type submap struct {
	chars     [][]rune
	usedTimes int
}

func (ltl *LayoutToLevel) parseSubmapsDir(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		name := strings.Replace(path+"/"+f.Name(), "//", "/", -1)
		ltl.parseSubmapFile(name)
	}
}

func (ltl *LayoutToLevel) parseSubmapFile(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sm := submap{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			ltl.submaps = append(ltl.submaps, sm)
			sm = submap{}
			continue
		}
		line := scanner.Text()
		var newArr []rune
		for _, chr := range line {
			if chr == '.' {
				chr = ' '
			}
			newArr = append(newArr, rune(chr))
		}
		sm.chars = append(sm.chars, newArr)
	}
	ltl.submaps = append(ltl.submaps, sm)
}

func (ltl *LayoutToLevel) applySubmaps() {
	for i := range ltl.submaps {
		ltl.applySubmapAtRandom(&ltl.submaps[i])
	}
}

func (ltl *LayoutToLevel) applySubmapAtRandom(sm *submap) {
	smH, smW := len(sm.chars), len(sm.chars[0])
	applicableCoords := make([][2]int, 0)
	for x := 0; x < len(ltl.charmap)-smW; x++ {
		for y := 0; y < len(ltl.charmap[x])-smH; y++ {
			if ltl.isSpaceEmpty(x, y, smW, smH) {
				applicableCoords = append(applicableCoords, [2]int{x, y})
			}
		}
	}
	if len(applicableCoords) > 0 {
		randCoordsIndex := ltl.rnd.Rand(len(applicableCoords))
		ltl.applySubmapAtCoords(sm, applicableCoords[randCoordsIndex][0], applicableCoords[randCoordsIndex][1])
	}
}


func (ltl *LayoutToLevel) applySubmapAtCoords(sm *submap, xx, yy int) bool {
	smH, smW := len(sm.chars), len(sm.chars[0])
	for x := 0; x < smW; x++ {
		for y := 0; y < smH; y++ {
			ltl.charmap[xx+x][yy+y] = sm.chars[y][x]
		}
	}
	return true
}

func (ltl *LayoutToLevel) isSpaceEmpty(xx, yy, w, h int) bool {
	for x := xx; x < xx+w; x++ {
		for y := yy; y < yy+h; y++ {
			if ltl.charmap[x][y] != ' ' {
				return false
			}
		}
	}
	return true
}
