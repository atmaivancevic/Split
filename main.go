package main

import (
	"fmt"
	"os"
)

const genomeDir = "/scratch/atmaGenomes/"

func splitGenome(genomeName, extension, identifier string) {

	os.Chdir(genomeDir + genomeName)

	inputFileName := identifier + "_unmasked." + extension

	// check if the genome unmasked file exists - if so, split it
	if _, err := os.Stat(inputFileName); err == nil {
		split(inputFileName)
		fmt.Println("Processed: " + genomeName + "/" + identifier + "_unmasked")
	} else {
		fmt.Println("File " + genomeName + "/" + identifier + "_unmasked does not exist")
	}

}

func main() {
	splitGenome("Alpaca", "fa", "vicPac2")
}
