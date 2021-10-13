package main

import (
	"flag"
	"log"
	"os"
	"github.com/a-h/generate"
	"github.com/hankc97/cmd/codegen/construct"
	"fmt"
)

var (
	i = flag.String("i", "", "input JSON file for schema generation")
	o = flag.String("o", "", "output file for JSON schema")
)

func main() {
	// create flags with errs
	flag.Parse()
	if *i == "" {
		flag.PrintDefaults()
		log.Fatal("required flag: -i, with JSON file path")
	}
	if *o == "" {
		flag.PrintDefaults()
		log.Fatal("required flag: -o, with output file path")
	}

	// create writer
	writer, err := os.Create(*o)
	if err != nil {
		log.Fatal("Error opening output file: ", err)
	}
	defer writer.Close()

	// place input file path into a slice
	inputFiles := []string{*i}
	// output map schema
	schemas, err := generate.ReadInputFiles(inputFiles, false)

	if err != nil {
		log.Fatal(err)
	}

	// generate Map of structs needed for output.go
	// generatedSchemaMap := generate.New(schemas...)
	// err = generatedSchemaMap.CreateTypes()

	constructor := construct.Constructor{Schemas : schemas}
	err = constructor.CreateStruct()

	if err != nil {
		log.Fatal("Failure generating structs:", err)
	}

	// gen map structs
	// fmt.Printf("%#v", generatedSchemaMap.Structs)

	// construct.Output(writer, generatedSchemaMap)
}