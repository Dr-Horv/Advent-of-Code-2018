package pkg

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteFileIfNew(path string, content []byte) error {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return errors.New("file already exists")
	}

	return ioutil.WriteFile(path, content, os.ModePerm)
}

func ReadFile(path string) []string {
	dat, err := ioutil.ReadFile(path)
	check(err)
	content := string(dat)
	lines := strings.Split(content, "\n")
	return lines
}
