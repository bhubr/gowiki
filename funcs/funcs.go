package funcs

import (
	"regexp"
)

func makeLink(bracketTitle []byte) []byte {
	title := string(bracketTitle[1 : len(bracketTitle)-1])
	str := "<a href=\"/view/" + title + "\">" + title + "</a>"
	return []byte(str)
}

func ReplaceLinks(s string) string {
	r, _ := regexp.Compile("\\[([a-zA-Z0-9]+)\\]")
	return string(r.ReplaceAllFunc([]byte(s), makeLink))
}
