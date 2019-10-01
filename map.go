package gq

import (
	"bytes"
	"encoding/json"
	"github.com/gramework/runtimer"
	"github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// Map 타입은 GQ 패키지 전반에서 사용할 기본 타입이며, map[string]interface{} 의 alias 이기도 합니다.
// GQ 패키지에서는 이를 "Map"이라는 이름으로 선언하고, SyncMap struct 에서는 mutex 로 동시성을 보장합니다.
//
// (map[string]interface{} 형태의 map을 사용함으로서 다양한 데이터 구조로 상호 전환이 가능한 장점이 있습니다.)
type Map map[string]interface{}

// NewMap 은 `Map 타입`으로 기본 인스턴스를 생성합니다.
//
// NewMap 이름으로 시작하는 메소드는 동시성을 보장하지 않습니다.
// 동시성 보장이 요구되는 `Map 타입`은 NewSyncMap 으로 시작하는 메소드를 사용하세요.
func NewMap() *Map {
	return &Map{}
}

// NewMapByMapType 은 map[string]interface{}타입 (map[string]string, map[string]int 등)의 데이터를 `Map 타입`으로 변환하여 인스턴스를 생성합니다.
func NewMapByMapTypes(raw interface{}) *Map {
	res := make(Map)
	switch raw.(type) {
	case map[string]interface{}:
		res = Map(raw.(map[string]interface{}))
		return &res
	case map[string]string:
		for k, v := range raw.(map[string]string) {
			res.SetAttr(k, v)
		}
	case map[string]int:
		for k, v := range raw.(map[string]int) {
			res.SetAttr(k, v)
		}
	case map[string]bool:
		for k, v := range raw.(map[string]bool) {
			res.SetAttr(k, v)
		}
	}
	return &res
}

// NewMapByJSONByte 은 JSON 포멧의 []byte 를 `Map 타입`으로 변환하여 인스턴스를 생성합니다.
func NewMapByJSONByte(raw []byte) *Map {
	return (&Map{}).SetJSONByte(raw)
}

// NewMapByMsgPackByte 은 JSON 포멧의 []byte 를 `Map 타입`으로 변환하여 인스턴스를 생성합니다.
func NewMapByMsgPackByte(raw []byte) *Map {
	return (&Map{}).SetMsgPackByte(raw)
}

// NewMapByStruct 은 struct 를 `Map 타입`으로 변환하여 인스턴스를 생성합니다.
func NewMapByStruct(raw interface{}) *Map {
	tmp, e := json.Marshal(raw)
	if e != nil {
		logrus.Errorf("NewMapByStruct Err : %s", e)
		return nil
	}
	t := &Map{}
	return t.SetJSONByte(tmp)
}

// GetJSONPretty 는 `Map 타입`에 정의된 데이터셋을 JSON string 으로 보기좋게 리턴합니다
func (t *Map) GetJSONPretty() string {
	res := &bytes.Buffer{}
	enc := json.NewEncoder(res)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "\t")
	if e := enc.Encode(t); e != nil {
		logrus.Errorf("GetJSONByte Err : %s\n%s", e)
		return ""
	}
	return res.String()
}

// GetJSONString 는 `Map 타입`에 정의된 데이터셋을 JSON 포멧(string) 으로 리턴합니다
func (t *Map) GetJSONString() string {
	res := t.GetJSONByte()
	if res != nil {
		// return string(res)
		return BytesToString(res)
	}
	return ""
}

// GetJSONByte 는 `Map 타입`에 정의된 데이터셋을 JSON 포멧(byte) 으로 리턴합니다
func (t *Map) GetJSONByte() []byte {
	res := &bytes.Buffer{}
	enc := json.NewEncoder(res)
	enc.SetEscapeHTML(false)
	if e := enc.Encode(t); e != nil {
		logrus.Errorf("GetJSONByte Err : %s\n%s", e)
		return nil
	}
	return res.Bytes()
}

// GetMsgPackString 는 `Map 타입`에 정의된 데이터셋을 MsgPack 포멧(string) 으로 리턴합니다
func (t *Map) GetMsgPackString() string {
	res := t.GetMsgPackByte()
	if res != nil {
		// return string(res)
		return BytesToString(res)
	}
	return ""
}

// GetMsgPackByte 는 `Map 타입`에 정의된 데이터셋을 MsgPack 포멧(byte) 으로 리턴합니다
func (t *Map) GetMsgPackByte() []byte {
	b, e := msgpack.Marshal(t)
	if e != nil {
		return nil
	}
	return b
}

