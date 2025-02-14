package task

import (
	"context"
	"go-inner/internal/utils"
	"math/rand"
	"sync"
)

const target = 5 * utils.GB
const times = 5 * utils.GB / utils.KB
const LargeObjectSize = 100000
const SmallObjectSize = 1

type AllocObjectTask struct {
	Workers    int
	ObjectSize int
}

var heap = [][]byte{}

func (t AllocObjectTask) Run(ctx context.Context) {
	sizePerGo := make([][]int, t.Workers)
	for i := 0; i < t.Workers; i++ {
		sum := 0
		size := []int{}
		for sum < (t.ObjectSize / t.Workers) {
			// random := rand.Intn(64 * utils.KB)
			random := 512 * utils.B

			size = append(size, random)
			sum += random
		}
		sizePerGo[i] = size
	}
	var wg sync.WaitGroup
	for i := 0; i < t.Workers; i++ {
		wg.Add(1)
		go func(ctx context.Context) {
			defer wg.Done()
			for _, val := range sizePerGo[i] {
				select {
				case <-ctx.Done():
					return
				default:
					Alloc(val)
				}
			}
		}(ctx)
	}
	wg.Wait()
}

func Alloc(size int) []byte {
	tmp := make([]byte, size)
	for i := 0; i < size; i++ {
		tmp[i] = 'a'
	}
	if (rand.Intn(2)+1)%2 != 0 {
		heap = append(heap, tmp)
	}
	return tmp
}

// func (t AllocObjectTask) Run(ctx context.Context) {
// 	for i := 0; i < times; i++ {
// 		select {
// 		case <-ctx.Done():
// 			return
// 		default:
// 			BatchAllocObject()
// 		}
// 	}
// }

// func BatchAllocObject() []byte {
// 	random := rand.Intn(1 * utils.KB)
// 	Num += uint64(random)
// 	tmp := make([]byte, random)
// 	for i := 0; i < random; i++ {
// 		tmp[i] = 'a'
// 	}
// 	if random%(rand.Intn(5)+1) == 0 {
// 		heap = append(heap, tmp)
// 	}
// 	return tmp
// }
