package main

import (
	"bufio"
	"compress/gzip"
	"jiebagou/jieba"

	"os"
	"testing"
)


func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		split(seg, false)
	}
}
func write(){
	fi, e := os.Open("dict.txt")
	if e != nil {
		panic(e)
	}
	defer fi.Close()
	fo, e := os.OpenFile("dict.gz", os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		panic(e)
	}
	defer fo.Close()
	en, e := gzip.NewWriterLevel(fo, gzip.BestCompression)
	if e != nil {
		panic(e)
	}
	defer en.Close()
	sc := bufio.NewScanner(fi)
	for sc.Scan() {
		text := sc.Text()
		if _,e = en.Write([]byte(text+"\n")); e != nil {
			panic(e)
		}

	}
	if e = sc.Err(); e != nil {
		panic(e)
	}
}

func read(){
	fi, e := os.Open("dict.gz")
	if e != nil {
		panic(e)
	}
	defer fi.Close()
	en, e := gzip.NewReader(fi)
	if e != nil {
		panic(e)
	}
	sc := bufio.NewScanner(en)
	/*	cbuffer := make([]byte, 0, bufio.MaxScanTokenSize)
		sc.Buffer(cbuffer, bufio.MaxScanTokenSize*50)*/
	for sc.Scan() {
		text := sc.Text()
		println(text)

	}
	if e = sc.Err(); e != nil {
		panic(e)
	}
}

func load() *jieba.Segmenter{
	fi, e := os.Open("dict.gz")
	if e != nil {
		panic(e)
	}
	defer fi.Close()
	en, e := gzip.NewReader(fi)
	if e != nil {
		panic(e)
	}
	loader:=jieba.TextDictionaryLoader{en}
	seg:=new(jieba.Segmenter)
	e=seg.LoadDictionaryWithOutLock(loader)
	if e != nil {
		panic(e)
	}
	return seg

}
func split(seg *jieba.Segmenter,pnt bool){
	words:=seg.CutAll(`接触自然语言处理(NLP)有段时间，理论知识有些了解，挺想动手写些东西，想想开源界关于NLP的东西肯定不少，其中分词是NLP的基础，遂在网上找了些资源，其中结巴分词是国内程序员用python开发的一个中文分词模块, 源码已托管在github: 源码地址 ，代码用python实现，源码中也有注释，但一些细节并没有相应文档，因此这里打算对源码进行分析，一来把知识分享，让更多的童鞋更快的对源码有个认识，二来使自己对分词这一块有个更深入的理解。`)
	for  i2 := range words {
		if pnt{
			println(i2)
		}
	}
}