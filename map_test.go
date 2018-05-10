package gq

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_NewMapByJSONByte(t *testing.T) {
	strjson := "{\"name\":\"Karl\",\"age\":38,\"messages\":[\"Hello\"],\"gender\":\"Male\"}"
	newmap := NewMapByJSONByte([]byte(strjson))
	t.Log(newmap.GetJSONString())
}

func Test_NewMapByStruct(t *testing.T) {
	teststruct := &TestStruct{
		Name:     "Karl",
		Age:      38,
		Messages: []interface{}{"Hello"},
		Gender:   "Male",
	}
	newmap := NewMapByStruct(teststruct)
	t.Log(newmap.GetJSONString())
}

func Test_GetJSONPretty(t *testing.T) {
	strjson := "{\"name\":\"Karl\",\"age\":38,\"messages\":[\"Hello\"],\"gender\":\"Male\"}"
	newmap := NewMapByJSONByte([]byte(strjson))
	t.Log(newmap.GetJSONPretty())
}

func Test_GetJSONString(t *testing.T) {
	strjson := "{\"name\":\"Karl\",\"age\":38,\"messages\":[\"Hello\"],\"gender\":\"Male\"}"
	newmap := NewMapByJSONByte([]byte(strjson))
	t.Log(newmap.GetJSONString())
}

func Test_GetJSONByte(t *testing.T) {
	strjson := "{\"name\":\"Karl\",\"age\":38,\"messages\":[\"Hello\"],\"gender\":\"Male\"}"
	newmap := NewMapByJSONByte([]byte(strjson))
	t.Log(newmap.GetJSONByte())
}

func Test_SetJSONString(t *testing.T) {
	strjson := "{\"name\":\"Karl\",\"age\":38}"
	testmap := NewMapByJSONByte([]byte(strjson))
	t.Log(testmap.GetJSONString())
	testmap.SetJSONString("{\"name\":\"Tomas\",\"age\":20}")
	t.Log(testmap.GetJSONString())
}

func Test_SetJSONByte(t *testing.T) {
	strjson := "{\"name\":\"Karl\",\"age\":38}"
	testmap := NewMapByJSONByte([]byte(strjson))
	t.Log(testmap.GetJSONString())
	testmap.SetJSONByte([]byte("{\"name\":\"Tomas\",\"age\":20}"))
	t.Log(testmap.GetJSONString())
}

func Test_SetStruct(t *testing.T) {
	strjson := "{\"name\":\"Karl\",\"age\":38,\"messages\":[\"Hello\"],\"gender\":\"Male\"}"
	testmap := NewMapByJSONByte([]byte(strjson))
	t.Log(testmap.GetJSONString())

	newstruct := &TestStruct{
		Name:     "Tomas",
		Age:      20,
		Messages: []interface{}{"Hello", "World"},
		Gender:   "Male",
	}
	testmap.SetStruct(newstruct)
	t.Log(testmap.GetJSONString())
}

func Test_SetAttr(t *testing.T) {
	newstruct := &TestStruct{
		Name:     "Tomas",
		Age:      20,
		Messages: []interface{}{"Hello", "World"},
		Gender:   "Male",
	}
	testmap := NewMapByStruct(newstruct)
	t.Log(testmap.GetJSONString())

	testmap.SetAttr("attr_int", 3)
	t.Log(testmap.GetJSONString())

	anonystruct := struct {
		Friends []*TestStruct `json:"friends"`
	}{
		[]*TestStruct{
			&TestStruct{
				Name:     "Tomas",
				Age:      20,
				Messages: []interface{}{"Hello", "World"},
				Gender:   "Male",
			},
		},
	}

	testmap.SetAttr("attr_struct", anonystruct)
	t.Log(testmap.GetJSONPretty())
}

func Test_SetAttrMap(t *testing.T) {
	newstruct := &TestStruct{
		Name:     "Tomas",
		Age:      20,
		Messages: []interface{}{"Hello", "World"},
		Gender:   "Male",
	}
	testmap := NewMapByStruct(newstruct)
	t.Log(testmap.GetJSONString())

	testmap.SetAttrMap(
		"misc",
		NewMapByJSONByte([]byte("{\"points\":[1,2,3,4],\"name\":\"cacao\"}")),
	)

	t.Log(testmap.GetJSONPretty())
}

