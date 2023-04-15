package main

import (
	"strconv"
	"strings"
)

func ApplyInstances(s string) string {
	t := strings.Fields(s)
	instances := []string{"hex", "bin", "cap", "up", "low"} //---					---> Gather instances
	var apply func(string) string                           //---					---> Common function

	//---> Change Function based on Instance
	for i := range t {
		for _, instance := range instances {
			switch instance {
			case "hex":
				apply = HexToDec
			case "bin":
				apply = BinToDec
			case "cap":
				apply = Capital
			case "up":
				apply = strings.ToUpper
			case "low":
				apply = strings.ToLower
			}

			//---> Single Target
			if t[i] == "("+instance+")" {
				for o := 1; o <= i; o++ { //---					---> Manage Multiple Instances
					if t[i-o] != "" {
						t[i-o] = apply(t[i-o])
						break //---					---> Avoid applying on all previous Words
					}
				}
				t[i] = ""
			}

			//---> Multiple Targets
			if i < len(t)-1 && t[i] == "("+instance+"," && t[i+1][len(t[i+1])-1] == 41 { //---> Last index must ')'
				b, err := strconv.Atoi(t[i+1][:len(t[i+1])-1]) //---					---> Convert the entire String
				if err == nil {
					for d := 1; d <= i && d <= b; d++ { //---					---> Manage Targets
						if t[i-d] == "" {
							b++ //---					---> Manage Multiple Instances
						} else {
							t[i-d] = apply(t[i-d])
						}
					}
				}
				t[i+1] = ""
				t[i] = ""
			}
		}
	}

	return FixSpaces(strings.Join(t, " ")) //---					---> Fix Double Spaces before Returning
}
