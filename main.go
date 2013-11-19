package main

import (
	"fmt"
	"os"
)

const genomeDir = "/scratch/atmaGenomes/"

func splitGenome(genomeName, prefix, extension, identifier string) {

	os.Chdir(genomeDir + genomeName)

	inputFileName := identifier + "_unmasked." + extension

	// check if the genome unmasked file exists - if so, split it
	if _, err := os.Stat(inputFileName); err == nil {
		split(inputFileName)
		fmt.Println("Processed: " + genomeName + "/" + identifier + "_unmasked")
	} else {
		fmt.Println("File " + genomeName + "/" + identifier + "_unmasked does not exist")
	}

	// maybe do a check if the genome has already been split
	//(i.e. are there any gi*.fa files in the directory already? If so, don't split)

}

func main() {
	splitGenome("Alpaca", "fa", "vicPac2")

	splitGenome("AmericanAlligator", "fa", "allMis1")

	splitGenome("Armadillo", "fa", "dasNov3")

	splitGenome("AtlanticCod", "fa", "gadMor1")

	splitGenome("Budgerigar", "fa", "melUnd1")

	splitGenome("Bushbaby", "fa", "otoGar3")

	splitGenome("Coelacanth", "fa", "latCha1")

	splitGenome("D_ananassae", "fa", "droAna3")

	splitGenome("D_erecta", "fa", "droEre2")

	splitGenome("D_grimshawi", "fa", "droGri2")

	splitGenome("D_mojavensis", "fa", "droMoj3")

	splitGenome("Dolphin", "fa", "turTru2")

	splitGenome("D_persimilis", "fa", "droPer1")

	splitGenome("D_pseudoobscura", "fa", "dp4")

	splitGenome("D_sechellia", "fa", "droSec1")

	splitGenome("D_virilis", "fa", "droVir3")

	splitGenome("Ferret", "fa", "musFur1")

	splitGenome("Frog", "fa", "xenTro3")

	splitGenome("Fugu", "fa", "fr3")

	splitGenome("GuineaPig", "fa", "cavPor3")

	splitGenome("Hedgehog", "fa", "eriEur1")

	splitGenome("KangarooRat", "fa", "dipOrd1")

	splitGenome("Lamprey", "fa", "petMar2")

	splitGenome("Manatee", "fa", "triMan1")

	splitGenome("MediumGroundfinch", "fa", "geoFor1")

	splitGenome("Megabat", "fa", "pteVam1")

	splitGenome("Microbat", "fa", "myoLuc2")

	splitGenome("MouseLemur", "fa", "micMur1")

	splitGenome("NakedMolerat", "fa", "hetGla2")

	splitGenome("PaintedTurtle", "fa", "chrPic1")

	splitGenome("Panda", "fa", "ailMel1")

	splitGenome("Pika", "fa", "ochPri2")

	splitGenome("RockHyrax", "fa", "proCap1")

	splitGenome("SeaHare", "fa", "aplCal1")

	splitGenome("SeaUrchin", "fa", "strPur2")

	splitGenome("Shrew", "fa", "sorAra1")

	splitGenome("Sloth", "fa", "choHof1")

	splitGenome("Squirrel", "fa", "speTri2")

	splitGenome("SquirrelMonkey", "fa", "saiBol1")

	splitGenome("Tarsier", "fa", "tarSyr1")

	splitGenome("Tenrec", "fa", "echTel2")

	splitGenome("TreeShrew", "fa", "tupBel1")

	splitGenome("Turkey", "fa", "melGal1")

	splitGenome("Wallaby", "fa", "macEug2")

	splitGenome("WhiteRhino", "fa", "cerSim1")

}
