package global

type PosTuple struct {
	B     byte
	S string
}


type PosTuples []*PosTuple

//Len length of states
func (ps PosTuples) Len() int {
	return len(ps)
}

//Less compare states
func (ps PosTuples) Less(i, j int) bool {
	if ps[i].B == ps[j].B {
		return ps[i].S < ps[j].S
	}
	return ps[i].B < ps[j].B
}

//Swap swap states
func (ps PosTuples) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
