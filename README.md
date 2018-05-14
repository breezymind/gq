# GQ [![GoDoc](https://godoc.org/github.com/breezymind/gq?status.svg)](https://godoc.org/github.com/breezymind/gq)

[ENG TRANSLATION](https://github.com/hero0926/gq/edit/master/README_ENG.md)


> GQ 는,  
> 
> JSON 포멧(string, []byte)의 데이터를 golang 의 특정 struct 또는 동시성을 보장하는 map[string]interface{} 형태로 변환하여 동적으로 속성을 가공 가능하도록 도와줍니다.

## Installation

```bash
go get "github.com/breezymind/gq"
```

## Usage
* GetAttrString 예시
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

* SetAttrJSONString 예시
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
> godoc 에는 더 다양한 예시가 있습니다. 아래 `예제보기` 를 통해서 보셔도 됩니다.

### map 
> `Map 타입은 GQ 패키지 전반에서 사용할 기본 타입이며, map[string]interface{} 의 alias` 이기도 합니다.
> 
> `NewMap 이름으로 시작하는 메소드는 동시성을 보장하지 않습니다.   
> 동시성 보장이 요구되는 Map 타입은 NewSyncMap 으로 시작하는 메소드를 사용`하세요.

[NewMapByJSONByte - 예제보기](https://godoc.org/github.com/breezymind/gq#example-NewMapByJSONByte)
> NewMapByJSONByte 은 JSON 포멧의 []byte 를 `Map 타입`으로 변환하여 인스턴스를 생성합니다.

[NewMapByStruct - 예제보기](https://godoc.org/github.com/breezymind/gq#example-NewMapByStruct)
> NewMapByStruct 은 struct 를 `Map 타입`으로 변환하여 인스턴스를 생성합니다.

[Map.GetAttrInt - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map-GetAttrInt)
> GetAttrInt 는 Map의 속성값을 integer 로 리턴 합니다.

[Map.GetAttrMap - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map-GetAttrMap)
> GetAttrMap 는 Map의 속성값을 Map타입으로 리턴합니다.

[Map.GetAttrQuery - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map-GetAttrQuery)
> GetAttrQuery 는 Map의 속성값을 dot(.)으로 접근하여 가져올 수 있게 합니다.

[Map.GetAttrString - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map-GetAttrString)
> GetAttrString 는 Map의 속성값을 string 으로 리턴 합니다.

[Map.GetJSONByte - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map-GetJSONByte)
> GetJSONByte 는 `Map 타입`에 정의된 데이터셋을 JSON 포멧(byte) 으로 리턴합니다

[Map.GetJSONPretty - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map-GetJSONPretty)
> GetJSONPretty 는 `Map 타입`에 정의된 데이터셋을 JSON string 으로 보기좋게 리턴합니다

[Map.GetJSONString - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map-GetJSONString)
> GetJSONString 는 `Map 타입`에 정의된 데이터셋을 JSON 포멧(string) 으로 리턴합니다

[Map.GetMapInterface - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map-GetMapInterface)
> GetMapInterface 는 `Map 타입`을 map[string]interface{} 로 리턴 합니다.

[Map.IsExistAttr - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map-IsExistAttr)
> IsExistAttr 는 Map의 속성에 특정 키가 있는지 확인합니다.

[Map.SetAttr - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map-SetAttr)
> SetAttr 는 `Map 타입` 데이터 셋에 새로운 속성값을 interface{} 타입으로 정의합니다.

[Map.SetAttrJSONByte - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map-SetAttrJSONByte)
> SetAttrJSONByte 는 `Map 타입` 데이터 셋에 새로운 속성값을 JSON 포멧형태의 값(byte)으로 정의 합니다.

[Map.SetAttrJSONString - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map-SetAttrJSONString)
> SetAttrJSONString 는 Map에 새로운 key/value를 정의하며, value 값은 JSON 값(string)으로 참조하여 정의 합니다.

[Map.SetAttrMap - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map-SetAttrMap)
> SetAttrMap 는 `Map 타입` 데이터 셋에 새로운 속성값을 `Map 타입`으로 정의합니다.

[Map.SetAttrQuery - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map-SetAttrQuery)
> SetAttrQuery 는 Map의 속성값을 dot(.)으로 접근하여 정의할 수 있습니다.

[Map.SetJSONByte - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map-SetJSONByte)
> SetJSONByte 는 JSON 포멧 형태의 []byte를 `Map 타입`으로 재정의합니다.

[Map.SetJSONString - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map-SetJSONString)
> SetJSONString 는 JSON 포멧 형태의 string을 `Map 타입`으로 재정의합니다.

[Map.SetStruct - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map-SetStruct)
> SetStruct 는 Struct를 `Map 타입`으로 재정의합니다.
---

### syncmap 
> `SyncMap 타입은 "Map" 타입과 동일한 메소드들을 제공하며, mutex 로 동시성을 보장`합니다.

---

### misc
[Map2Struct - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map2Struct)
> Map2Struct 는 map[string]interface{}를 struct 로 변환합니다

[InterfaceSlice2StringSlice - 예제보기](https://godoc.org/github.com/breezymind/gq#example-InterfaceSlice2StringSlice)
> InterfaceSlice2StringSlice 는 Interface{} 타입의 슬라이스를 String 타입의 슬라이스로 변경합니다

[JSONByte2Struct - 예제보기](https://godoc.org/github.com/breezymind/gq#example-JSONByte2Struct)
> JSONByte2Struct 는 JSON 포멧 형태의 []byte를 struct 로 변환합니다

[JSONString2Struct - 예제보기](https://godoc.org/github.com/breezymind/gq#example-JSONString2Struct)
> JSONString2Struct 는 JSON 포멧 형태의 string을 struct 로 변환합니다

[Struct2JSONByte - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Struct2JSONByte)
> Struct2JSONString 는 Struct를 JSON 포멧으로 변환하여 string 로 리턴합니다

[Struct2JSONString - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Struct2JSONString)
> Struct2JSONByte 는 Struct를 JSON 포멧으로 변환하여 []byte 로 리턴합니다

## Todos

- [ ] gq map, test example, godoc 작성
- [x] gq syncmap, test example, godoc 작성
- [x] misc, test example 작성
- [x] misc, godoc comment

## License
[MIT license](https://opensource.org/licenses/MIT)