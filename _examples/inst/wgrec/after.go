package wgrec

import (
	"sync"

	oraclert "github.com/zhuyanxi/gfuzz/pkg/oraclert"
)

type RandomStruct struct {
	wg sync.WaitGroup
}

func Hello() {
	wg := sync.WaitGroup{}
	r := RandomStruct{}
	oraclert.StoreOpInfo("Add", 1)
	wg.Add(1)
	oraclert.StoreOpInfo("Wait", 2)
	wg.Wait()
	oraclert.StoreOpInfo("Done", 3)
	wg.Done()
	oraclert.StoreOpInfo("Wait", 4)

	r.wg.Wait()
}
