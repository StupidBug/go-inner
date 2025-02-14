package grammar

import "testing"

func testdefer1() (a int) {
	defer func() { a = 1 }()
	return
}

func testdefer2() int {
	a := 0
	defer func() { a = 1 }()
	return a
}

func TestDeferExecute(t *testing.T) {
	t.Logf("test defer 1: %d \n", testdefer1()) // test defer 1: 1
	t.Logf("test defer 2: %d \n", testdefer2()) // test defer 2: 0
}
