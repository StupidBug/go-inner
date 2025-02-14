package task

import (
	"context"
	"sync"
)

func RunWithGoroutine(task Task, n int, ctx context.Context) *sync.WaitGroup {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			task.Run(ctx)
		}()
	}
	return &wg
}
