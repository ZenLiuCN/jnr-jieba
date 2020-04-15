package global

import "fmt"

type ProbState struct {
	Prob  float64
	State byte
}

func (p ProbState) String() string {
	return fmt.Sprintf("(%f, %x)", p.Prob, p.State)
}

type ProbStates []*ProbState

//Len length of states
func (ps ProbStates) Len() int {
	return len(ps)
}

//Less compare states
func (ps ProbStates) Less(i, j int) bool {
	if ps[i].Prob == ps[j].Prob {
		return ps[i].State < ps[j].State
	}
	return ps[i].Prob < ps[j].Prob
}

//Swap swap states
func (ps ProbStates) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
