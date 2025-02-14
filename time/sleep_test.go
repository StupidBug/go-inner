package time

import (
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	time.Sleep(1 * time.Second)
	time.Sleep(1 * time.Second)
}
