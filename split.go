package main

import (
	"code.google.com/p/biogo/alphabet"
	"code.google.com/p/biogo/io/seqio"
	"code.google.com/p/biogo/io/seqio/fasta"
	"code.google.com/p/biogo/seq/linear"

	"bufio"
	"fmt"
	"os"
)

func split(inputFileName string) {
	// open inputFile in read-only mode
	inputFile, inputError := os.Open(inputFileName)
	if inputError != nil {
		fmt.Println("An error occurred opening the input file.")
		return
	}
	defer inputFile.Close()

	sc := seqio.NewScanner(fasta.NewReader(inputFile, linear.NewSeq("", nil, alphabet.DNA)))
	for sc.Next() {
		s := sc.Seq().(*linear.Seq)

		// open outputFile in write-only mode
		outputFile, outputError := os.Create(s.ID + ".fa")
		if outputError != nil {
			panic(outputError)
			// fmt.Println("Could not create file " + s.ID + ".fa")
			return
		}
		defer outputFile.Close()

		outputWriter := bufio.NewWriter(outputFile)

		// write header
		outputWriter.WriteString(">" + s.ID + "\n")
		// write FASTA sequence
		outputWriter.WriteString(s.String() + "\n")

		// write the buffer completely to the file
		outputWriter.Flush()
	}
	if sc.Error() != nil {
		panic(sc.Error())
	}

}
