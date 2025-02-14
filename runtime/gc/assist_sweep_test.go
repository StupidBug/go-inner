package gc

import (
	"context"
	"go-inner/internal/task"
	"go-inner/internal/utils"
	"runtime"
	"sync"
	"testing"
)

func AsyncTriggerGC() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		runtime.GC()
	}()
	return &wg
}

const target = 50 * utils.GB

// Total: 13.5 s
// GC incremental sweep: 315 ms
// GC mark assist: 40 ms
// STW: 0.4 ms
func TestNoBgsweep(t *testing.T) {
	ctx := context.Background()
	task.AllocObjectTask{Workers: utils.GOMAXPROCS, ObjectSize: target}.Run(ctx)
	utils.PrintMemstats()
}

// Total: 12.9 s
// GC incremental sweep: 33 ms
// GC mark assist: 40 ms
// STW: 0.8 ms
func TestBgsweep(t *testing.T) {
	ctx := context.Background()
	task.AllocObjectTask{Workers: utils.GOMAXPROCS - 1, ObjectSize: target}.Run(ctx)
	utils.PrintMemstats()
}
