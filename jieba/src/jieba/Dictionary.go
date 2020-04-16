package jieba

import (
	"bufio"
	"encoding/gob"
	"io"
	"math"
	"strconv"
	"strings"
	"sync"
)

type Dictionary struct {
	total, logTotal float64
	freqMap         map[string]float64
	sync.RWMutex
}

func (d *Dictionary) AddToken(token Token) {
	d.Lock()
	d.freqMap[token.Text] = token.Frq
	d.total += token.Frq
	runes := []rune(token.Text)
	n := len(runes)
	for i := 0; i < n; i++ { //TODO: n-1?
		frag := string(runes[:i+1])
		if _, ok := d.freqMap[frag]; !ok {
			d.freqMap[frag] = 0.0
		}
	}
	d.Unlock()
	d.logTotal = math.Log(d.total)
}
func (d *Dictionary) Frequency(key string) (float64, bool) {
	d.RLock()
	freq, ok := d.freqMap[key]
	d.RUnlock()
	return freq, ok
}
func (d *Dictionary) LoadFrom(read DictionaryReader) error {
	t, e := read.Load()
	d.Lock()
	defer d.Unlock()
	go func(tch <-chan Token) {
		for token := range tch {
			d.AddToken(token)
		}
	}(t)
	return <-e
}
func (d *Dictionary) LoadFromWithOutLock(read DictionaryReader) error {
	t, e := read.Load()
	go func(tch <-chan Token) {
		for token := range tch {
			d.AddToken(token)
		}
	}(t)
	return <-e
}


type DictionaryReader interface {
	Load() (<-chan Token, <-chan error)
}

type TextDictionaryLoader struct {
	io.Reader
}
func (l TextDictionaryLoader) Load()(<-chan Token, <-chan error){
	t:=make(chan Token )
	e:=make(chan error )
	go func() {
		defer func() {
			close(t)
			close(e)
		}()
		var token Token
		var line string
		var fields []string
		var err error
		scanner := bufio.NewScanner(l)
		for scanner.Scan() {
			line = scanner.Text()
			fields = strings.Split(line, " ")
			token.Text = strings.TrimSpace(strings.Replace(fields[0], "\ufeff", "", 1))
			if length := len(fields); length > 1 {
				token.Frq, err = strconv.ParseFloat(fields[1], 64)
				if err != nil {
					e <- err
					return
				}
				if length > 2 {
					token.Pos = strings.TrimSpace(fields[2])
				}
			}
			t <- token
		}
		if err = scanner.Err(); err != nil {
			e <- err
		}
	}()
	return t,e
}


type GobDictionaryLoader struct {
	*gob.Decoder
}
func (l GobDictionaryLoader) Load()(<-chan Token, <-chan error){
	t:=make(chan Token )
	e:=make(chan error )
	go func() {
		defer func() {
			close(t)
			close(e)
		}()
		var token *Token
		var err error
		for{
			token=new(Token)
			err=l.Decode(token)
			if err!=nil{
				e<-err
				return
			}
			t<-*token
		}

	}()
	return t,e
}