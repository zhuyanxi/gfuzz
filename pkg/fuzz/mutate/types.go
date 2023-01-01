package mutate

import (
	"github.com/zhuyanxi/gfuzz/pkg/fuzz/gexecfuzz"
	"github.com/zhuyanxi/gfuzz/pkg/oraclert/config"
	"github.com/zhuyanxi/gfuzz/pkg/oraclert/output"
)

type OrtConfigMutateStrategy interface {
	Mutate(g *gexecfuzz.GExecFuzz, curr *config.Config, o *output.Output, energy int) ([]*config.Config, error)
}
