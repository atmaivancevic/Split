package main

import (
	"fmt"
	"os"
)

const genomeDir = "scratch/atmaGenomes/"

func splitGenome(genomeName, extension, identifier string) {

	os.Chdir(genomeDir + genomeName)

	inputFileName := identifier + "." + extension

	split(inputFileName)
	fmt.Println("Processed: " + genomeName + "/" + identifier)

}

func main() {
	splitGenome("Alpaca", "fa", "vicPac2")
}
