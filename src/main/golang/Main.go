package main

import "C"
import (
	"compress/gzip"
	"jiebaJnr/jieba"
	"os"

	"strings"
)

func main() {

}

var (
	seg *jieba.Segmenter
)

type TokenizerMode int

const (
	CUT_ALL TokenizerMode = iota
	CUT_HMM
	CUT_SEARCH_HMM
	CUT_SEARCH_NO_HMM
	CUT_NO_HMM
)
//export Initial
func Initial(fileName *C.char) {
	src:=C.GoString(fileName)
	if len(src)==0{
		src="dict.gz"
	}
	fi, e := os.Open(src)
	if e != nil {
		panic(e)
	}
	defer fi.Close()
	en, e := gzip.NewReader(fi)
	if e != nil {
		panic(e)
	}
	loader := jieba.TextDictionaryLoader{en}
	seg = new(jieba.Segmenter)
	e = seg.LoadDictionaryWithOutLock(loader)
	if e != nil {
		panic(e)
	}
}

//export Tokenizer
func Tokenizer(src *C.char, join *C.char, mode int) *C.char {
	srg := C.GoString(src)
	buf := new(strings.Builder)
	var ch <-chan string
	switch TokenizerMode(mode) {
	case CUT_HMM:
		ch = seg.Cut(srg, true)
	case CUT_SEARCH_HMM:
		ch = seg.CutForSearch(srg, true)
	case CUT_SEARCH_NO_HMM:
		ch = seg.CutForSearch(srg, false)
	case CUT_NO_HMM:
		ch = seg.Cut(srg, true)
	default:
		ch = seg.CutAll(srg)
	}
	for r := range ch {
		buf.WriteString(r)
		buf.WriteString(C.GoString(join))
	}
	return C.CString(buf.String())
}

//export AddWord
func AddWord(word *C.char,freq C.double){
	w:=C.GoString(word)
	seg.AddWord(w,float64(freq))
}
//export RemoveWord
func RemoveWord(word *C.char){
	w:=C.GoString(word)
	seg.DeleteWord(w)
}
