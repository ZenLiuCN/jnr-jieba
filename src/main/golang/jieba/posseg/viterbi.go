package posseg

import (
	."jiebaJnr/jieba/global"
	"math"
	"sort"
)

const(
	minFloat=-3.14e100
)
var(
	minInf=math.Inf(-1)
	allState []PosTuple
)

func init()  {
	for tuple, _ := range probTrans {
		allState= append(allState,tuple )
	}
}

func viterbi(obs []rune) (float64, []byte) {
	path := make(map[int]map[*PosTuple]byte)
	path[0]=make(map[*PosTuple]byte)
	V := make([]map[*PosTuple]float64, len(obs))
	V[0] = make(map[*PosTuple]float64)
	for _, tuple := range charStateTab[obs[0]] {
		V[0][tuple]=probStart[tuple]+probEmit[tuple][obs[0]]
		path[0][tuple]=0
	}
	for i, ob := range obs[1:] {
		path[i]=make(map[*PosTuple]byte)
		V[i] = make(map[*PosTuple]float64)
		preStates := make(PosTuples, 0)
		for tuple, _ := range path[i-1] {
			if len(probTrans[tuple])>0{
				preStates= append(preStates, tuple)
			}
		}
		preStateExpectNext:=make(PosTuples,0)
		for _, tuple := range preStates {
			for posTuple, m := range probTrans {
				if posTuple==tuple{
					for p, _ := range m {
						preStateExpectNext= append(preStateExpectNext, p)
					}
				}
			}
		}
	}
	return v.Prob, path[v.State]
}

