package type_

import (
	"reflect"
	"testing"
)

// 字符串类型的切片的类型还是字符串
func TestStringSliceType(t *testing.T) {
	a := "123"
	t.Log(a[2:] == "3")          //true
	t.Log(reflect.TypeOf(a[2:])) //string
}
