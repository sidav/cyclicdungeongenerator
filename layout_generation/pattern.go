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
