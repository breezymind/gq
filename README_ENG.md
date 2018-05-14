# GQ [![GoDoc](https://godoc.org/github.com/breezymind/gq?status.svg)](https://godoc.org/github.com/breezymind/gq)

> GQ Converts,
> 
> JSON format(string, []byte) data to specific golang struct or map[string]interface{} which supports concurrency so it helps you dynamically process properties.

## Installation

```bash
go get "github.com/breezymind/gq"
```

## Usage
* GetAttrString Example
```go
strjson := "{\"espresso\": 1.33222, \"americano\": \"1234\"}"
testmap := NewMapByJSONByte([]byte(strjson))

fmt.Println(testmap.GetJSONPretty())
// {
// 	"americano": "1234",
// 	"espresso": 1.33222
// }

fmt.Println(testmap.GetAttrString("espresso"))
// 1
fmt.Println(testmap.GetAttrString("americano"))
// 1234
```

* SetAttrJSONString Example
```go
newstruct := &TestStruct{
    Name:     "Tomas",
    Age:      20,
    Messages: []interface{}{"Hello", "World"},
    Gender:   "Male",
}
testmap := NewMapByStruct(newstruct)

fmt.Println(testmap.GetJSONString())

// {"age":20,"gender":"Male","messages":["Hello","World"],"name":"Tomas"}

testmap.SetAttrJSONString(
    "misc",
    "{\"points\":[1,2,3,4],\"name\":\"cacao\"}",
)
fmt.Println(testmap.GetJSONPretty())

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
```
> godoc has various example of usage. You can check it by godoc gq website.

### map 
> `Map is basic Type of GQ  package. it is alias of the map[string]interface{}.
> 
> `The method name that starts with NewMap - dosen't guarantee concurrency.   
>  For Map types that require concurrency guarantees, use methods starting with NewSyncMap.

[NewMapByJSONByte - Examples](https://godoc.org/github.com/breezymind/gq#example-NewMapByJSONByte)
> NewMapByJSONByte makes instance by converting JSON format's []byte to `Map` Type.

[NewMapByStruct - Examples](https://godoc.org/github.com/breezymind/gq#example-NewMapByStruct)
> NewMapByStruct makes instance by converting struct to `Map` Type.

[Map.GetAttrInt - Examples](https://godoc.org/github.com/breezymind/gq#example-Map-GetAttrInt)
> GetAttrInt returns Map's attributes to integer.

[Map.GetAttrMap - Examples](https://godoc.org/github.com/breezymind/gq#example-Map-GetAttrMap)
> GetAttrMap returns  Map's attributes to  Map.

[Map.GetAttrQuery - Examples](https://godoc.org/github.com/breezymind/gq#example-Map-GetAttrQuery)
> GetAttrQuery allows to get Map's attribute by using dot(.).

[Map.GetAttrString - Examples](https://godoc.org/github.com/breezymind/gq#example-Map-GetAttrString)
> GetAttrString returns Map's attributes to string다.

[Map.GetJSONByte - Examples](https://godoc.org/github.com/breezymind/gq#example-Map-GetJSONByte)
> GetJSONByte returns `Map Types` defined datasets to JSON format(byte).

[Map.GetJSONPretty - Examples](https://godoc.org/github.com/breezymind/gq#example-Map-GetJSONPretty)
> GetJSONPretty returns `Map Types` defined datasets to Json string, but prettier.

[Map.GetJSONString - Examples](https://godoc.org/github.com/breezymind/gq#example-Map-GetJSONString)
> GetJSONString returns `Map Types` defined datasets to JSON format(string).

[Map.GetMapInterface - Examples](https://godoc.org/github.com/breezymind/gq#example-Map-GetMapInterface)
> GetMapInterface returnt `Map Type` to  map[string]interface{}.

[Map.IsExistAttr - Examples](https://godoc.org/github.com/breezymind/gq#example-Map-IsExistAttr)
> IsExistAttr check the Map's attributes has a specific key.

[Map.SetAttr - Examples](https://godoc.org/github.com/breezymind/gq#example-Map-SetAttr)
> SetAttr defines a new attribute value in the `Map type` dataset as the interface{}.

[Map.SetAttrJSONByte - Examples](https://godoc.org/github.com/breezymind/gq#example-Map-SetAttrJSONByte)
> SetAttrJSONByte defines a value (byte) in JSON format.

[Map.SetAttrJSONString - Examples](https://godoc.org/github.com/breezymind/gq#example-Map-SetAttrJSONString)
> SetAttrJSONString defines a new key / value in the Map, and the value is defined by referring to the JSON value (string).

[Map.SetAttrMap - Examples](https://godoc.org/github.com/breezymind/gq#example-Map-SetAttrMap)
> SetAttrMap defines a new attribute value in the `Map type` dataset as `Map type`.

[Map.SetAttrQuery - Examples](https://godoc.org/github.com/breezymind/gq#example-Map-SetAttrQuery)
> SetAttrQuery can be defined by accessing the attribute values ​​of the Map with dot (.).

[Map.SetJSONByte - Examples](https://godoc.org/github.com/breezymind/gq#example-Map-SetJSONByte)
> SetJSONByte overrides the [] byte in JSON format to `Map type`.

[Map.SetJSONString - Examples](https://godoc.org/github.com/breezymind/gq#example-Map-SetJSONString)
> SetJSONString overrides a JSON format string as `Map type`.

[Map.SetStruct - Examples](https://godoc.org/github.com/breezymind/gq#example-Map-SetStruct)
> SetStruct overrides a Struct as `Map type`.
---

### syncmap 
> `SyncMap` type provides the same methods as 'Map' type, and guarantees concurrency with mutex.

---

### misc
[Map2Struct - Examples](https://godoc.org/github.com/breezymind/gq#example-Map2Struct)
> Map2Struct converts map [string] interface {} to a struct

[InterfaceSlice2StringSlice - Examples](https://godoc.org/github.com/breezymind/gq#example-InterfaceSlice2StringSlice)
> InterfaceSlice2StringSlice changes a Slice of Interface {} type to a String type slice

[JSONByte2Struct - Examples](https://godoc.org/github.com/breezymind/gq#example-JSONByte2Struct)
> JSONByte2Struct converts a [] byte in JSON format to a struct

[JSONString2Struct - Examples](https://godoc.org/github.com/breezymind/gq#example-JSONString2Struct)
> JSONString2Struct converts a string in JSON format to a struct

[Struct2JSONByte - Examples](https://godoc.org/github.com/breezymind/gq#example-Struct2JSONByte)
> Struct2JSONString converts the Struct to JSON format and returns it as a string

[Struct2JSONString - Examples](https://godoc.org/github.com/breezymind/gq#example-Struct2JSONString)
> Struct2JSONByte converts the Struct into JSON format and returns it as [] byte

## Todos

- [ ] gq map, test example, godoc
- [x] gq syncmap, test example, godoc
- [x] misc, test example
- [x] misc, godoc comment

## License
[MIT license](https://opensource.org/licenses/MIT)
