package json

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestUnmarshalSlice(t *testing.T) {
	var d struct{ Data []string }
	assert.True(t, nil == d.Data)

	json.Unmarshal([]byte{}, &d)
	assert.True(t, nil == d.Data)

	json.Unmarshal([]byte("{\"Data\":[\"1\"]}"), &d)
	assert.False(t, nil == d.Data)

	json.Unmarshal([]byte("{\"Data\":[]}"), &d)
	assert.False(t, nil == d.Data)
}
