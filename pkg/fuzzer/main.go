package fuzzer

import (
	"context"
	"log"
	"os"
	"sync/atomic"

	"github.com/zhuyanxi/gfuzz/pkg/fuzz/api"
	"github.com/zhuyanxi/gfuzz/pkg/fuzz/config"
	"github.com/zhuyanxi/gfuzz/pkg/fuzz/gexecfuzz"
	"github.com/zhuyanxi/gfuzz/pkg/gexec"
	ortconfig "github.com/zhuyanxi/gfuzz/pkg/oraclert/config"
)

var (
	runningTasks int32
	exited       uint32
)

// Reply run the fuzzing with given oracle runtime configuration and given executable
func Replay(fctx *api.Context, ge gexec.Executable, config *config.Config, rtConfig *ortconfig.Config) {
	ctx := context.Background()
	i := api.NewExecInput(fctx.GetAutoIncGlobalID(), 0, config.OutputDir, ge, rtConfig, api.ReplayStage)
	o, err := Run(ctx, fctx.Cfg, i)
	if err != nil {
		log.Printf("%s: %s", i.ID, err)
	}
	err = HandleExec(ctx, i, o, fctx, nil)
	if err != nil {
		log.Printf("%s: %s", i.ID, err)
	}
}

// Main starts fuzzing with a given list of executables and configuration
func Main(fctx *api.Context, execs []gexec.Executable, config *config.Config,
	interestHdl api.InterestHandler, scorer api.ScoreStrategy) {
	if len(execs) == 0 {
		log.Println("no executables found, exit.")
		os.Exit(0)
	}

	for _, e := range execs {
		log.Printf("found executable: %s", e)
	}

	// initialize interested inputs by generating init stage input for each executables
	fctx.EachGExecFuzz(func(g *gexecfuzz.GExecFuzz) {
		i := api.NewInitExecInput(fctx, g.Exec)
		ii := api.NewUnexecutedInterestInput(i)
		ii.Reason = api.InitStg
		fctx.Interests.Add(ii)
	})

	exitCh := make(chan struct{})
	// endless loop to handle interested inputs
	go func() {
		TryLoopInterestList(fctx, interestHdl, exitCh)
	}()

	// start a group of workers to handle fuzz execution in parallel
	startWorkers(config.MaxParallel, func(ctx context.Context) {
		execWorker(ctx, fctx, interestHdl, exitCh)
	})

}

func TryLoopInterestList(fctx *api.Context, interestHdl api.InterestHandler, exitCh chan struct{}) {
	if fctx.Interests.IsLooping() || atomic.LoadInt32(&runningTasks) > 0 {
		return
	}

	handled := fctx.Interests.Each(interestHdl)

	if atomic.LoadInt32(&runningTasks) == 0 && !fctx.Interests.Dirty && !handled {
		if atomic.LoadUint32(&exited) == 0 {
			atomic.StoreUint32(&exited, 1)
			log.Printf("nothing to fuzz, exiting...")
			close(exitCh)
		}
	}
}

// execWorker handles a execution inputs from channel
func execWorker(ctx context.Context,
	fc *api.Context,
	interestHdl api.InterestHandler,
	exitCh chan struct{}) {
	logger := getWorkerLogger(ctx)

	for {
		select {
		case i := <-fc.ExecInputCh:
			logger.Printf("received %s", i.ID)
			atomic.AddInt32(&runningTasks, 1)
			o, err := Run(ctx, fc.Cfg, i)
			if err != nil {
				logger.Printf("%s: %s", i.ID, err)
			}
			fc.IncNumOfRun()
			err = HandleExec(ctx, i, o, fc, interestHdl)
			if err != nil {
				logger.Printf("%s: %s", i.ID, err)
			}
			logger.Printf("finished %s", i.ID)

			atomic.AddInt32(&runningTasks, -1)

			TryLoopInterestList(fc, interestHdl, exitCh)
		case <-exitCh:
			logger.Printf("exited")
			return
		}

	}
}
