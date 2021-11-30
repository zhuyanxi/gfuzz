package mutate

import (
	"crypto/rand"
	"fmt"
	"gfuzz/pkg/fuzz/gexecfuzz"
	"gfuzz/pkg/oraclert/config"
	"gfuzz/pkg/oraclert/output"
	"gfuzz/pkg/selefcm"
	"math/big"
)

type RandomMutateStrategy struct{}

func (d *RandomMutateStrategy) Mutate(g *gexecfuzz.GExecFuzz, curr *config.Config, o *output.Output, energy int) ([]*config.Config, error) {
	var cfgs []*config.Config

	for mutate_idx := 0; mutate_idx < energy; mutate_idx++ {
		cfg := config.NewConfig()
		cfg.SelEfcm.SelTimeout = curr.SelEfcm.SelTimeout
		cfg.SelEfcm.SelTimeout += 1000
		if cfg.SelEfcm.SelTimeout > 10000 {
			cfg.SelEfcm.SelTimeout = 1000
		}
		mutateMethod := getRandomWithMax(10)
		// get all select records we have seen so far for this executable
		records := g.GetAllSelectRecords()
		numOfSelects := len(records)
		if numOfSelects == 0 {
			return nil, nil
		}

		if mutateMethod < 8 {
			// Mutate one select per time
			mutateWhichSelect := getRandomWithMax(numOfSelects)
			numOfSelectCases := records[mutateWhichSelect].Cases
			if numOfSelectCases == 0 {
				return nil, fmt.Errorf("cannot randomly mutate an input with zero number of cases in select %d", mutateWhichSelect)
			}
			mutateToWhatValue := getRandomWithMax(int(numOfSelectCases))

			cfg.SelEfcm.Efcms = append(cfg.SelEfcm.Efcms, selefcm.SelEfcm{
				ID:   records[mutateWhichSelect].ID,
				Case: mutateToWhatValue,
			})
		} else {
			// Mutate random number of select
			mutateChance := getRandomWithMax(numOfSelects)
			for mutateIdx := 0; mutateIdx < mutateChance; mutateIdx++ {
				mutateWhichSelect := getRandomWithMax(numOfSelects)
				numOfSelectCases := records[mutateWhichSelect].Cases
				if numOfSelectCases == 0 {
					return nil, fmt.Errorf("cannot randomly mutate an input with zero number of cases in select %d", mutateWhichSelect)
				}
				mutateToWhatValue := getRandomWithMax(int(numOfSelectCases))
				cfg.SelEfcm.Efcms = append(cfg.SelEfcm.Efcms, selefcm.SelEfcm{
					ID:   records[mutateWhichSelect].ID,
					Case: mutateToWhatValue,
				})
			}
		}
		cfgs = append(cfgs, cfg)
	}

	return cfgs, nil
}

func getRandomWithMax(max int) int {
	mutateMethod, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		fmt.Println("Crypto/rand returned non-nil errors: ", err)
	}
	return int(mutateMethod.Int64())
}
