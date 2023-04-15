package main

import (
	"strings"
)

func FixPunctuation(s string) string {
	if len(s) > 1 && isPunctuation(s[0]) {
		if s[1] != 32 && !isPunctuation(s[1]) {
			s = AddSpace(s, 1)
		}
	}
	for i := range s {
		if i > 0 && i < len(s)-1 {
			if isPunctuation(s[i]) {
				if s[i+1] != 32 && !isPunctuation(s[i+1]) {
					s = AddSpace(s, i+1)
				}
				if s[i-1] == 32 {
					s = Delete(s, i-1)
				}
			}
		}
	}
	if len(s) > 1 && isPunctuation(s[len(s)-1]) && s[len(s)-2] == 32 {
		s = Delete(s, len(s)-2)
	}
	return FixSpaces(s)
}

func FixParentheses(s string) string {
	out := strings.NewReplacer("(", " (", ")", ") ")
	in := strings.NewReplacer("( ", "(", " )", ")")
	s = out.Replace(s)
	s = in.Replace(in.Replace(s))
	return FixSpaces(s)
}