// SetJSONString 는 JSON 포멧 형태의 string을 `Map 타입`으로 재정의합니다.
func (t *Map) SetJSONString(raw string) *Map {
	return t.SetJSONByte(StringToBytes(raw))
}

// SetJSONByte 는 JSON 포멧 형태의 []byte를 `Map 타입`으로 재정의합니다.
func (t *Map) SetJSONByte(raw []byte) *Map {
	if e := json.Unmarshal(raw, t); e != nil {
		// logrus.Errorf("SetJSONByte Err : %s\n%s\n", e, string(raw))
		logrus.Errorf("SetJSONByte Err : %s\n%s\n", e, BytesToString(raw))
		return nil
	}
	return t
}

// SetMsgPackByte 는 MsgPack 포멧 형태의 []byte를 `Map 타입`으로 재정의합니다.
func (t *Map) SetMsgPackByte(raw []byte) *Map {
	if e := msgpack.Unmarshal(raw, t); e != nil {
		// logrus.Errorf("SetJSONByte Err : %s\n%s\n", e, string(raw))
		logrus.Errorf("SetMsgPackByte Err : %s\n%s\n", e, BytesToString(raw))
		return nil
	}
	return t
}

// SetStruct 는 Struct를 `Map 타입`으로 재정의합니다.
func (t *Map) SetStruct(raw interface{}) *Map {
	res, e := json.Marshal(raw)
	if e != nil {
		logrus.Errorf("SetStruct Err : %s", e)
		return nil
	}
	return t.SetJSONByte(res)
}

// func (t *Map) SetInterface(raw interface{}) *Map {
// 	tmp := Map(raw.(map[string]interface{}))
// 	t = &tmp
// 	return t
// }

// SetAttr 는 `Map 타입` 데이터 셋에 새로운 속성값을 interface{} 타입으로 정의합니다.
func (t *Map) SetAttr(k string, v interface{}) *Map {
	(*t)[k] = v
	return t
}

// SetAttrMap 는 `Map 타입` 데이터 셋에 새로운 속성값을 `Map 타입`으로 정의합니다.
func (t *Map) SetAttrMap(k string, v *Map) *Map {
	(*t)[k] = v
	return t
}

// SetAttrJSONString 는 Map에 새로운 key/value를 정의하며, value 값은 JSON 값(string)으로 참조하여 정의 합니다.
func (t *Map) SetAttrJSONString(k, v string) *Map {
	return t.SetAttrJSONByte(k, StringToBytes(v))
}

// SetAttrJSONByte 는 `Map 타입` 데이터 셋에 새로운 속성값을 JSON 포멧형태의 값(byte)으로 정의 합니다.
func (t *Map) SetAttrJSONByte(k string, v []byte) *Map {
	attr := make((map[string]interface{}))
	if e := json.Unmarshal(v, &attr); e != nil {
		logrus.Errorf("SetByte Err : %s", e)
		return nil
	}
	(*t)[k] = attr
	return t
}

// SetAttrQuery 는 Map의 속성값을 dot(.)으로 접근하여 정의할 수 있습니다.
func (t *Map) SetAttrQuery(k string, v interface{}) *Map {
	q := strings.Split(k, ".")
	max := len(q)

	if max < 2 {
		return t.SetAttr(k, v)
	}

	var mapctx *Map
	for idx, attr := range q {
		switch true {
		case idx == (max - 1):
			return mapctx.SetAttr(attr, v)
		default:
			var tmp *Map
			if idx < 1 {
				tmp = t.GetAttrMap(attr)
			} else {
				tmp = mapctx.GetAttrMap(attr)
			}
			if tmp == nil {
				break
			}
			mapctx = tmp
		}
	}
	return nil
}

// IsExistAttr 는 Map의 속성에 특정 키가 있는지 확인합니다.
func (t *Map) IsExistAttr(k string) bool {
	if t.GetAttrInterface(k) != nil {
		return true
	}
	return false
}

// GetAttrMap 는 Map의 속성값을 Map타입으로 리턴합니다.
func (t *Map) GetAttrMap(k string) *Map {
	if v, ok := (*t)[k]; ok {
		switch v.(type) {
		case *Map:
			return v.(*Map)
		default:
			res := Map(v.(map[string]interface{}))
			return &res
		}
	}
	return nil
}

