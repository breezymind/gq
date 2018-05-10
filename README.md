# GQ [![GoDoc](https://godoc.org/github.com/breezymind/gq?status.svg)](https://godoc.org/github.com/breezymind/gq)

> GQ 는,  
> string, []byte 등으로 표현된 json 텍스트 포멧의 데이터를 특정 struct 또는 동시성을 보장하는 map[string]interface{} 형태로 변환하여 동적으로 속성을 가공 가능 하도록 도와줍니다.

## Todos

- [ ] gq map, test example, godoc 작성
- [x] gq syncmap, test example, godoc 작성
- [x] misc, test example 작성
- [x] misc, godoc comment

## Installation

```bash
go get "github.com/breezymind/gq"
```

## Usage

### map 
: 동시성을 보장하지 않는 데이터를 가공 할때 사용합니다. 

[NewMapByJSONByte - 예제보기](https://godoc.org/github.com/breezymind/gq#example-NewMapByJSONByte)
> NewMapByJSONByte 은 JSON 포멧의 []byte 를 `Map 타입`으로 변환하여 인스턴스를 생성합니다.

[NewMapByStruct - 예제보기](https://godoc.org/github.com/breezymind/gq#example-NewMapByStruct)
> NewMapByStruct 은 struct 를 `Map 타입`으로 변환하여 인스턴스를 생성합니다.

[Map.GetAttrInt - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map.GetAttrInt)
> GetAttrInt 는 Map의 속성값을 integer 로 리턴 합니다.

[Map.GetAttrMap - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map.GetAttrMap)
> GetAttrMap 는 Map의 속성값을 Map타입으로 리턴합니다.

[Map.GetAttrQuery - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map.GetAttrQuery)
> GetAttrQuery 는 Map의 속성값을 dot(.)으로 접근하여 가져올 수 있게 합니다.

[Map.GetAttrString - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map.GetAttrString)
> GetAttrString 는 Map의 속성값을 string 으로 리턴 합니다.

[Map.GetJSONByte - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map.GetJSONByte)
> GetJSONByte 는 `Map 타입`에 정의된 데이터셋을 JSON 포멧(byte) 으로 리턴합니다

[Map.GetJSONPretty - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map.GetJSONPretty)
> GetJSONPretty 는 `Map 타입`에 정의된 데이터셋을 JSON string 으로 보기좋게 리턴합니다

[Map.GetJSONString - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map.GetJSONString)
> GetJSONString 는 `Map 타입`에 정의된 데이터셋을 JSON 포멧(string) 으로 리턴합니다

[Map.GetMapInterface - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map.GetMapInterface)
> GetMapInterface 는 `Map 타입`을 map[string]interface{} 로 리턴 합니다.

[Map.IsExistAttr - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map.IsExistAttr)
> IsExistAttr 는 Map의 속성에 특정 키가 있는지 확인합니다.

[Map.SetAttr - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map.SetAttr)
> SetAttr 는 `Map 타입` 데이터 셋에 새로운 속성값을 interface{} 타입으로 정의합니다.

[Map.SetAttrJSONByte - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map.SetAttrJSONByte)
> SetAttrJSONByte 는 `Map 타입` 데이터 셋에 새로운 속성값을 JSON 포멧형태의 값(byte)으로 정의 합니다.

[Map.SetAttrJSONString - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map.SetAttrJSONString)
> SetAttrJSONString 는 Map에 새로운 key/value를 정의하며, value 값은 JSON 값(string)으로 참조하여 정의 합니다.

[Map.SetAttrMap - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map.SetAttrMap)
> SetAttrMap 는 `Map 타입` 데이터 셋에 새로운 속성값을 `Map 타입`으로 정의합니다.

[Map.SetAttrQuery - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map.SetAttrQuery)
> SetAttrQuery 는 Map의 속성값을 dot(.)으로 접근하여 정의할 수 있습니다.

[Map.SetJSONByte - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map.SetJSONByte)
> SetJSONByte 는 JSON 포멧 형태의 []byte를 `Map 타입`으로 재정의합니다.

[Map.SetJSONString - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map.SetJSONString)
> SetJSONString 는 JSON 포멧 형태의 string을 `Map 타입`으로 재정의합니다.

[Map.SetStruct - 예제보기](https://godoc.org/github.com/breezymind/gq#example-Map.SetStruct)
> SetStruct 는 Struct를 `Map 타입`으로 재정의합니다.
---

### syncmap 
> GQ Map 에서 동시성을 보장하려면 Mutex 를 사용하는 syncmap 을 사용하여야 합니다.
>
> map 의 메소드와 동일하며 인스턴스 생성만 NewSyncMap 와 같은 syncmap 생성 메소드를 사용 하시면 됩니다.

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

## License
[MIT license](https://opensource.org/licenses/MIT)
