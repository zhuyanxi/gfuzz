package oraclert

import (
	"runtime"
	"sync/atomic"
	"time"
)

func GetSelEfcmSwitchCaseIdx(filename string, origLine string, origCases int) int {
	runtime.StoreLastMySwitchSelectNumCase(origCases)
	runtime.StoreLastMySwitchLineNum(origLine)

	if efcmStrat == nil {
		// if strategy is not initialized, return -1
		runtime.StoreLastMySwitchChoice(-1)
		return -1
	}
	selectID := filename + ":" + origLine
	idx := efcmStrat.GetCase(selectID)
	//fmt.Printf("[oraclert] index %d is chosen for %s\n", idx, selectID)
	if idx != -1 {
		runtime.StoreLastMySwitchChoice(idx)
		return idx
	} else {
		atomic.AddUint32(&notSelEfcmCount, 1)
		runtime.StoreLastMySwitchChoice(-1)
		return -1 // let switch choose the default case
	}
}

func StoreLastMySwitchChoice(choice int) {
	if choice == -1 {
		atomic.AddUint32(&origSelCount, 1)
	}
	runtime.StoreLastMySwitchChoice(choice)
}

func SelEfcmTimeout() <-chan time.Time {
	// if this channel wins, remember to call "runtime.StoreLastMySwitchChoice(-1)", which means we will use the original select
	return time.After(time.Duration(selTimeout) * time.Millisecond)
}
