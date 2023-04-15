package main

import (
	"os"
)

func main() {
	Args := os.Args[1:]
	if len(Args) > 1 {
		content, err := os.ReadFile(Args[0])
		if err == nil {
			txt := string(content)
			txt = FixSpaces(txt)
			txt = FixPunctuation(txt)
			txt = FixParentheses(txt)
			txt = ApplyInstances(txt)
			txt = FixPunctuation(txt)
			// txt = FixApostrophic(txt)
			// txt = FixAn(txt)
			content = []byte(txt)

			os.WriteFile(Args[1], content, 0700)
		}
	}
}
