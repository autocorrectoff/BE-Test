package main

import (
	"flag"
	"github.com/autocorrectoff/BE-Test/engine"
	"github.com/autocorrectoff/BE-Test/utils"
	"log"
	"os"
)

func main() {
	inputFile := flag.String("input-file", "", "File to load data from")
	outputFile := flag.String("output-file", "output.html", "File to dump processed data to")
	flag.Parse()

	log.Println("Processing data")
	str, err := engine.Process(*inputFile)
	utils.HandleError(err)

	file, err := os.Create(*outputFile)
	utils.HandleError(err)
	defer file.Close()
	file.WriteString(str)

	log.Println("Completed")
}
