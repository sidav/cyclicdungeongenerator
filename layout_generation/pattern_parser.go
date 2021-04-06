package layout_generation

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type PatternParser struct {
	currentSplitLine   []string
	WriteLinesInResult bool
}

func (pp *PatternParser) ListPatternFilenamesInPath(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	var names []string
	for _, f := range files {
		if strings.Contains(f.Name(), ".ptn") {
			name := strings.Replace(path+"/"+f.Name(), "//", "/", -1)
			names = append(names, name)
		}
	}
	return names
}

func (pp *PatternParser) ParsePatternFile(filename string) *pattern {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	pat := pattern{Filename: filename}

	currLine := 1
	for scanner.Scan() {
		if currLine == 1 {
			pat.Name = scanner.Text()
		} else {
			if scanner.Text() != "" {
				newInstr := pp.parseLineToInstruction(scanner.Text())
				if newInstr != nil {
					if pp.WriteLinesInResult {
						newInstr.instructionText = fmt.Sprintf("%d: %s", currLine, scanner.Text())
					}
					pat.instructions = append(pat.instructions, newInstr)
				}
			}
		}
		currLine++
	}
	return &pat
}

func (pp *PatternParser) parseLineToInstruction(line string) *patternStep {
	pp.currentSplitLine = strings.Split(line, " ")

	action := strings.Replace(strings.ToUpper(pp.currentSplitLine[0]), "_", "", -1)
	if action[0] == "#"[0] { // that's a comment
		return nil
	}
	switch action {
	// ADDROOMATEMPTY ROOMNAME room_name FX fromX FY fromY TX toX TY toY MINEMPTYNEAR minemptycellsnear TAGS tags
	case "ADDROOMATEMPTY":
		return &patternStep{
			actionType:        ACTION_PLACE_NODE_AT_EMPTY,
			minEmptyCellsNear: pp.getIntAfterIdentifier("MINEMPTYNEAR"),
			nameOfNode:        pp.getStringAfterIdentifier("ROOMNAME"),
			fromX:             pp.getIntAfterIdentifier("FX"),
			fromY:             pp.getIntAfterIdentifier("FY"),
			toX:               pp.getIntAfterIdentifier("TX"),
			toY:               pp.getIntAfterIdentifier("TY"),
			tags:              pp.getStringAfterIdentifier("TAGS"),
		}
	case "PLACEPATH": // PLACEPATH PATHID id FROM fromname TO toname
		return &patternStep{
			actionType:        ACTION_PLACE_PATH_FROM_TO,
			minEmptyCellsNear: pp.getIntAfterIdentifier("MINEMPTYNEAR"),
			nameFrom:          pp.getStringAfterIdentifier("FROM"),
			nameTo:            pp.getStringAfterIdentifier("TO"),
			pathNumber:        pp.getIntAfterIdentifier("PATHID"),
			allowCrossPaths:   strings.ToUpper(pp.getStringAfterIdentifier("ALLOWINTERSECT")) == "TRUE",
		}
	case "PLACERANDOMROOMS": // PLACERANDOMROOMS MIN minrooms MAX maxrooms
		return &patternStep{
			actionType:        ACTION_PLACE_RANDOM_CONNECTED_NODES,
			minEmptyCellsNear: pp.getIntAfterIdentifier("MINEMPTYNEAR"),
			nameFrom:          pp.getStringAfterIdentifier("FROM"),
			nameTo:            pp.getStringAfterIdentifier("TO"),
			countFrom:         pp.getIntAfterIdentifier("MIN"),
			countTo:           pp.getIntAfterIdentifier("MAX"),
			pathNumber:        pp.getIntAfterIdentifier("PATHID"),
		}
	case "PLACEROOMNEARPATH": // PLACEROOMNEARPATH PATHID id ROOMNAME newroomname TAGS tag
		return &patternStep{
			actionType:        ACTION_PLACE_NODE_NEAR_PATH,
			minEmptyCellsNear: pp.getIntAfterIdentifier("MINEMPTYNEAR"),
			nameOfNode:        pp.getStringAfterIdentifier("ROOMNAME"),
			nameFrom:          pp.getStringAfterIdentifier("FROM"),
			nameTo:            pp.getStringAfterIdentifier("TO"),
			countFrom:         pp.getIntAfterIdentifier("MIN"),
			countTo:           pp.getIntAfterIdentifier("MAX"),
			pathNumber:        pp.getIntAfterIdentifier("PATHID"),
			tags:              pp.getStringAfterIdentifier("TAGS"),
		}
	case "PLACEROOMATPATH": // PLACEROOMATPATH PATHID id ROOMNAME newroomname TAGS tag
		return &patternStep{
			actionType:        ACTION_PLACE_NODE_AT_PATH,
			minEmptyCellsNear: pp.getIntAfterIdentifier("MINEMPTYNEAR"),
			nameOfNode:        pp.getStringAfterIdentifier("ROOMNAME"),
			nameFrom:          pp.getStringAfterIdentifier("FROM"),
			nameTo:            pp.getStringAfterIdentifier("TO"),
			countFrom:         pp.getIntAfterIdentifier("MIN"),
			countTo:           pp.getIntAfterIdentifier("MAX"),
			pathNumber:        pp.getIntAfterIdentifier("PATHID"),
			tags:              pp.getStringAfterIdentifier("TAGS"),
		}
	case "LOCKROOMFROMPATH": // LOCKROOMFROMPATH PATHID id ROOMNAME roomname LOCKID lockid
		return &patternStep{
			actionType:        ACTION_SET_NODE_CONNECTION_LOCKED_FROM_PATH,
			minEmptyCellsNear: pp.getIntAfterIdentifier("MINEMPTYNEAR"),
			nameOfNode:        pp.getStringAfterIdentifier("ROOMNAME"),
			nameFrom:          pp.getStringAfterIdentifier("FROM"),
			nameTo:            pp.getStringAfterIdentifier("TO"),
			countFrom:         pp.getIntAfterIdentifier("MIN"),
			countTo:           pp.getIntAfterIdentifier("MAX"),
			pathNumber:        pp.getIntAfterIdentifier("PATHID"),
			lockNumber:        pp.getIntAfterIdentifier("LOCKID"),
		}
	case "SETROOMTAGS": // SETROOMTAGS ROOMNAME roomname TAGS tags
		return &patternStep{
			actionType:        ACTION_SET_NODE_TAGS,
			minEmptyCellsNear: pp.getIntAfterIdentifier("MINEMPTYNEAR"),
			nameOfNode:        pp.getStringAfterIdentifier("ROOMNAME"),
			nameFrom:          pp.getStringAfterIdentifier("FROM"),
			nameTo:            pp.getStringAfterIdentifier("TO"),
			countFrom:         pp.getIntAfterIdentifier("MIN"),
			countTo:           pp.getIntAfterIdentifier("MAX"),
			pathNumber:        pp.getIntAfterIdentifier("PATHID"),
			lockNumber:        pp.getIntAfterIdentifier("LOCKID"),
			tags:              pp.getStringAfterIdentifier("TAGS"),
		}
	}
	panic("UNKNOWN ACTION IDENTIFIER " + action)
}

func (pp *PatternParser) getStringAfterIdentifier(ident string) string {
	for i := range pp.currentSplitLine {
		if i > 0 {
			str := strings.Replace(strings.ToUpper(pp.currentSplitLine[i-1]), "_", "", -1)
			if str == ident {
				return pp.currentSplitLine[i]
			}
		}
	}
	return ""
}

func (pp *PatternParser) getIntAfterIdentifier(ident string) int {
	for i := range pp.currentSplitLine {
		if i > 0 {
			str := strings.Replace(strings.ToUpper(pp.currentSplitLine[i-1]), "_", "", -1)
			if str == ident {
				val, err := strconv.Atoi(pp.currentSplitLine[i])
				if err != nil {
					panic("Broken!")
				}
				return val
			}
		}
	}
	return 0
}
