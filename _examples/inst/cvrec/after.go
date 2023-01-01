package cvrec

import (
	"sync"

	oraclert "github.com/zhuyanxi/gfuzz/pkg/oraclert"
)

func Hello() {
	m := sync.Mutex{}

	c := sync.NewCond(&m)
	oraclert.StoreOpInfo("Broadcast", 1)

	c.Broadcast()
	oraclert.StoreOpInfo("Signal", 2)

	c.Signal()
	oraclert.StoreOpInfo("Wait", 3)

	c.Wait()
}
