package core

import (
	"regexp"
)

func OutfileName(s string) (replaced string) {
	reg := regexp.MustCompile("\\.tmpl$")
	replaced = reg.ReplaceAllLiteralString(s, "")
	if s == replaced {
		replaced += ".out"
	}
	return replaced
}
