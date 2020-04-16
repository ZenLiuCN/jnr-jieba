package finalseg

import (
	."jiebaJnr/jieba/global"
	"regexp"
	"sort"
)

const minFloat = -3.14e100

var (
	prevStatus = map[byte][]byte{
		'B': {'E', 'S'},
		'M': {'M', 'B'},
		'S': {'S', 'E'},
		'E': {'B', 'M'},
	}
)



func viterbi(obs []rune, States []byte) (float64, []byte) {
	path := make(map[byte][]byte)
	V := make([]map[byte]float64, len(obs))
	V[0] = make(map[byte]float64)
	for _, y := range States {
		if val, ok := probEmit[y][obs[0]]; ok {
			V[0][y] = val + probStart[y]
		} else {
			V[0][y] = minFloat + probStart[y]
		}
		path[y] = []byte{y}
	}

	for t := 1; t < len(obs); t++ {
		newPath := make(map[byte][]byte)
		V[t] = make(map[byte]float64)
		for _, y := range States {
			ps0 := make(ProbStates, 0)
			var emP float64
			if val, ok := probEmit[y][obs[t]]; ok {
				emP = val
			} else {
				emP = minFloat
			}
			for _, y0 := range prevStatus[y] {
				var transP float64
				if tp, ok := probTrans[y0][y]; ok {
					transP = tp
				} else {
					transP = minFloat
				}
				prob0 := V[t-1][y0] + transP + emP
				ps0 = append(ps0, &ProbState{Prob: prob0, State: y0})
			}
			sort.Sort(sort.Reverse(ps0))
			V[t][y] = ps0[0].Prob
			pp := make([]byte, len(path[ps0[0].State]))
			copy(pp, path[ps0[0].State])
			newPath[y] = append(pp, y)
		}
		path = newPath
	}
	ps := make(ProbStates, 0)
	for _, y := range []byte{'E', 'S'} {
		ps = append(ps, &ProbState{V[len(obs)-1][y], y})
	}
	sort.Sort(sort.Reverse(ps))
	v := ps[0]
	return v.Prob, path[v.State]
}

var (
	reHan  = regexp.MustCompile(`\p{Han}+`)
	reSkip = regexp.MustCompile(`(\d+\.\d+|[a-zA-Z0-9]+)`)
)

func cutHan(sentence string) chan string {
	result := make(chan string)
	go func() {
		runes := []rune(sentence)
		_, posList := viterbi(runes, []byte{'B', 'M', 'E', 'S'})
		begin, next := 0, 0
		for i, char := range runes {
			pos := posList[i]
			switch pos {
			case 'B':
				begin = i
			case 'E':
				result <- string(runes[begin : i+1])
				next = i + 1
			case 'S':
				result <- string(char)
				next = i + 1
			}
		}
		if next < len(runes) {
			result <- string(runes[next:])
		}
		close(result)
	}()
	return result
}

func Cut(sentence string) chan string {
	result := make(chan string)
	s := sentence
	var hans string
	var hanLoc []int
	var noneHanLoc []int
	go func() {
		for {
			//loop for index of Hans
			hanLoc = reHan.FindStringIndex(s)
			if hanLoc == nil {
				if len(s) == 0 {
					break
				}
			} else if hanLoc[0] == 0 {
				hans = s[hanLoc[0]:hanLoc[1]]
				s = s[hanLoc[1]:]
				for han := range cutHan(hans) {
					// force split words
					if _, ok := ForceSplitWords[han]; !ok {
						result <- han
					} else {
						for _,i2 := range []rune(han) {
							result <- string(i2)
						}
					}

				}
				continue
			}
			noneHanLoc = reSkip.FindStringIndex(s)
			if noneHanLoc == nil {
				if len(s) == 0 {
					break
				}
			} else if noneHanLoc[0] == 0 {
				nonHans := s[noneHanLoc[0]:noneHanLoc[1]]
				s = s[noneHanLoc[1]:]
				if nonHans != "" {
					result <- nonHans
					continue
				}
			}
			var loc []int
			if hanLoc == nil && noneHanLoc == nil {
				if len(s) > 0 {
					result <- s
					break
				}
			} else if hanLoc == nil {
				loc = noneHanLoc
			} else if noneHanLoc == nil {
				loc = hanLoc
			} else if hanLoc[0] < noneHanLoc[0] {
				loc = hanLoc
			} else {
				loc = noneHanLoc
			}
			result <- s[:loc[0]]
			s = s[loc[0]:]
		}
		close(result)
	}()
	return result
}
