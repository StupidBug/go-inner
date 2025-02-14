package task

import (
	"go-inner/internal/utils"
	"math/rand"
	"testing"
)

func TestAllocLargeObject(t *testing.T) {
	sum := 0
	for i := 0; i < times; i++ {
		sum += rand.Intn(utils.KB)
	}
	print(sum)
}