func Test_SetAttrJSONString(t *testing.T) {
	newstruct := &TestStruct{
		Name:     "Tomas",
		Age:      20,
		Messages: []interface{}{"Hello", "World"},
		Gender:   "Male",
	}
	testmap := NewMapByStruct(newstruct)
	t.Log(testmap.GetJSONString())

	testmap.SetAttrJSONString(
		"misc",
		"{\"points\":[1,2,3,4],\"name\":\"cacao\"}",
	)

	t.Log(testmap.GetJSONPretty())
}

func Test_SetAttrJSONByte(t *testing.T) {
	newstruct := &TestStruct{
		Name:     "Tomas",
		Age:      20,
		Messages: []interface{}{"Hello", "World"},
		Gender:   "Male",
	}
	testmap := NewMapByStruct(newstruct)
	t.Log(testmap.GetJSONString())

	testmap.SetAttrJSONByte(
		"misc",
		[]byte("{\"points\":[1,2,3,4],\"name\":\"cacao\"}"),
	)

	t.Log(testmap.GetJSONPretty())
}

func Test_SetAttrQuery(t *testing.T) {
	strjson := "{\"misc\":{\"name\":\"cacao\"},\"name\":\"Tomas\"}"

	testmap := NewMapByJSONByte([]byte(strjson))

	testmap.SetAttrQuery("misc.name", "cocoa")
	t.Log(testmap.GetJSONString())

	testmap.SetAttrQuery("misc.name", 3)
	t.Log(testmap.GetJSONString())

	anonystruct := struct {
		AttrA string `json:"attr_a"`
		AttrB string `json:"attr_b"`
	}{
		AttrA: "TestA",
		AttrB: "TestB",
	}

	testmap.SetAttrQuery("misc.name", anonystruct)
	t.Log(testmap.GetJSONPretty())
}

func Test_IsExistAttr(t *testing.T) {
	newstruct := &TestStruct{
		Name:     "Tomas",
		Age:      20,
		Messages: []interface{}{"Hello", "World"},
		Gender:   "Male",
	}
	testmap := NewMapByStruct(newstruct)
	t.Log(testmap.GetJSONString())
	t.Log(testmap.IsExistAttr("age"))
}

func Test_GetAttrMap(t *testing.T) {
	strjson := "{\"misc\":{\"name\":\"cacao\"},\"name\":\"Tomas\"}"
	testmap := NewMapByJSONByte([]byte(strjson))

	miscmap := testmap.GetAttrMap("misc")
	t.Log(miscmap.GetJSONString())

	t.Log(miscmap.GetAttrString("name"))
}

func Test_GetAttrQuery(t *testing.T) {
	strjson := "{\"shop\":{\"name\":{\"first\":\"hollys\",\"last\":\"coffee\"},\"location\":[\"a\",\"b\"]}}"
	testmap := NewMapByJSONByte([]byte(strjson))
	t.Log(testmap.GetJSONPretty())

	t.Log(testmap.GetAttrQuery("shop.name.first"))
	t.Log(reflect.TypeOf(testmap.GetAttrQuery("shop.name.first")))

	t.Log(testmap.GetAttrQuery("shop.location"))
	t.Log(reflect.TypeOf(testmap.GetAttrQuery("shop.location")))
}

func Test_GetAttrInterface(t *testing.T) {
	strjson := "{\"shop\":{\"name\":{\"first\":\"hollys\",\"last\":\"coffee\"},\"location\":[\"a\",\"b\"]}}"
	testmap := NewMapByJSONByte([]byte(strjson))

	t.Log(testmap.GetAttrInterface("shop"))
	t.Log(reflect.TypeOf(testmap.GetAttrInterface("shop")))
}

func Test_GetMapInterface(t *testing.T) {
	strjson := "{\"shop\":{\"name\":{\"first\":\"hollys\",\"last\":\"coffee\"},\"location\":[\"a\",\"b\"]}}"
	testmap := NewMapByJSONByte([]byte(strjson))

	t.Log(testmap.GetMapInterface())
	t.Log(reflect.TypeOf(testmap.GetMapInterface()))
}

func Test_GetAttrInt(t *testing.T) {
	strjson := "{\"espresso\": 1.33222, \"americano\": \"1234\", \"latte\": \"2.234\" }"
	testmap := NewMapByJSONByte([]byte(strjson))

	t.Log(testmap.GetJSONPretty())

	t.Log(testmap.GetAttrInt("espresso"))
	t.Log(testmap.GetAttrInt("americano"))
	t.Log(testmap.GetAttrInt("latte"))
}

func Test_GetAttrString(t *testing.T) {
	strjson := "{\"espresso\": 1.33222, \"americano\": \"1234\"}"
	testmap := NewMapByJSONByte([]byte(strjson))

	t.Log(testmap.GetJSONPretty())

	t.Log(testmap.GetAttrString("espresso"))
	t.Log(testmap.GetAttrString("americano"))
}

