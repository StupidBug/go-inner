package time

import (
	"context"
	"time"

	"go-inner/internal/task"
)

func ScheduleWithSleep(task task.Task, ctx context.Context) {
	for {
		task.Run(ctx)
		time.Sleep(time.Second)
	}
}

func ScheduleWithTick(task task.Task, ctx context.Context) {
	t1 := time.Tick(3 * time.Second)
	for {
		select {
		case <-t1:
			task.Run(ctx)
		}
	}
}

func ScheduleWithTicker(task task.Task, ctx context.Context) {
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			task.Run(ctx)
			t.Stop()
		}
	}
}
