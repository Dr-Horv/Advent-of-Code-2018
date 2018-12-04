package pkg

import (
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(path string) []string {
	dat, err := ioutil.ReadFile(path)
	check(err)
	content := string(dat)
	lines := strings.Split(content, "\n")
	return lines
}
