package expandup

import "strings"

type RawDoc []*Line

func parseDoc(b []byte) *RawDoc {
	s := strings.Split(string(b), "\n")
	rd := &RawDoc{}
	for i, v := range s {
		*rd = append(*rd, &Line{
			Number:  i,
			Content: v,
		})
	}
	return rd
}
