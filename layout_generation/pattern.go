package layout_generation

type pattern struct {
	Name         string
	Filename     string
	instructions []*patternStep
}

func (p *pattern) getTotalConnectionsForNodeWithName(name string) int {
	conns := 0
	for i := range p.instructions {
		aType := p.instructions[i].actionType
		if aType == ACTION_PLACE_PATH_FROM_TO {
			if p.instructions[i].nameFrom == name || p.instructions[i].nameTo == name {
				conns++
			}
		}
	}
	return conns
}

func (p *pattern) getTotalNodesToBePlacedAtPath(pathId int) int {
	nodes := 0
	for i := range p.instructions {
		aType := p.instructions[i].actionType
		if aType == ACTION_PLACE_NODE_AT_PATH && p.instructions[i].pathNumber == pathId {
			nodes++
		}
	}
	return nodes
}

func (p *pattern) getAllMinDistancesForNode(nodeName string) *map[string]int {
	nodedistmap := make(map[string]int, 0)
	for i := range p.instructions {
		aType := p.instructions[i].actionType
		if aType == ACTION_PLACE_PATH_FROM_TO {
			if p.instructions[i].pathNumber == 0 {
				continue
			}
			if p.instructions[i].nameFrom == nodeName || p.instructions[i].nameTo == nodeName {
				nodeForDist := p.instructions[i].nameFrom
				if nodeForDist == nodeName {
					nodeForDist = p.instructions[i].nameTo
				}
				nodedistmap[nodeForDist] = p.getTotalNodesToBePlacedAtPath(p.instructions[i].pathNumber) + 1
			}
		}
	}
	return &nodedistmap
}

//func (p *pattern) getAllNonzeroPathIdsForNodeWithName(name string) []int {
//	ids := make([]int, 0)
//	for i := range p.instructions {
//		aType := p.instructions[i].actionType
//		if aType == ACTION_PLACE_PATH_FROM_TO {
//			if p.instructions[i].nameFrom == name || p.instructions[i].nameTo == name && p.instructions[i].pathNumber != 0 {
//				ids = append(ids, p.instructions[i].pathNumber)
//			}
//		}
//	}
//	return ids
//}