func Test_GetAttrSlice(t *testing.T) {
	strjson := "{\"coffee\": [\"americano\",\"espresso\",\"latte\"]}"
	testmap := NewMapByJSONByte([]byte(strjson))

	t.Log(testmap.GetJSONPretty())

	for i, j := range testmap.GetAttrStringSlice("coffee") {
		t.Logf("%d : %s", i, j)
	}
}

// func Test_DelAttr(k string) *Map {
// }

// func Test_Keys() []string {
// }

// func Test_Values() []interface{} {
// }

// func Test_Clone() *Map {
// }

func ExampleNewMapByJSONByte() {
	strjson := "{\"name\":\"Karl\",\"age\":38,\"messages\":[\"Hello\"],\"gender\":\"Male\"}"
	newmap := NewMapByJSONByte([]byte(strjson))
	fmt.Println(newmap.GetJSONString())
}

func ExampleNewMapByStruct() {
	teststruct := &TestStruct{
		Name:     "Karl",
		Age:      38,
		Messages: []interface{}{"Hello"},
		Gender:   "Male",
	}
	newmap := NewMapByStruct(teststruct)
	fmt.Println(newmap.GetJSONString())
}

func ExampleMap_GetJSONPretty() {
	strjson := "{\"name\":\"Karl\",\"age\":38,\"messages\":[\"Hello\"],\"gender\":\"Male\"}"
	newmap := NewMapByJSONByte([]byte(strjson))
	fmt.Println(newmap.GetJSONPretty())
	// Output:
	// {
	// 	"age": 38,
	// 	"gender": "Male",
	// 	"messages": [
	// 		"Hello"
	// 	],
	// 	"name": "Karl"
	// }
}

func ExampleMap_GetJSONString() {
	strjson := "{\"name\":\"Karl\",\"age\":38,\"messages\":[\"Hello\"],\"gender\":\"Male\"}"
	newmap := NewMapByJSONByte([]byte(strjson))
	fmt.Println(newmap.GetJSONString())
	// Output:
	// {"age":38,"gender":"Male","messages":["Hello"],"name":"Karl"}
}

func ExampleMap_GetJSONByte() {
	strjson := "{\"name\":\"Karl\",\"age\":38,\"messages\":[\"Hello\"],\"gender\":\"Male\"}"
	newmap := NewMapByJSONByte([]byte(strjson))
	fmt.Println(newmap.GetJSONByte())
	// Output:
	// [123 34 97 103 101 34 58 51 56 44 34 103 101 110 100 101 114 34 58 34 77 97 108 101 34 44 34 109 101 115 115 97 103 101 115 34 58 91 34 72 101 108 108 111 34 93 44 34 110 97 109 101 34 58 34 75 97 114 108 34 125]
}

func ExampleMap_SetJSONString() {
	strjson := "{\"name\":\"Karl\",\"age\":38}"
	testmap := NewMapByJSONByte([]byte(strjson))
	fmt.Println(testmap.GetJSONString())
	testmap.SetJSONString("{\"name\":\"Tomas\",\"age\":20}")
	fmt.Println(testmap.GetJSONString())
	// Output:
	// {"age":38,"name":"Karl"}
	// {"age":20,"name":"Tomas"}
}

func ExampleMap_SetJSONByte() {
	strjson := "{\"name\":\"Karl\",\"age\":38}"
	testmap := NewMapByJSONByte([]byte(strjson))
	fmt.Println(testmap.GetJSONString())
	testmap.SetJSONByte([]byte("{\"name\":\"Tomas\",\"age\":20}"))
	fmt.Println(testmap.GetJSONString())
	// Output:
	// {"age":38,"name":"Karl"}
	// {"age":20,"name":"Tomas"}
}

func ExampleMap_SetStruct() {
	strjson := "{\"name\":\"Karl\",\"age\":38,\"messages\":[\"Hello\"],\"gender\":\"Male\"}"
	testmap := NewMapByJSONByte([]byte(strjson))
	fmt.Println(testmap.GetJSONString())

	newstruct := &TestStruct{
		Name:     "Tomas",
		Age:      20,
		Messages: []interface{}{"Hello", "World"},
		Gender:   "Male",
	}
	testmap.SetStruct(newstruct)
	fmt.Println(testmap.GetJSONString())
	// Output:
	// {"age":38,"gender":"Male","messages":["Hello"],"name":"Karl"}
	// {"age":20,"gender":"Male","messages":["Hello","World"],"name":"Tomas"}
}

