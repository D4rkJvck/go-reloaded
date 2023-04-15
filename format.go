package main

import (
	"strings"
	"unicode"
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

func FixApostrophic(s string) string {
	open := true

	if len(s) > 1 && s[0] == 39 {
		open = false
		if s[1] == 32 {
			s = Delete(s, 1)
		}
	}

	for i := range s {
		if i > 0 && i < len(s)-1 {
			if s[i] == 39 && (!unicode.IsLetter(rune(s[i-1])) || !unicode.IsLetter(rune(s[i+1]))) {
				if open {
					if s[i-1] != 32 {
						s = AddSpace(s, i)
					}
					if s[i+1] == 32 {
						s = Delete(s, i+1)
					}
					open = false
				} else {
					if s[i+1] != 32 && !isPunctuation(s[i+1]) {
						s = AddSpace(s, i+1)
					}
					if s[i-1] == 32 {
						s = Delete(s, i-1)
					}
					open = true
				}
			}
		}
	}
	if len(s) > 1 && !open && s[len(s)-1] == 39 && s[len(s)-2] == 32 {
		s = Delete(s, len(s)-2)
	}
	return s
}


func FixAn(s string) string {
	t := strings.Fields(s)
	for i, word := range t {
		if i < len(t)-1 {
			if word == "a" || word == "A" {
				if isVowel(t[i+1][0]) {
					t[i] += "n"
				}
			}
		}
	}
	return strings.Join(t, " ")
}