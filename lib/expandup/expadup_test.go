package expandup

import "strings"

type Config struct {
	OpenPrefix  string
	OpenSuffix  string
	ClosePrefix string
	CloseSuffix string

	IgnoreIndents bool
}

type Document struct {
	s []string
}

func ToDocument(raw string) *Document {
	d := &Document{}
	d.s = strings.Split(raw, "\n")

	return d
}
