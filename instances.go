package main

import (
	"strconv"
	"strings"
)

func ApplyInstances(s string) string {
	t := strings.Fields(s)
	instances := []string{"hex", "bin", "cap", "up", "low"}
	var apply func(string) string

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
			if t[i] == "("+instance+")" {
				for o := 1; o <= i; o++ {
					if t[i-o] == "" {
						continue
					} else {
						t[i-o] = apply(t[i-o])
						break
					}
				}
				t[i] = ""
			}
			if i < len(t)-1 && t[i] == "("+instance+"," && t[i+1][len(t[i+1])-1] == 41 {
				b, err := strconv.Atoi(t[i+1][:len(t[i+1])-1])
				if err == nil {
					for d := 1; d <= b && d <= i; d++ {
						if t[i-d] == "" {
							b++
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
	return FixSpaces(strings.Join(t, " "))
}
