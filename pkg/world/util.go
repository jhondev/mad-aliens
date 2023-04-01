package world

import "strings"

// TODO: make this join generic
func JoinAliens(aliens AliensStat, sep string) string {
	if len(aliens) == 0 {
		return ""
	}

	buf := strings.Builder{}
	maxLen := 0
	i := 0

	for k := range aliens {
		if len(k) > maxLen {
			maxLen = len(k)
		}
		if i > 0 {
			buf.WriteString(sep)
		}
		buf.Grow(buf.Len() + len(k) + len(sep))
		buf.WriteString(string(k))
		i++
	}

	return buf.String()
}
