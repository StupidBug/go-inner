package type_

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestBoolSize(t *testing.T) {
	fmt.Printf("size of bool: %d bytes", unsafe.Sizeof(bool(true))) // size of int: 1 bytes
}
