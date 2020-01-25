package mysql

import "strings"

var keyWords = map[string]bool{
	"CONSTRAINT": true,
}

func removeTrailingComma(s string) string {
	if len(s) == 0 {
		return s
	}
	if len(s) == 1 && s == "," {
		return ""
	}
	s = strings.TrimSpace(s)
	for s[len(s)-1] == ',' {
		s = s[:len(s)-1]
	}
	return s
}

func removeKeywords(s string) string {
	out := []string{}
	for _, line := range strings.Split(s, "\n") {
		fields := strings.Fields(line)
		if keyWords[fields[0]] {
			out[len(out)-1] = removeTrailingComma(out[len(out)-1])
			continue
		}
		out = append(out, line)
	}
	return strings.Join(out, "\n")
}