// GetAttrQuery 는 Map의 속성값을 dot(.)으로 접근하여 가져올 수 있게 합니다.
func (t *Map) GetAttrQuery(k string) interface{} {
	q := strings.Split(k, ".")
	max := len(q)

	if max < 2 {
		return t.GetAttrInterface(k)
	}

	var mapctx *Map
	for idx, attr := range q {
		switch true {
		case idx == (max - 1):
			return mapctx.GetAttrInterface(attr)
		default:
			var tmp *Map
			if idx < 1 {
				tmp = t.GetAttrMap(attr)
			} else {
				tmp = mapctx.GetAttrMap(attr)
			}
			if tmp == nil {
				break
			}
			mapctx = tmp
		}
	}
	return nil
}

// GetAttrInterface 는 Map의 속성값을 interface{} 로 리턴 합니다.
func (t *Map) GetAttrInterface(k string) interface{} {
	if v, ok := (*t)[k]; ok {
		return v
	}
	return nil
}

// GetMapInterface 는 `Map 타입`을 map[string]interface{} 로 리턴 합니다.
func (t *Map) GetMapInterface() map[string]interface{} {
	return *t
}

// GetAttrInt 는 Map의 속성값을 integer 로 리턴 합니다.
func (t *Map) GetAttrInt(k string) int {
	v := t.GetAttrInterface(k)
	switch v.(type) {
	case string:
		v, _ := runtimer.Atoi(v.(string))
		return v
	case float64:
		return int(v.(float64))
	case uint64:
		return int(v.(uint64))
	case int64:
		return int(v.(int64))
	case uint32:
		return int(v.(uint32))
	case int32:
		return int(v.(int32))
	case uint16:
		return int(v.(uint16))
	case int16:
		return int(v.(int16))
	case uint8:
		return int(v.(uint8))
	case int8:
		return int(v.(int8))
	default:
		return v.(int)
	}
}

// GetAttrString 는 Map의 속성값을 string 으로 리턴 합니다.
func (t *Map) GetAttrString(k string) string {
	v := t.GetAttrInterface(k)
	if v == nil {
		return ""
	}
	switch v.(type) {
	case float64:
		return strconv.Itoa(int(v.(float64)))
	case int:
		return strconv.Itoa(v.(int))
	default:
		return v.(string)
	}
}

// GetAttrSlice 는 Map의 속성값을 []interface{} 으로 리턴 합니다.
func (t *Map) GetAttrSlice(k string) []interface{} {
	v := t.GetAttrInterface(k)
	if v == nil {
		return []interface{}{}
	}
	switch v.(type) {
	case []interface{}:
		return v.([]interface{})
	default:
		return []interface{}{}
	}
}

// GetAttrStringSlice 는 Map의 속성값을 []string 으로 리턴 합니다.
func (t *Map) GetAttrStringSlice(k string) []string {
	return InterfaceSlice2StringSlice(t.GetAttrSlice(k))
}

// DelAttr 는 Map의 속성값을 삭제 합니다.
func (t *Map) DelAttr(k string) *Map {
	delete((*t), k)
	return t
}

// Keys 는 Map의 key 값을 []string 으로 리턴합니다.
func (t *Map) Keys() []string {
	keys := reflect.ValueOf(t).Elem().MapKeys()
	res := make([]string, len(keys))
	for k, v := range keys {
		res[k] = v.String()
	}
	sort.Strings(res)
	return res
}

// Values 는 Map의 value 값을 []interface{} 으로 리턴합니다.
func (t *Map) Values() []interface{} {
	tmp := make([]interface{}, 0)
	for _, key := range t.Keys() {
		tmp = append(tmp, t.GetAttrInterface(key))
	}
	return tmp
}

// Clone 는 `Map 타입` 을 복제하여 리턴합니다.
func (t *Map) Clone() *Map {
	src := map[string]interface{}(*t)
	srcval := reflect.ValueOf(src)
	dst := make(map[string]interface{}, len(src))
	dstval := reflect.ValueOf(dst)
	for _, k := range srcval.MapKeys() {
		dstval.SetMapIndex(k, srcval.MapIndex(k))
	}
	l := Map(dst)
	return &l
}

// ToKeyValueSlice 는 `Map 타입` 을 `KVSlice 타입(Key-Value 데이터 구조의 Slice)`으로 변경하여 리턴합니다.
func (t *Map) ToKeyValueSlice() KVSlice {
	res := []*KV{}
	val := t.Values()
	for idx, key := range t.Keys() {
		res = append(res, &KV{key, val[idx]})
	}
	return res
}
