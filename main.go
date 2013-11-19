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

	// maybe do a check if the genome has already been split
	//(i.e. are there any gi*.fa files in the directory already? If so, don't split)

}

func main() {
	// splitGenome("Alpaca", "fa", "vicPac2")

	// splitGenome("AmericanAlligator", "fa", "allMis1")

	// splitGenome("Armadillo", "fa", "dasNov3")

	// splitGenome("AtlanticCod", "fa", "gadMor1")

	// splitGenome("Budgerigar", "fa", "melUnd1")

	// splitGenome("Bushbaby", "fa", "otoGar3")

	// splitGenome("Coelacanth", "fa", "latCha1")

	// splitGenome("D_ananassae", "fa", "droAna3")

	// splitGenome("D_erecta", "fa", "droEre2")

	// splitGenome("D_grimshawi", "fa", "droGri2")

	// splitGenome("D_mojavensis", "fa", "droMoj3")

	// splitGenome("Dolphin", "fa", "turTru2")

	// splitGenome("D_persimilis", "fa", "droPer1")

	// splitGenome("D_pseudoobscura", "fa", "dp4")

	// splitGenome("D_sechellia", "fa", "droSec1")

	// splitGenome("D_virilis", "fa", "droVir3")

	// splitGenome("Ferret", "fa", "musFur1")

	// splitGenome("Frog", "fa", "xenTro3")

	// splitGenome("Fugu", "fa", "fr3")

	// splitGenome("GuineaPig", "fa", "cavPor3")

	// splitGenome("Hedgehog", "fa", "eriEur1")

	// splitGenome("KangarooRat", "fa", "dipOrd1")

	// splitGenome("Lamprey", "fa", "petMar2")

	// splitGenome("Manatee", "fa", "triMan1")

	// splitGenome("MediumGroundfinch", "fa", "geoFor1")

	// splitGenome("Megabat", "fa", "pteVam1")

	// splitGenome("Microbat", "fa", "myoLuc2")

	// splitGenome("MouseLemur", "fa", "micMur1")

	// splitGenome("NakedMolerat", "fa", "hetGla2")

	// splitGenome("PaintedTurtle", "fa", "chrPic1")

	// splitGenome("Panda", "fa", "ailMel1")

	// splitGenome("Pika", "fa", "ochPri2")

	// splitGenome("RockHyrax", "fa", "proCap1")

	// splitGenome("SeaHare", "fa", "aplCal1")

	// splitGenome("SeaUrchin", "fa", "strPur2")

	// splitGenome("Shrew", "fa", "sorAra1")

	// splitGenome("Sloth", "fa", "choHof1")

	// splitGenome("Squirrel", "fa", "speTri2")

	// splitGenome("SquirrelMonkey", "fa", "saiBol1")

	// splitGenome("Tarsier", "fa", "tarSyr1")

	// splitGenome("Tenrec", "fa", "echTel2")

	// splitGenome("TreeShrew", "fa", "tupBel1")

	// splitGenome("Turkey", "fa", "melGal1")

	// splitGenome("Wallaby", "fa", "macEug2")

	// splitGenome("WhiteRhino", "fa", "cerSim1")

	for _, gen := range []struct {
		dir  string
		file string
	}{
		{"Alpaca", "vicPac2"},
		{"AmericanAlligator", "allMis1"},
		{"Armadillo", "dasNov3"},
		{"AtlanticCod", "gadMor1"},
		{"Budgerigar", "melUnd1"},
		{"Bushbaby", "otoGar3"},
		{"Coelacanth", "latCha1"},
		{"D_ananassae", "droAna3"},
		{"D_erecta", "droEre2"},
		{"D_grimshawi", "droGri2"},
		{"D_mojavensis", "droMoj3"},
		{"Dolphin", "turTru2"},
		{"D_persimilis", "droPer1"},
		{"D_pseudoobscura", "dp4"},
		{"D_sechellia", "droSec1"},
		{"D_virilis", "droVir3"},
		{"Ferret", "musFur1"},
		{"Frog", "xenTro3"},
		{"Fugu", "fr3"},
		{"GuineaPig", "cavPor3"},
		{"Hedgehog", "eriEur1"},
		{"KangarooRat", "dipOrd1"},
		{"Lamprey", "petMar2"},
		{"Manatee", "triMan1"},
		{"MediumGroundfinch", "geoFor1"},
		{"Megabat", "pteVam1"},
		{"Microbat", "myoLuc2"},
		{"MouseLemur", "micMur1"},
		{"NakedMolerat", "hetGla2"},
		{"PaintedTurtle", "chrPic1"},
		{"Panda", "ailMel1"},
		{"Pika", "ochPri2"},
		{"RockHyrax", "proCap1"},
		{"SeaHare", "aplCal1"},
		{"SeaUrchin", "strPur2"},
		{"Shrew", "sorAra1"},
		{"Sloth", "choHof1"},
		{"Squirrel", "speTri2"},
		{"SquirrelMonkey", "saiBol1"},
		{"Tarsier", "tarSyr1"},
		{"Tenrec", "echTel2"},
		{"TreeShrew", "tupBel1"},
		{"Turkey", "melGal1"},
		{"Wallaby", "macEug2"},
		{"WhiteRhino", "cerSim1"},
	} {
		splitGenome(gen.dir, "fa", gen.file)
	}
}
