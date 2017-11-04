package expandup

type Line struct {
	Number  int
	Content string
}

type Block []*Line

// func (ln *Line) NextLine(blc *Block, current *Line) *Line {
// 	if len(blc) > current.Number+1 {
// 		return
// 	}
// }
//
// func (ln *Line) IsStartMarker() bool {
// 	s := strings.TrimSpace(ln.Content)
// 	if strings.HasPrefix(s, START_MARKER_PREFIX) && strings.HasSuffix(s, START_MARKER_SUFFIX) {
// 		return true
// 	}
// 	return false
// }
//
// func (ln *Line) IsEndMarker() bool {
// 	s := strings.TrimSpace(ln.Content)
// 	if s == END_MARKER {
// 		return true
// 	}
// 	return false
// }
