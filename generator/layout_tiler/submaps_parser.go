package layout_tiler

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func (ltl *LayoutTiler) parseSubmapsDir(path string) {
	if path != "" {
		ltl.parseSubmapsDirRecursively(path, "")
	}
}

func (ltl *LayoutTiler) parseSubmapsDirRecursively(path, tag string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return
		log.Fatal(err)
	}
	for _, f := range files {
		name := strings.Replace(path+"/"+f.Name(), "//", "/", -1)
		if f.IsDir() {
			ltl.parseSubmapsDirRecursively(path+"/"+f.Name(), f.Name())
		} else {
			ltl.parseSubmapFile(name, tag)
		}
	}
}

func (ltl *LayoutTiler) parseSubmapFile(filename, tag string) {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sm := submap{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			ltl.submaps[tag] = append(ltl.submaps[tag], sm)
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
	ltl.submaps[tag] = append(ltl.submaps[tag], sm)
}
