package mutate

import (
	"github.com/zhuyanxi/gfuzz/pkg/fuzz/gexecfuzz"
	"github.com/zhuyanxi/gfuzz/pkg/oraclert/config"
	"github.com/zhuyanxi/gfuzz/pkg/oraclert/output"
)

type NoMutateStrategy struct {
}

func (d *NoMutateStrategy) Mutate(g *gexecfuzz.GExecFuzz, curr *config.Config, o *output.Output, energy int) ([]*config.Config, error) {
	var cfgs []*config.Config
	for i := 0; i < energy; i++ {
		cfg := config.NewConfig()
		cfgs = append(cfgs, cfg)
	}

	return cfgs, nil
}
