package interest

import (
	"testing"

	"github.com/zhuyanxi/gfuzz/pkg/oraclert/config"
	"github.com/zhuyanxi/gfuzz/pkg/selefcm"
	"github.com/zhuyanxi/gfuzz/pkg/utils/hash"
)

func TestCfgHashEq(t *testing.T) {
	cfg1 := &config.Config{
		SelEfcm: selefcm.SelEfcmConfig{
			SelTimeout: 500,
			Efcms: []selefcm.SelEfcm{
				{
					ID:   "abc.go:123",
					Case: 1,
				},
			},
		},
	}
	cfg2 := &config.Config{
		SelEfcm: selefcm.SelEfcmConfig{
			SelTimeout: 500,
			Efcms: []selefcm.SelEfcm{
				{
					ID:   "abc.go:123",
					Case: 1,
				},
			},
		},
	}
	if hash.AsSha256(cfg1) != hash.AsSha256(cfg2) {
		t.Fail()
	}
}

func TestCfgHashNotEq(t *testing.T) {
	cfg1 := &config.Config{
		SelEfcm: selefcm.SelEfcmConfig{
			SelTimeout: 1000,
			Efcms: []selefcm.SelEfcm{
				{
					ID:   "abc.go:123",
					Case: 1,
				},
			},
		},
	}
	cfg2 := &config.Config{
		SelEfcm: selefcm.SelEfcmConfig{
			SelTimeout: 500,
			Efcms: []selefcm.SelEfcm{
				{
					ID:   "abc.go:123",
					Case: 1,
				},
			},
		},
	}
	if hash.AsSha256(cfg1) == hash.AsSha256(cfg2) {
		t.Fail()
	}
}
