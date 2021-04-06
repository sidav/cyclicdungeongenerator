package layout_to_tiled_map

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

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
