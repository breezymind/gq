package gq

import (
	"fmt"
	"reflect"
	"testing"
)

type TestStruct struct {
	Name     string        `json:"name"`
	Age      int           `json:"age"`
	Messages []interface{} `json:"messages"`
	Gender   string        `json:"gender"`
}

// NOTE: Test
func Test_Struct2JSONByte(t *testing.T) {
	teststruct := &TestStruct{
		Name:     "Karl",
		Age:      38,
		Messages: []interface{}{"Hello"},
		Gender:   "Male",
	}
	t.Log(teststruct)
	t.Log(Struct2JSONByte(teststruct))
}

func Test_Struct2JSONString(t *testing.T) {
	teststruct := &TestStruct{
		Name:     "Karl",
		Age:      38,
		Messages: []interface{}{"Hello"},
		Gender:   "Male",
	}
	t.Log(teststruct)
	t.Log(Struct2JSONString(teststruct))
}

func Test_Map2Struct(t *testing.T) {
	testmap := map[string]interface{}{
		"Name":     "Karl",
		"Age":      38,
		"Messages": []string{"Hello"},
		"Gender":   "Male",
		"OnlyMap":  "OnlyMap",
	}
	refts := &TestStruct{}

	Map2Struct(testmap, refts)

	t.Log(refts)
	t.Log(refts.Name)
	t.Log(refts.Messages)
	t.Log(reflect.TypeOf(refts))
}
func Test_JSONString2Struct(t *testing.T) {
	strjson := "{\"name\":\"Karl\",\"age\":38,\"messages\":[\"Hello\"],\"gender\":\"Male\"}"
	refts := &TestStruct{}

	JSONString2Struct(strjson, refts)

	t.Log(refts)
	t.Log(refts.Name)
	t.Log(reflect.TypeOf(refts))
}
func Test_JSONByte2Struct(t *testing.T) {
	bytejson := []byte("{\"name\":\"Karl\",\"age\":38,\"messages\":[\"Hello\"],\"gender\":\"Male\"}")
	refts := &TestStruct{}

	JSONByte2Struct(bytejson, refts)

	t.Log(refts)
	t.Log(refts.Name)
	t.Log(reflect.TypeOf(refts))
}

func Test_InterfaceSlice2StringSlice(t *testing.T) {
	infslice := []interface{}{"Hello", "World"}
	t.Log(reflect.TypeOf(infslice))

	strslice := InterfaceSlice2StringSlice(infslice)
	t.Log(strslice)
	t.Log(reflect.TypeOf(strslice))
}

// NOTE: Example
func ExampleStruct2JSONByte() {
	teststruct := &TestStruct{
		Name:     "Karl",
		Age:      38,
		Messages: []interface{}{"Hello"},
		Gender:   "Male",
	}
	fmt.Println(teststruct)
	fmt.Println(Struct2JSONByte(teststruct))
	// Output:
	// &{Karl 38 [Hello] Male}
	// [123 34 110 97 109 101 34 58 34 75 97 114 108 34 44 34 97 103 101 34 58 51 56 44 34 109 101 115 115 97 103 101 115 34 58 91 34 72 101 108 108 111 34 93 44 34 103 101 110 100 101 114 34 58 34 77 97 108 101 34 125]
}

func ExampleStruct2JSONString() {
	teststruct := &TestStruct{
		Name:     "Karl",
		Age:      38,
		Messages: []interface{}{"Hello"},
		Gender:   "Male",
	}
	fmt.Println(teststruct)
	fmt.Println(Struct2JSONString(teststruct))
	// Output:
	// &{Karl 38 [Hello] Male}
	// {"name":"Karl","age":38,"messages":["Hello"],"gender":"Male"}
}

func ExampleMap2Struct() {
	testmap := map[string]interface{}{
		"Name":     "Karl",
		"Age":      38,
		"Messages": []string{"Hello"},
		"Gender":   "Male",
		"OnlyMap":  "OnlyMap",
	}
	refts := &TestStruct{}

	Map2Struct(testmap, refts)

	fmt.Println(refts)
	fmt.Println(refts.Name)
	fmt.Println(refts.Messages)
	fmt.Println(reflect.TypeOf(refts))

	// Output:
	// &{Karl 38 [Hello] Male}
	// Karl
	// [Hello]
	// *gq.TestStruct
}
func ExampleJSONString2Struct() {
	strjson := "{\"name\":\"Karl\",\"age\":38,\"messages\":[\"Hello\"],\"gender\":\"Male\"}"
	refts := &TestStruct{}

	JSONString2Struct(strjson, refts)

	fmt.Println(refts)
	fmt.Println(refts.Name)
	fmt.Println(reflect.TypeOf(refts))

	// Output:
	// &{Karl 38 [Hello] Male}
	// Karl
	// *gq.TestStruct
}
func ExampleJSONByte2Struct() {
	bytejson := []byte("{\"name\":\"Karl\",\"age\":38,\"messages\":[\"Hello\"],\"gender\":\"Male\"}")
	refts := &TestStruct{}

	JSONByte2Struct(bytejson, refts)

	fmt.Println(refts)
	fmt.Println(refts.Name)
	fmt.Println(reflect.TypeOf(refts))

	// Output:
	// &{Karl 38 [Hello] Male}
	// Karl
	// *gq.TestStruct
}

func ExampleInterfaceSlice2StringSlice() {
	infslice := []interface{}{"Hello", "World"}
	fmt.Println(reflect.TypeOf(infslice))

	strslice := InterfaceSlice2StringSlice(infslice)
	fmt.Println(strslice)
	fmt.Println(reflect.TypeOf(strslice))

	// Output:
	// []interface {}
	// [Hello World]
	// []string
}
