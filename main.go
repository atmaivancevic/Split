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
		// {"Alpaca", "vicPac2"},
		// {"AmericanAlligator", "allMis1"},
		// {"Armadillo", "dasNov3"},
		// {"AtlanticCod", "gadMor1"},
		// {"Baboon", "papAnu2"},
		// {"Budgerigar", "melUnd1"},
		// {"Bushbaby", "otoGar3"},
		// {"Cat", "felCat5"},
		// {"Chicken", "galGal4"},
		// {"Chimp", "panTro4"},
		// {"Coelacanth", "latCha1"},
		// {"Cow", "bosTau6"},
		// {"D_ananassae", "droAna3"},
		// {"D_erecta", "droEre2"},
		// {"D_grimshawi", "droGri2"},
		// {"D_mojavensis", "droMoj3"},
		// {"Dog", "canFam3"},
		// {"Dolphin", "turTru2"},
		// {"D_persimilis", "droPer1"},
		// {"D_pseudoobscura", "dp4"},
		// {"D_sechellia", "droSec1"},
		// {"D_virilis", "droVir3"},
		// {"Elephant", "loxAfr3"},
		// {"Ferret", "musFur1"},
		// {"Frog", "xenTro3"},
		// {"Fugu", "fr3"},
		// {"Gibbon", "nomLeu3"},
		// {"Gorilla", "gorGor3"},
		// {"GuineaPig", "cavPor3"},
		// {"Hedgehog", "eriEur1"},
		// {"KangarooRat", "dipOrd1"},
		// {"Lamprey", "petMar2"},
		// {"Lizard", "anoCar2"},
		// {"Manatee", "triMan1"},
		// {"Marmoset", "calJac3"},
		// {"Medaka", "oryLat2"},
		// {"MediumGroundfinch", "geoFor1"},
		// {"Megabat", "pteVam1"},
		// {"Microbat", "myoLuc2"},
		// {"MouseLemur", "micMur1"},
		// {"NakedMolerat", "hetGla2"},
		// {"NileTilapia", "oreNil2"},
		// {"PaintedTurtle", "chrPic1"},
		// {"Panda", "ailMel1"},
		// {"Pig", "susScr3"},
		// {"Pika", "ochPri2"},
		// {"Platypus", "ornAna1"},
		// {"Rabbit", "oryCun2"},
		// {"Rat", "rn5"},
		// {"Rhesus", "rheMac3"},
		// {"RockHyrax", "proCap1"},
		// {"SeaHare", "aplCal1"},
		// {"SeaUrchin", "strPur2"},
		// {"Sheep", "oviAri3"},
		// {"Stickleback", "gasAcu1"},
		// {"Shrew", "sorAra1"},
		// {"Sloth", "choHof1"},
		// {"Squirrel", "speTri2"},
		// {"SquirrelMonkey", "saiBol1"},
		// {"Tarsier", "tarSyr1"},
		// {"TasDevil", "sarHar1"},
		// {"Tenrec", "echTel2"},
		// {"TreeShrew", "tupBel1"},
		// {"Turkey", "melGal1"},
		// {"Wallaby", "macEug2"},
		// {"WhiteRhino", "cerSim1"},
		// {"ZebraFinch", "taeGut1"},
		// {"Zebrafish", "danRer7"},
		// {"CapeElephantShrew", "EleEdw1.0"},
		// {"Aardvark", "OryAfe1.0"},
		// {"CapeGoldenMole", "ChrAsi1.0"},
		// {"AmericanPika", "OchPri3.0"},
		// {"Mouse", "GRCm38.p2"},
		// {"GoldenHamster", "MesAur1.0"},
		// {"PrairieVole", "MicOch1.0"},
		// {"LesserEgyptianJerboa", "JacJac1.0"},
		// {"Chinchilla", "ChiLan1.0"},
		// {"Rat", "Rnor_5.0"},
		// {"Degu", "OctDeg1.0"},
		// {"ChineseHamster", "CriGri_1.0"},
		// {"DeerMouse", "Pman_1.0"},
		// {"GreenMonkey", "Chlorocebus_sabeus_1.0"},
		// {"CrabeatingMacaque", "Macaca_fascicularis_5.0"},
		// {"Bonobo", "panpan1"},
		// {"SumatranOrangutan", "P_pygmaeus_2.0.2"},
		// {"Human", "GRCh38"},
		// {"Tarsier", "Tarsius_syrichta_2.0.1"},
		// {"Marmoset", "Callithrix_jacchus_3.2"},
		// {"ChineseTreeShrew", "TupChi_1.0"},
		// {"WeddellSeal", "LepWed1.0"},
		// {"Walrus", "Oros_1.0"},
		// {"SiberianTiger", "PanTig1.0"},
		// {"TibetanAntelope", "PHO1.0"},
		// {"KillerWhale", "Oorc_1.1"},
		// {"Goat", "CHIR_1.0"},
		// {"Yak", "BosGru_v2.0"},
		// {"MinkeWhale", "BalAcu1.0"},
		// {"WaterBuffalo", "UMD_CASPUR_WB_2.0"},
		// {"SpermWhale", "Physeter_macrocephalus_2.0.2"},
		// {"Baiji", "Lipotes_vexillifer_v1"},
		// {"BactrianCamel", "CB1"},
		// {"MouseearedBat", "ASM32734v1"},
		// {"BlackFlyingFox", "ASM32557v1"},
		// {"BigBrownBat", "EptFus1.0"},
		// {"StrawcolouredFruitBat", "ASM46528v1"},
		// {"GreaterFalseVampireBat", "ASM46534v1"},
		// {"ParnellsMustachedBat", "ASM46540v1"},
		// {"Hedgehog", "EriEur2.0"},
		// {"CommonShrew", "SorAra2.0"},
		// {"StarnosedMole", "ConCri1.0"},
		// {"BurmesePython", "Python_molurus_bivittatus_5.0.2"},
		// {"SpinySoftshellTurtle", "ASM38561v1"},
		// {"GreenSeaTurtle", "CheMyd_1.0"},
		// {"ChineseSoftshellTurtle", "PelSin_1.0"},
		// {"ChineseAlligator", "ASM45574v1"},
		// {"ScarletMacaw", "SMACv1.1"},
		// {"WhitethroatedSparrow", "Zonotrichia_albicollis_1.0.1"},
		// {"Mallard", "BGI_duck_1.0"},
		// {"RockDove", "Cliv_1.0"},
		// {"SakerFalcon", "F_cherrug_v1.0"},
		// {"PeregrineFalcon", "F_peregrinus_v1.0"},
		// {"PuertoRicanParrot", "AV1"},
		// {"GroundTit", "PseHum1.0"},
		// {"AtlanticCanary", "SCA1"},
		// {"JapaneseQuail", "Coja_1.0"},
		// {"CollaredFlycatcher", "FicAlb1.5"},
		// {"WesternClawedFrog", "Xtropicalis_v7"},
		// {"AustralianGhostshark", "Callorhinchus_milii_6.1.3"},
		// {"YellowbellyPufferfish", "version_1_of_Takifugu_flavidus_genome"},
		// {"MexicanTetra", "Astyanax_mexicanus_1.0.2"},
		// {"Medaka", "ASM31367v1"},
		// {"ZebraMbuna", "MetZeb1.1"},
		// {"SpottedGar", "LepOcu1"},
		// {"SouthernPlatyfish", "Xiphophorus_maculatus_4.4.2"},
		// {"AstatotilapiaBurtoni", "AstBur1.0"},
		// {"NeolamprologusBrichardi", "NeoBri1.0"},
		// {"FlameBackCichlid", "PunNye1.0"},
		// {"ArticLamprey", "LetJap1.0"},
		// {"GreaterHorseshoeBat", "ASM46549v1"},
		// {"BrandtsBat", "ASM41265v1"},
		// {"VaseTunicate", "KH"},
		// {"StarAscidian", "356a_chromosome_assembly"},
		// {"BatStar", "Pmin_1.0"},
		// {"GreenSeaUrchin", "Lvar_0.4"},
		// {"PurpleSeaUrchin", "Spur_3.1"},
		// {"AcornWorm", "Skow_1.1"},
		// {"Schisto", "ASM23792v2"},
		// {"ChineseLiverFluke", "C_sinensis_2.0"},
		// {"RodentTapeworm", "HMIC001"},
		// {"HyperTapeworm", "EGRAN001"},
		// {"FoxTapeworm", "EMULTI001"},
		// {"C_elegans", "WBcel235"},
		// {"Microworm", "Pred3"},
		// {"PigRoundworm", "AscSuum_1.0"},
		// {"BeneficialNematode", "Heterorhabditis_bacteriophora_7.0"},
		// {"PorkWorm", "Trichinella_spiralis_3.7.1"},
		// {"C_11", "Caenorhabditis_sp11_JU1373_3.0.1"},
		// {"C_brenneri", "C_brenneri_6.0.1b"},
		// {"C_angaria", "ps1010rel4"},
		// {"C_japonica", "C_japonica_7.0.1"},
		// {"HumanHookworm", "N_americanus_v1"},
		// {"S_monti", "S_monti_v1"},
		// {"RedDeerNematode", "EEL001"},
		// {"O_vol", "OVOC001"},
		// {"WireWorm", "HCON"},
		// {"AsianBeetle", "Agla_1.0"},
		// {"D_pseudo", "Dpse_3.0"},
		// {"Housefly", "Musca_domestica_2.0.2"},
		// {"Dragonfly", "Lful_1.0"},
		// {"D_miranda", "DroMir_2.2"},
		// {"MountainPineBeetle", "DendPond_male_1.0"},
		// {"A_quad", "Anop_quad_QUAD4_A_V1"},
		// {"A_funestus", "Anop_fune_FUMOZ_V1"},
		// {"A_epiro", "Anop_epir_epiroticus2_V1"},
		// {"A_alb", "Anop_albi_ALBI9_A_V1"},
		// {"A_dirus", "Anop_diru_WRAIR2_V1"},
		// {"A_christyi", "Anop_chri_ACHKN1017_V1"},
		// {"A_arab", "Anop_arab_DONG5_A_V1"},
		// {"A_minimus", "Anop_mini_MINIMUS1_V1"},
		// {"Medfly", "Ccap_1.0"},
		// {"TurnipSawfly", "Aros_1.0"},
		// {"D_eugra", "Deug_2.0"},
		// {"D_rho", "Drho_2.0"},
		// {"D_bipect", "Dbip_2.0"},
		// {"D_taka", "Dtak_2.0"},
		// {"D_kikka", "Dkik_2.0"},
		// {"D_elegans", "Dele_2.0"},
		// {"D_bia", "Dbia_2.0"},
		// {"D_ficus", "Dfic_2.0"},
		// {"DiamondbackMoth", "DBM_FJ_V1.1"},
		// {"PostmanButterfly", "ASM31383v2"},
		// {"D_albo", "DroAlb_1.0"},
		// {"Sandfly", "Llon_1.0"},
		// {"WesternPredatoryMite", "Mocc_1.0"},
		// {"TobaccoHornworm", "Msex_1.0"},
		// {"PhleboSandfly", "Ppap_1.0"},
		// {"AlfalfaLeafcutterBee", "MROT_1.0"},
		// {"DwarfHoneyBee", "Aflo_1.0"},
		// {"S_maritima", "Smar_1.0"},
		// {"RedSpiderMite", "ASM23943v1"},
		// {"MonarchButterfly", "DanPle_1.0"},
		// {"CommonEasternBumbleBee", "BIMP_2.0"},
		// {"BufftailedBumbleBee", "Bter_1.0"},
		// {"ArgentineAnt", "Lhum_UMD_V04"},
		// {"PanamanianLeafcutterAnt", "Aech_3.9"},
		// {"PeaAphid", "Acyr_2.0"},
		// {"WaterFlea", "V1.0"},
		// {"RedFireAnt", "Si_gnG"},
		// {"RedHarvesterAnt", "Pbar_UMD_V03"},
		// {"AssassinBug", "Rhodnius_prolixus_3.0.1"},
		// {"LeafcutterAnt", "Attacep1.0"},
		// {"HessianFly", "Mdes_1.0"},
		// {"Silkworm", "ASM15162v1"},
		// {"JumpingAnt", "HarSal_1.0"},
		// {"BlackCarpenterAnt", "CamFlo_1.0"},
		// {"A_melas", "Anop_mela_CM1001059_A_V2"},
		// {"A_farauti", "Anop_fara_FAR1_V2"},
		// {"A_merus", "Anop_meru_MAF_V1"},
		// {"AtlanticHorseshoeCrab", "Limulus_polyphemus_2.1.2"},
		// {"GreenDrake", "Edan_1.0"},
		// {"PollinatingWasp", "CerSol_1.0"},
		// {"ColoradoPotatoBeetle", "Ldec_1.5"},
		// {"AsianCitrusPsyllid", "Diaci_psyllid_genome_assembly_version_1.1"},
		// {"A_atro", "Anop_atro_EBRO_V1"},
		// {"A_cul", "Anop_culi_species_A_37_1_V1"},
		// {"A_maculatus", "Anop_macu_maculatus3_V1"},
		// {"D_suzukii", "Dsuzukii.v01"},
		// {"CommonHouseSpider", "Ptep_1.0"},
		// {"GiantHoneyBee", "Apis_dorsata_1.3"},
		// {"A_sinensis", "AS2"},
		// {"A_steph", "ASM30077v2"},
		// {"D_willistoni", "dwil_caf1"},
		// {"N_gir", "Ngir_1.0"},
		// {"N_long", "Nlon_1.0"},
		// {"CatusWorm", "Priapulus_caudatus_4.0.1"},
		// {"CaliforniaSeaHare", "AplCal3.0"},
		// {"PolychaeteWorm", "Capca1"},
		// {"Leech", "Helobdella_robusta_v1.0"},
		// {"OwlLimpet", "Helro1"},
		// {"PacificOyster", "oyster_v9"},
		// {"Rotifer", "ASM_PRJEB1171_v1"},
		// {"FreshwaterSnail", "ASM45736v1"},
		// {"FreshwaterPolyp", "Hydra_RP_1.0"},
		// {"WartyCombJelly", "MneLei_Aug2011"},
		// {"SeaSponge", "v1.0"},
		// {"C_semi", "Cse_v1.0"},
		// {"AmazonMolly", "Poecilia_formosa_5.1.2"},
		// {"FlagRockfish", "SRub1.0"},
		// {"TigerRockfish", "Snig1.0"},
		// {"ElephantShark", "calMil1"},
		// {"Guppy", "Guppy_female_1.0_MT"},
		// {"BlindMolerat", "S.galili_v1.0"},
		// {"BlackCormorant", "ASM70892v1"},
		// {"Elephant", "LAv4"},
		{"PaintedTurtle", "Chrysemys_picta_bellii_3.0.3"},
	} {
		splitGenome(gen.dir, "fa", gen.file)
	}
}
