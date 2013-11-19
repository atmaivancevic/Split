package main

import (
	"fmt"
	"os"
)

const genomeDir = "scratch/atmaGenomes/"

func splitGenome(genomeName, extension, identifier string) {

	os.Chdir(genomeDir + genomeName)

	inputFileName := identifier + "_unmasked." + extension

	split(inputFileName)
	fmt.Println("Processed: " + genomeName + "/" + identifier + "_unmasked")

}

func main() {
	splitGenome("Alpaca", "fa", "vicPac2")
}
