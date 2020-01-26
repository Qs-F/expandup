package expandup

import "strings"

type Doc struct {
	s []string
}

func filter(s []string) (ret []string) {
	for _, v := range s {
		if v == "" {
			continue
		}
		ret = append(ret, v)
	}

	return ret
}

func Document(raw string) *Doc {
	d := &Doc{}
	d.s = strings.Split(raw, "\n")

	return d
}
