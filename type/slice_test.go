package type_

import (
	"fmt"
	"go-inner/internal/debug"
	"go-inner/internal/utils"
	"testing"
)

func TestSliceNil(t *testing.T) {
	a := []byte{}
	fmt.Println(nil == a) // false
	fmt.Println(len(a))   // 0
	var b []byte
	fmt.Println(nil == b) // true
	fmt.Println(len(b))   // 0
	c := new([]byte)
	fmt.Println(nil == *c) // true
	fmt.Println(len(*c))   // 0
}

type LargeSlice struct {
	slice [utils.KB * 128]byte
}

type SmallSlice struct {
	slice [utils.KB * 1]byte
}

// go test -gcflags="-m" type/slice_test.go
// go test -gcflags='all=-N -l -S' -trimpath type/slice_test.go 2>&1 | grep -e "slice_test"
func TestSliceEscape(t *testing.T) {
	_ = make([]byte, utils.KB*64)           // make([]byte, 65536) does not escape
	_ = make([]byte, utils.KB*128)          // make([]byte, 131072) escapes to heap
	_ = new(LargeSlice)                     // new(LargeSlice) escapes to heap
	_ = new(SmallSlice)                     // new(LargeSlice) escapes to heap
	debug.DoNothing([utils.KB * 128]byte{}) // [131072]byte{} does not escape
	var a [utils.KB * 128]byte              //
	debug.DoNothing(a)                      // a does not escape
}
