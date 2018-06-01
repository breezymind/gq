package gq

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_RSortByValue(t *testing.T) {
	strjson := "{\"banana\":999,\"apple\":38}"
	testmap := NewMapByJSONByte([]byte(strjson))

	t.Log(testmap.GetJSONString())

	kvslice := testmap.ToKeyValueSlice()

	t.Log(reflect.TypeOf(kvslice))
	t.Log(kvslice.GetJSONString())

	kvslice.RSortByValue()
	t.Log("kvslice.SortByValue() after")
	t.Log(kvslice.GetJSONString())
	t.Log(kvslice.GetJSONPretty())
}

func ExampleRSortByValue() {
	strjson := "{\"banana\":999,\"apple\":38}"
	testmap := NewMapByJSONByte([]byte(strjson))

	kvslice := testmap.ToKeyValueSlice()

	fmt.Println(reflect.TypeOf(kvslice))
	fmt.Println(kvslice.GetJSONString())

	kvslice.RSortByValue()
	fmt.Println("kvslice.SortByValue() after")
	fmt.Println(kvslice.GetJSONString())
	fmt.Println(kvslice.GetJSONPretty())

	// Output:
	// {"apple":38,"banana":999}
	//
	// gq.KVSlice
	// [{"k":"apple","v":38},{"k":"banana","v":999}]
	//
	// kvslice.SortByValue() after
	// [{"k":"banana","v":999},{"k":"apple","v":38}]
}
