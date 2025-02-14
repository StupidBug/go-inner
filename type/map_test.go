package type_

import "testing"

func TestReadMapWithNoKey(t *testing.T) {
	m := make(map[string]string)
	if m["a"] != "" {
		t.Fatalf("not expected")
	}
	m1 := make(map[string]struct{})
	if m1["a"] != struct{}{} {
		t.Fatalf("not expected")
	}
}
