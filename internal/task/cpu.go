package task

import (
	"context"
	"math"
)

type Task interface {
	Run(context.Context)
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

type CPUTask struct {
}

func (t CPUTask) Run(ctx context.Context) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			isPrime(i)
		}
	}
}
