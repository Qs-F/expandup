package expandup

import "crypto/sha1"

type Def struct {
	Ident   string
	Content string
	SHA1    SHA1
}

func NewDef(ident string, content string) *Def {
	return &Def{
		Ident:   ident,
		Content: content,
		SHA1:    Sum(content),
	}
}

type SHA1 [sha1.Size]byte

func Sum(src string) SHA1 {
	return SHA1(sha1.Sum([]byte(src)))
}

func Comp(s1, s2 SHA1) bool {
	return s1 == s2
}

type Config struct {
	OpenPrefix  string
	OpenSuffix  string
	ClosePrefix string
	CloseSuffix string

	IgnoreIndents bool
}
