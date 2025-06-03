package type_

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestByteSlice(t *testing.T) {
	a := "1"
	s, _ := json.Marshal(a)
	fmt.Printf("dfas: %s", s)
}