func ExampleMap_SetAttr() {
	newstruct := &TestStruct{
		Name:     "Tomas",
		Age:      20,
		Messages: []interface{}{"Hello", "World"},
		Gender:   "Male",
	}
	testmap := NewMapByStruct(newstruct)
	fmt.Println(testmap.GetJSONString())

	testmap.SetAttr("attr_int", 3)
	fmt.Println(testmap.GetJSONString())

	anonystruct := struct {
		Friends []*TestStruct `json:"friends"`
	}{
		[]*TestStruct{
			&TestStruct{
				Name:     "Tomas",
				Age:      20,
				Messages: []interface{}{"Hello", "World"},
				Gender:   "Male",
			},
		},
	}

	testmap.SetAttr("attr_struct", anonystruct)
	fmt.Println(testmap.GetJSONPretty())

	// Output:
	// {"age":20,"gender":"Male","messages":["Hello","World"],"name":"Tomas"}
	// {"age":20,"attr_int":3,"gender":"Male","messages":["Hello","World"],"name":"Tomas"}
	// {
	// 	"age": 20,
	// 	"attr_int": 3,
	// 	"attr_struct": {
	// 		"friends": [
	// 			{
	// 				"name": "Tomas",
	// 				"age": 20,
	// 				"messages": [
	// 					"Hello",
	// 					"World"
	// 				],
	// 				"gender": "Male"
	// 			}
	// 		]
	// 	},
	// 	"gender": "Male",
	// 	"messages": [
	// 		"Hello",
	// 		"World"
	// 	],
	// 	"name": "Tomas"
	// }
}

func ExampleMap_SetAttrMap() {
	newstruct := &TestStruct{
		Name:     "Tomas",
		Age:      20,
		Messages: []interface{}{"Hello", "World"},
		Gender:   "Male",
	}
	testmap := NewMapByStruct(newstruct)
	fmt.Println(testmap.GetJSONString())

	testmap.SetAttrMap(
		"misc",
		NewMapByJSONByte([]byte("{\"points\":[1,2,3,4],\"name\":\"cacao\"}")),
	)

	fmt.Println(testmap.GetJSONPretty())

	// Output:
	// {"age":20,"gender":"Male","messages":["Hello","World"],"name":"Tomas"}
	// {
	// 	"age": 20,
	// 	"gender": "Male",
	// 	"messages": [
	// 		"Hello",
	// 		"World"
	// 	],
	// 	"misc": {
	// 		"name": "cacao",
	// 		"points": [
	// 			1,
	// 			2,
	// 			3,
	// 			4
	// 		]
	// 	},
	// 	"name": "Tomas"
	// }
}

func ExampleMap_SetAttrJSONString() {
	newstruct := &TestStruct{
		Name:     "Tomas",
		Age:      20,
		Messages: []interface{}{"Hello", "World"},
		Gender:   "Male",
	}
	testmap := NewMapByStruct(newstruct)
	fmt.Println(testmap.GetJSONString())

	testmap.SetAttrJSONString(
		"misc",
		"{\"points\":[1,2,3,4],\"name\":\"cacao\"}",
	)
	fmt.Println(testmap.GetJSONPretty())

	// Output:
	// {"age":20,"gender":"Male","messages":["Hello","World"],"name":"Tomas"}
	// {
	// 	"age": 20,
	// 	"gender": "Male",
	// 	"messages": [
	// 		"Hello",
	// 		"World"
	// 	],
	// 	"misc": {
	// 		"name": "cacao",
	// 		"points": [
	// 			1,
	// 			2,
	// 			3,
	// 			4
	// 		]
	// 	},
	// 	"name": "Tomas"
	// }
}

func ExampleMap_SetAttrJSONByte() {
	newstruct := &TestStruct{
		Name:     "Tomas",
		Age:      20,
		Messages: []interface{}{"Hello", "World"},
		Gender:   "Male",
	}
	testmap := NewMapByStruct(newstruct)
	fmt.Println(testmap.GetJSONString())

	testmap.SetAttrJSONByte(
		"misc",
		[]byte("{\"points\":[1,2,3,4],\"name\":\"cacao\"}"),
	)

	fmt.Println(testmap.GetJSONPretty())
}

