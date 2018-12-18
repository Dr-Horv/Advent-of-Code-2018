package main

import (
	"fmt"
	"github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func check(e error, msg string) {
	if e != nil {
		fmt.Println(msg)
		os.Exit(1)
	}
}

func main() {
	day := "18"
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Usage: setup.go token")
		os.Exit(1)
	}

	token := args[0]
	packagePath := createPackage(day)
	createSolverFile(packagePath, day)
	createTestFile(packagePath, day)
	fetchInput(packagePath, day, token)
}

func createPackage(day string) string {
	path := filepath.Join("./", "internal", "day"+day)
	err := os.MkdirAll(path, os.ModePerm)
	check(err, fmt.Sprintf("Failed to create folder: %v", path))
	return path
}

func fetchInput(path string, day string, token string) {
	client := &http.Client{}
	dayAsInt, _ := strconv.Atoi(day)
	url := fmt.Sprintf("https://adventofcode.com/2018/day/%v/input", dayAsInt)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Cookie", fmt.Sprintf("session=%v", token))
	resp, err := client.Do(req)
	inputPath := filepath.Join(path, "input")
	body, err := ioutil.ReadAll(resp.Body)
	check(err, "Failed to read body")
	err = pkg.WriteFileIfNew(inputPath, body)
	check(err, "Failed to write input file")
}

func createSolverFile(path string, day string) {
	solverPath := filepath.Join(path, "day"+day+".go")
	fileTemplate := "package day%v\n" +
		"\n" +
		"func Solve(lines []string, partOne bool) string {\n" +
		"	return \"\"\n" +
		"}\n"
	fileContent := fmt.Sprintf(fileTemplate, day)
	err := pkg.WriteFileIfNew(solverPath, []byte(fileContent))
	check(err, "Failed to write solver file")
}

func createTestFile(path string, day string) {
	testPath := filepath.Join(path, "day"+day+"_test.go")
	testTemplate := "package day%v\n" +
		"\n" +
		"import \"testing\"\n" +
		"\n" +
		"func TestSolve(t *testing.T) {\n" +
		"\n" +
		"	answer := Solve([]string{\"line1\"}, false)\n" +
		"\n" +
		"	if answer != \"expected\" {\n" +
		"		t.Error(\"Expected SOMETHING, got \", answer)\n" +
		"	}\n" +
		"}"
	testContent := fmt.Sprintf(testTemplate, day)
	err := pkg.WriteFileIfNew(testPath, []byte(testContent))
	check(err, "Failed to write test file")
}
