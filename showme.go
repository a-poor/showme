package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

const InputFile = "sample.md"

func clean(s string) string {
	stripped := strings.TrimSpace(s)
	re := regexp.MustCompile(`\n{2,}`)
	return re.ReplaceAllLiteralString(stripped, "\n\n")
}

func main() {
	fmt.Println(strings.Repeat("=", 60))

	// Read the file.
	file, err := ioutil.ReadFile(InputFile)
	if err != nil {
		panic(err)
	}

	src := string(file)
	fmt.Println(src)
	fmt.Println(strings.Repeat("=", 60))

	cleaned := clean(src)

	fmt.Println(cleaned)
	fmt.Println(strings.Repeat("=", 60))

	// fmt.Println(strings.Repeat("=", 60))
}
