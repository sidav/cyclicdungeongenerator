package layout_generation

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type PatternParser struct {
	currentSplitLine []string
}

func (pp *PatternParser) ParsePatternFile(filename string) *pattern {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	pat := pattern{}

	currLine := 1
	for scanner.Scan() {
		if currLine == 1 {
			pat.name = scanner.Text()
		} else {
			if scanner.Text() != "" {
				pat.instructions = append(pat.instructions, *pp.parseLineToInstruction(scanner.Text()))
			}
		}
		currLine++
	}
	return &pat
}

func (pp *PatternParser) parseLineToInstruction(line string) *patternStep {
	pp.currentSplitLine = strings.Split(line, " ")

	action := strings.ToUpper(pp.currentSplitLine[0])
	switch action {
	case "ADDROOMATEMPTY": // ADDROOMATEMPTY ROOMNAME room_name FX fx FY fy TX tx TY ty MINEMPTYCELLSNEAR minemptycellsnear
		return &patternStep{
			actionType:        ACTION_PLACE_NODE_AT_EMPTY,
			minEmptyCellsNear: pp.getIntAfterIdentifier("MINEMPTYCELLSNEAR"),
			nameOfNode:        pp.getStringAfterIdentifier("ROOMNAME"),
		}
	case "PLACEPATH": // PLACEPATH PATHID id FROM fromname TO toname
		return &patternStep{
			actionType:        ACTION_PLACE_PATH_FROM_TO,
			minEmptyCellsNear: pp.getIntAfterIdentifier("MINEMPTYCELLSNEAR"),
			nameFrom:          pp.getStringAfterIdentifier("FROM"),
			nameTo:            pp.getStringAfterIdentifier("TO"),
			pathNumber:        pp.getIntAfterIdentifier("PATHID"),
		}
	case "PLACERANDOMROOMS": // PLACERANDOMROOMS MIN minrooms MAX maxrooms
		return &patternStep{
			actionType:        ACTION_PLACE_RANDOM_CONNECTED_NODES,
			minEmptyCellsNear: pp.getIntAfterIdentifier("MINEMPTYCELLSNEAR"),
			nameFrom:          pp.getStringAfterIdentifier("FROM"),
			nameTo:            pp.getStringAfterIdentifier("TO"),
			countFrom:         pp.getIntAfterIdentifier("MIN"),
			countTo:           pp.getIntAfterIdentifier("MAX"),
			pathNumber:        pp.getIntAfterIdentifier("PATHID"),
		}
	case "PLACEROOMNEARPATH": // PLACEROOMNEARPATH PATHID id ROOMNAME newroomname
		return &patternStep{
			actionType:        ACTION_PLACE_NODE_NEAR_PATH,
			minEmptyCellsNear: pp.getIntAfterIdentifier("MINEMPTYCELLSNEAR"),
			nameOfNode:        pp.getStringAfterIdentifier("ROOMNAME"),
			nameFrom:          pp.getStringAfterIdentifier("FROM"),
			nameTo:            pp.getStringAfterIdentifier("TO"),
			countFrom:         pp.getIntAfterIdentifier("MIN"),
			countTo:           pp.getIntAfterIdentifier("MAX"),
			pathNumber:        pp.getIntAfterIdentifier("PATHID"),
		}
	case "LOCKROOMFROMPATH": // LOCKROOMFROMPATH PATHID id ROOMNAME roomname LOCKID lockid
		return &patternStep{
			actionType:        ACTION_SET_NODE_CONNECTION_LOCKED_FROM_PATH,
			minEmptyCellsNear: pp.getIntAfterIdentifier("MINEMPTYCELLSNEAR"),
			nameOfNode:        pp.getStringAfterIdentifier("ROOMNAME"),
			nameFrom:          pp.getStringAfterIdentifier("FROM"),
			nameTo:            pp.getStringAfterIdentifier("TO"),
			countFrom:         pp.getIntAfterIdentifier("MIN"),
			countTo:           pp.getIntAfterIdentifier("MAX"),
			pathNumber:        pp.getIntAfterIdentifier("PATHID"),
			lockNumber:        pp.getIntAfterIdentifier("LOCKID"),
		}
	case "ADDROOMSTATUS": // ADDROOMSTATUS ROOMNAME roomname STATUS status
		return &patternStep{
			actionType:        ACTION_SET_NODE_STATUS,
			minEmptyCellsNear: pp.getIntAfterIdentifier("MINEMPTYCELLSNEAR"),
			nameOfNode:        pp.getStringAfterIdentifier("ROOMNAME"),
			nameFrom:          pp.getStringAfterIdentifier("FROM"),
			nameTo:            pp.getStringAfterIdentifier("TO"),
			countFrom:         pp.getIntAfterIdentifier("MIN"),
			countTo:           pp.getIntAfterIdentifier("MAX"),
			pathNumber:        pp.getIntAfterIdentifier("PATHID"),
			lockNumber:        pp.getIntAfterIdentifier("LOCKID"),
			status:            pp.getStringAfterIdentifier("STATUS"),
		}
	}
	panic("UNKNOWN ACTION IDENTIFIER " + action)
}

func (pp *PatternParser) getStringAfterIdentifier(ident string) string {
	for i := range pp.currentSplitLine {
		if i > 0 && strings.ToUpper(pp.currentSplitLine[i-1]) == ident {
			return pp.currentSplitLine[i]
		}
	}
	return ""
}

func (pp *PatternParser) getIntAfterIdentifier(ident string) int {
	for i := range pp.currentSplitLine {
		if i > 0 && strings.ToUpper(pp.currentSplitLine[i-1]) == ident {
			val, err := strconv.Atoi(pp.currentSplitLine[i])
			if err != nil {
				panic("Broken!")
			}
			return val
		}
	}
	return 0
}
