package main

import (
	"code.google.com/p/biogo/alphabet"
	"code.google.com/p/biogo/io/seqio"
	"code.google.com/p/biogo/io/seqio/fasta"
	"code.google.com/p/biogo/seq/linear"

	"bufio"
	"os"
	"strconv"
)

// this function splits a genome into smaller files,
// where each new file consists of one header with seq ID and the corresponding FASTA sequence
// returns: 0 if all files are created and none are skipped
// 			1 if some files are created and some are skipped
//			2 if no files are created and all are skipped
//			3 otherwise (no files were created or skipped)
func split(inputFileName string) int {

	fileSkipped := false
	fileCreated := false

	// open input file in read-only mode
	inputFile, inputError := os.Open(inputFileName)
	if inputError != nil {
		panic(inputError)
	}
	defer inputFile.Close()

	sc := seqio.NewScanner(fasta.NewReader(inputFile, linear.NewSeq("", nil, alphabet.DNA)))
	seqNo := 1;
	for sc.Next() {
		s := sc.Seq().(*linear.Seq)

		// check if the output file already exists
		if _, err := os.Stat("seq" + strconv.Itoa(seqNo) + ".fa"); os.IsNotExist(err) {
			// open output file in write-only mode
			outputFile, outputError := os.Create("seq" + strconv.Itoa(seqNo) + ".fa")
			if outputError != nil {
				panic(outputError)
			}
			outputWriter := bufio.NewWriter(outputFile)
			// write header
			outputWriter.WriteString(">" + s.ID + "\n")
			// write FASTA sequence
			outputWriter.WriteString(s.String() + "\n")
			// write the buffer completely to the file
			outputWriter.Flush()
			// close the output file
			outputFile.Close()

			seqNo++

			fileCreated = true
		} else {
			fileSkipped = true
		}
	}
	if sc.Error() != nil {
		panic(sc.Error())
	}

	if fileCreated && (!fileSkipped) {
		return 0
	} else if fileCreated && fileSkipped {
		return 1
	} else if (!fileCreated) && fileSkipped {
		return 2
	} else {
		return 3
	}
}
