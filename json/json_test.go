package json

import (
	"encoding/json"
	"reflect"
	"testing"
)

type AnyStruct struct {
	Member any `json:"member"`
}

// any 类型在经过序列化和反序列化后,不会变为之前的类型
// 如果之前是结构体,那么会反序列为 map
func TestMarshalAny(t *testing.T) {
	a1 := AnyStruct{Member: int(1)}
	s, _ := json.Marshal(a1)
	var a2 AnyStruct
	_ = json.Unmarshal(s, &a2)
	_, ok := a2.Member.(int)
	t.Log(ok)
	t.Log(reflect.TypeOf(a2.Member))
}
