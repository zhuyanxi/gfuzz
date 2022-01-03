package main

import (
	"os"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {

	// Fuzzing Target
	GoModDir     string   `long:"gomod" description:"Directory contains go.mod"`
	TestFunc     []string `long:"func" description:"Only run specific test function in the test"`
	TestPkg      []string `long:"pkg" description:"Only run test functions in the specific package"`
	TestBinGlobs []string `long:"bin" description:"A list of globs for Go test bins."`
	Ortconfig    string   `long:"ortconfig" description:"Only run once with given ortconfig (replay)"`
	Repeat       int      `long:"repeat" description:"how many replay" default:"1"`

	// Fuzzer
	OutputDir string `long:"out" description:"Directory for fuzzing output"`
	Parallel  int    `long:"parallel" description:"Number of workers to fuzz parallel" default:"5"`
	InstStats string `long:"instStats" description:"This parameter consumes a file path to a statistics file generated by isnt."`
	Version   bool   `long:"version" description:"Print version and exit"`

	// Fuzzing
	GlobalTuple         bool `long:"globalTuple" description:"Whether prev_location is global or per channel"`
	ScoreSdk            bool `long:"scoreSdk" description:"Recording/scoring if channel comes from Go SDK"`
	ScoreAllPrim        bool `long:"scoreAllPrim" description:"Recording/scoring other primitives like Mutex together with channel"`
	TimeDivideBy        int  `long:"timedivideby" description:"Durations in time/sleep.go will be divided by this int number"`
	OracleRtDebug       bool `long:"ortdebug"`
	SelEfcmTimeout      int  `long:"setimeout" default:"500" description:"default select enforcement timeout"`
	FixedSelEfcmTimeout bool `long:"fixedsetimeout" description:"disable automatically select enforcement timeout mutating"`
	ScoreBasedEnergy    bool `long:"scoreenergy"`
	AllowDupCfg         bool `long:"allowdupcfg" description:"allow duplicated randomly generated configuration to be run"`
	IsIgnoreFeedback    bool `long:"ignorefeedback" description:"Is ignoring the feedback, and save every mutated seed into the fuzzing queue"`
	RandMutateEnergy    int  `long:"randMutateEnergy" description:"Determine the energy of random mutations. If == 100 (default), then each seed would mutate 100 times in the rand mutation stage"`
	IsDisableScore      bool `long:"disablescore" description:"Is disable score to priority testing case. "`
	NoSelEfcm           bool `long:"nose" description:"Disable select enforcement"`

	NfbRandEnergy         bool `long:"nfbrandenergy" description:"should energy be randomly generated in non-feedback"`
	NfbRandSelEfcmTimeout bool `long:"nfbrandsetimeout" description:"should timeout of select enforcement be randomly generated in non-feedback"`
	MemRandStrat          bool `long:"memrandstrat" description:"prioritize generating non-triggered case in given energy"`
}

func parseFlags() {

	if _, err := flags.Parse(&opts); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}

}
