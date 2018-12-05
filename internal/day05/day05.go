package day05

import "unicode"

func Solve(lines []string, partOne bool) string {

	s := lines[0]
	
	for {
		nextS := process(s)
		
		if s == nextS {
			break
		}
		
		s = nextS
	}
	

	return string(len(s))
}

func process(s string) string {
	nextS := ""
	for i := 0; i < (len(s) - 2); i++ {


		if ((unicode.IsLower(s[i]) && unicode.IsUpper(s[i+1])) || (unicode.IsUpper(s[i]) && unicode.isLower(s[i+1]))) {

		}


	}
}