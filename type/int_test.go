package type_

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestIntSize(t *testing.T) {
	fmt.Printf("size of int: %d bytes", unsafe.Sizeof(int(0))) // size of int: 8 bytes
}
