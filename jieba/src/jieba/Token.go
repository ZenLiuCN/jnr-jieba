package jieba

type Token struct {
	Text string
	Frq  float64
	Pos  string
}

func NewToken(text string, frq float64, pos string) Token {
	return Token{
		Text: text,
		Frq:  frq,
		Pos:  pos,
	}
}
