package exec

import (
	"encoding/json"
	"gfuzz/pkg/gexec"
	"gfuzz/pkg/oraclert/config"
	"gfuzz/pkg/oraclert/output"
	"path"
	"path/filepath"
)

// Stage indicates how we treat/response to an input and corresponding output
type Stage string

const (
	// InitStage simply run the empty without any mutation
	InitStage Stage = "init"

	// DeterStage is to create input by tweak select choice one by one
	DeterStage Stage = "deter"

	// CalibStage choose an input from queue to run (prepare for rand)
	CalibStage Stage = "calib"

	// RandStage randomly mutate select choice
	RandStage Stage = "rand"
)

// Input contains all information about a single execution
// (usually by fuzzer)
type Input struct {
	// ID is the unique identifer for this execution.
	ID string
	// OracleRtConfig is the configuration for the oracle runtime.
	OracleRtConfig *config.Config
	// Exec is the command to trigger a program with oracle runtime.
	Exec gexec.Executable
	// OutputDir is the output directory for this execution
	OutputDir string

	Stage Stage
}

// Output contains all useful information after a single execution
type Output struct {
	OracleRtOutput *output.Output
	BugIDs         []string
	IsTimeout      bool
}

func (i *Input) GetOrtConfigFilePath() (string, error) {
	return filepath.Abs(path.Join(i.OutputDir, "ort_config"))
}

func (i *Input) GetOutputFilePath() (string, error) {
	return filepath.Abs(path.Join(i.OutputDir, "stdout"))
}

func (i *Input) GetOrtOutputFilePath() (string, error) {
	return filepath.Abs(path.Join(i.OutputDir, "ort_output"))
}

func Serialize(l *Input) ([]byte, error) {
	if l == nil {
		return []byte{}, nil
	}

	return json.Marshal(l)
}

func Deserilize(data []byte) (*Input, error) {
	l := Input{}
	err := json.Unmarshal(data, &l)
	if err != nil {
		return nil, err
	}
	return &l, nil
}
