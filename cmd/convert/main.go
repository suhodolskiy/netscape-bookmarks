// Tool to convert bookmark netscape file to json
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	bookmarks "github.com/suhodolskiy/netscape-bookmarks"
)

var (
	inputPath  string
	outputPath string
)

func init() {
	flag.StringVar(&inputPath, "input", "", "input file")
	flag.StringVar(&outputPath, "output", "./example.json", "output file")
	flag.Parse()
}

func main() {
	if inputPath == "" {
		panic(fmt.Errorf("input file path is required"))
	}

	input, err := os.OpenFile(inputPath, os.O_RDONLY, 0644)

	if err != nil {
		panic(fmt.Errorf("failed to open input file (%s)", err))
	}

	data, err := bookmarks.Parse(input)

	if err != nil {
		panic(fmt.Errorf("failed to parse input file (%s)", err))
	}

	output, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		panic(fmt.Errorf("failed to marshal data (%s)", err))
	}

	if err := ioutil.WriteFile(outputPath, output, 0644); err != nil {
		panic(fmt.Errorf("failed to write output file (%s)", err))
	}
}
