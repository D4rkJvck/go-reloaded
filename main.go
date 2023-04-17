package main

import "os"

func main() {
	Args := os.Args[1:]
	if len(Args) > 1 {
		content, err := os.ReadFile(Args[0]) //---					---> Access File Content from File name
		if err == nil {

			txt := string(content) //---					---> Convert for Fortmatting

			//---> Format to Correct Instances
			txt = FixSpaces(txt)
			txt = FixPunctuation(txt)
			txt = FixParentheses(txt)

			txt = ApplyInstances(FixAn(txt)) //---					---> Apply then Delete (Add FixAn to apply on 'n')

			//---> Final Formatting
			txt = FixPunctuation(txt)
			txt = FixApostrophic(txt)
			txt = FixAn(txt) //---					---> Case: 'a/A' before Instance

			os.WriteFile(Args[1], []byte(txt), 0700) //---					---> Write byte table version
		}
	}
}
