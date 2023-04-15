package main

import (
	"strconv"
	"strings"
)

func FixSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func isPunctuation(b byte) bool {
	p := ",.;:?!"
	for range p {
		if strings.Contains(p, string(b)) {
			return true
		}
	}
	return false
}

func AddSpace(s string, i int) string {
	var newStr strings.Builder
	newStr.WriteString(s[:i])
	newStr.WriteRune(32)
	newStr.WriteString(s[i:])
	return newStr.String()
}

func Delete(s string, i int) string {
	var newStr strings.Builder
	newStr.WriteString(s[:i])
	newStr.WriteString(s[i+1:])
	return newStr.String()
}

func HexToDec(s string) string {
	if isHexadecimal(s) {
		hex, err := strconv.ParseInt(s, 16, 64)
		if err == nil {
			s = strconv.Itoa(int(hex))
		}
	}
	return s
}

func isHexadecimal(s string) bool {
	hex := "0123456789AaBbCcDdEeFf"
	for _, c := range s {
		if !strings.Contains(hex, string(c)) {
			return false
		}
	}
	return true
}

func BinToDec(s string) string {
	if isBinary(s) {
		bin, err := strconv.ParseInt(s, 2, 64)
		if err == nil {
			s = strconv.Itoa(int(bin))
		}
	}
	return s
}

func isBinary(s string) bool {
	for _, c := range s {
		if c != 48 && c != 49 {
			return false
		}
	}
	return true
}

func Capital(s string) string { //---					---> Deprecated but functional
	return strings.Title(strings.ToLower(s))
}

func isVowel(b byte) bool {
	v := "AaEeHhIiOoUu"
	return strings.Contains(v, string(b))
}
