package jieba

import "regexp"

const (
	MIN_FLOAT = -3.14e100
)
var(
	PrevStatus=map[byte]string{
		'B': "ES",
		'M': "MB",
		'S': "SE",
		'E': "BM",
	}
)
var (
	reHan  = regexp.MustCompile(`\p{Han}+`)
	reSkip = regexp.MustCompile(`(\d+\.\d+|[a-zA-Z0-9]+)`)
)