func ExampleMap_SetAttrQuery() {
	strjson := "{\"misc\":{\"name\":\"cacao\"},\"name\":\"Tomas\"}"

	testmap := NewMapByJSONByte([]byte(strjson))

	testmap.SetAttrQuery("misc.name", "cocoa")
	fmt.Println(testmap.GetJSONString())

	testmap.SetAttrQuery("misc.name", 3)
	fmt.Println(testmap.GetJSONString())

	anonystruct := struct {
		AttrA string `json:"attr_a"`
		AttrB string `json:"attr_b"`
	}{
		AttrA: "TestA",
		AttrB: "TestB",
	}

	testmap.SetAttrQuery("misc.name", anonystruct)
	fmt.Println(testmap.GetJSONPretty())
	// Output:
	// {"misc":{"name":"cocoa"},"name":"Tomas"}
	// {"misc":{"name":3},"name":"Tomas"}
	// {
	// 	"misc": {
	// 		"name": {
	// 			"attr_a": "TestA",
	// 			"attr_b": "TestB"
	// 		}
	// 	},
	// 	"name": "Tomas"
	// }
}

func ExampleMap_IsExistAttr() {
	newstruct := &TestStruct{
		Name:     "Tomas",
		Age:      20,
		Messages: []interface{}{"Hello", "World"},
		Gender:   "Male",
	}
	testmap := NewMapByStruct(newstruct)
	fmt.Println(testmap.GetJSONString())
	fmt.Println(testmap.IsExistAttr("age"))
	// Output:
	// {"age":20,"gender":"Male","messages":["Hello","World"],"name":"Tomas"}
	// true
}

func ExampleMap_GetAttrMap() {
	strjson := "{\"misc\":{\"name\":\"cacao\"},\"name\":\"Tomas\"}"
	testmap := NewMapByJSONByte([]byte(strjson))

	miscmap := testmap.GetAttrMap("misc")
	fmt.Println(miscmap.GetJSONString())
	fmt.Println(miscmap.GetAttrString("name"))
	// Output:
	// {"name":"cacao"}
	// cacao
}

func ExampleMap_GetAttrQuery() {
	strjson := "{\"shop\":{\"name\":{\"first\":\"hollys\",\"last\":\"coffee\"},\"location\":[\"a\",\"b\"]}}"
	testmap := NewMapByJSONByte([]byte(strjson))
	fmt.Println(testmap.GetJSONPretty())

	fmt.Println(testmap.GetAttrQuery("shop.name.first"))
	fmt.Println(reflect.TypeOf(testmap.GetAttrQuery("shop.name.first")))

	fmt.Println(testmap.GetAttrQuery("shop.location"))
	fmt.Println(reflect.TypeOf(testmap.GetAttrQuery("shop.location")))

	// Output:
	// {
	// 	"shop": {
	// 		"location": [
	// 			"a",
	// 			"b"
	// 		],
	// 		"name": {
	// 			"first": "hollys",
	// 			"last": "coffee"
	// 		}
	// 	}
	// }

	// hollys
	// string

	// [a b]
	// []interface {}
}

func ExampleMap_GetMapInterface() {
	strjson := "{\"shop\":{\"name\":{\"first\":\"hollys\",\"last\":\"coffee\"},\"location\":[\"a\",\"b\"]}}"
	testmap := NewMapByJSONByte([]byte(strjson))

	fmt.Println(testmap.GetMapInterface())
	fmt.Println(reflect.TypeOf(testmap.GetMapInterface()))

	// Output:
	// map[shop:map[name:map[first:hollys last:coffee] location:[a b]]]
	// map[string]interface {}
}

func ExampleMap_GetAttrInt() {
	strjson := "{\"espresso\": 1.33222, \"americano\": \"1234\", \"latte\": \"2.234\" }"
	testmap := NewMapByJSONByte([]byte(strjson))

	fmt.Println(testmap.GetJSONPretty())

	fmt.Println(testmap.GetAttrInt("espresso"))
	fmt.Println(testmap.GetAttrInt("americano"))
	fmt.Println(testmap.GetAttrInt("latte"))

	// Output
	// {
	// 	"americano": "1234",
	// 	"espresso": 1.33222,
	// 	"latte": "2.234"
	// }

	// 1
	// 1234
	// 2
}

func ExampleMap_GetAttrString() {
	strjson := "{\"espresso\": 1.33222, \"americano\": \"1234\"}"
	testmap := NewMapByJSONByte([]byte(strjson))

	fmt.Println(testmap.GetJSONPretty())

	fmt.Println(testmap.GetAttrString("espresso"))
	fmt.Println(testmap.GetAttrString("americano"))

	// Output:
	// {
	// 	"americano": "1234",
	// 	"espresso": 1.33222
	// }
	// 1
	// 1234
}
