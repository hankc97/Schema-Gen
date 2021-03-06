package main

import (
	"flag"
	"log"
	"os"
	"gopkg.in/yaml.v3"
	"github.com/hankc97/fixtures/gen_schema"
	"github.com/hankc97/parse/lint"
)

func main() {
	inputFilename := flag.String("i", "", "name of the file to lint")
	flag.Parse()

	if *inputFilename == "" {
		log.Fatalf("-i is required")
	}
	if err := realMain(*inputFilename); err != nil {
		log.Fatal(err)
	}
}

func realMain(inputFilename string) error {
	input, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	defer input.Close()

	sink := &lint.ProblemSink{Filename: inputFilename, Output: os.Stdout}
	node := new(gen_schema.WorkflowRoot)
	
	if err := yaml.NewDecoder(input).Decode(&node); err != nil {
		return err
	}

	if err := lint.LintWorkflow(sink, node); err != nil {
		return err
	}

	return nil
}

// 2 levels to parsing
// checks for basic yaml file strucutre
// then we create our own lint functions

