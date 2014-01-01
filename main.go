package main

import (
	"fmt"
	"os"
)

const genomeDir = "/scratch/atmaGenomes/"

// uses the split function on each genome

func splitGenome(genomeName, extension, identifier string) {

	os.Chdir(genomeDir + genomeName)

	inputFileName := identifier + "_unmasked." + extension

	// check if the genome unmasked file exists - if so, split it
	if _, err := os.Stat(inputFileName); err == nil {
		status := split(inputFileName)
		// print a status message
		switch status {
		case 0:
			fmt.Println("Processed: " + genomeName + "/" + identifier + "_unmasked")
		case 1:
			fmt.Println("Processed: " + genomeName + "/" + identifier + "_unmasked; some files were skipped")
		case 2:
			fmt.Println("Processed: " + genomeName + "/" + identifier + "_unmasked; no new files created")
		default:
			fmt.Println("Something went wrong with " + genomeName + "/" + identifier + "_unmasked")
		}
	} else {
		fmt.Println("File " + genomeName + "/" + identifier + "_unmasked does not exist")
	}
}

func main() {

	for _, gen := range []struct {
		dir  string
		file string
	}{
		{"Alpaca", "vicPac2"},
		{"AmericanAlligator", "allMis1"},
		{"Armadillo", "dasNov3"},
		{"AtlanticCod", "gadMor1"},
		{"Baboon", "papAnu2"},
		{"Budgerigar", "melUnd1"},
		{"Bushbaby", "otoGar3"},
		{"Cat", "felCat5"},
		{"Chicken", "galGal4"},
		{"Chimp", "panTro4"},
		{"Coelacanth", "latCha1"},
		{"Cow", "bosTau6"},
		{"D_ananassae", "droAna3"},
		{"D_erecta", "droEre2"},
		{"D_grimshawi", "droGri2"},
		{"D_mojavensis", "droMoj3"},
		{"Dog", "canFam3"},
		{"Dolphin", "turTru2"},
		{"D_persimilis", "droPer1"},
		{"D_pseudoobscura", "dp4"},
		{"D_sechellia", "droSec1"},
		{"D_virilis", "droVir3"},
		{"Elephant", "loxAfr3"},
		{"Ferret", "musFur1"},
		{"Frog", "xenTro3"},
		{"Fugu", "fr3"},
		{"Gibbon", "nomLeu3"},
		{"Gorilla", "gorGor3"},
		{"GuineaPig", "cavPor3"},
		{"Hedgehog", "eriEur1"},
		{"KangarooRat", "dipOrd1"},
		{"Lamprey", "petMar2"},
		{"Lizard", "anoCar2"},
		{"Manatee", "triMan1"},
		{"Marmoset", "calJac3"},
		{"Medaka", "oryLat2"},
		{"MediumGroundfinch", "geoFor1"},
		{"Megabat", "pteVam1"},
		{"Microbat", "myoLuc2"},
		{"MouseLemur", "micMur1"},
		{"NakedMolerat", "hetGla2"},
		{"NileTilapia", "oreNil2"},
		{"PaintedTurtle", "chrPic1"},
		{"Panda", "ailMel1"},
		{"Pig", "susScr3"},
		{"Pika", "ochPri2"},
		{"Platypus", "ornAna1"},
		{"Rabbit", "oryCun2"},
		{"Rat", "rn5"},
		{"Rhesus", "rheMac3"},
		{"RockHyrax", "proCap1"},
		{"SeaHare", "aplCal1"},
		{"SeaUrchin", "strPur2"},
		{"Sheep", "oviAri3"},
		{"Stickleback", "gasAcu1"},
		{"Shrew", "sorAra1"},
		{"Sloth", "choHof1"},
		{"Squirrel", "speTri2"},
		{"SquirrelMonkey", "saiBol1"},
		{"Tarsier", "tarSyr1"},
		{"TasDevil", "sarHar1"},
		{"Tenrec", "echTel2"},
		{"TreeShrew", "tupBel1"},
		{"Turkey", "melGal1"},
		{"Wallaby", "macEug2"},
		{"WhiteRhino", "cerSim1"},
		{"ZebraFinch", "taeGut1"},
		{"Zebrafish", "danRer7"},
	} {
		splitGenome(gen.dir, "fa", gen.file)
	}
}
