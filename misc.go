package gq

import (
	"unsafe"

	"github.com/gramework/runtimer"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack"
)

// Struct2MsgPackByte 는 Struct를 MsgPack 포멧으로 변환하여 []byte 로 리턴합니다
func Struct2MsgPackByte(s interface{}) []byte {
	t, e := msgpack.Marshal(s)
	if e != nil {
		logrus.Errorf("Struct2MsgPackByte Err : %s", e)
		return nil
	}
	return t
}

// Struct2JSONByte 는 Struct를 JSON 포멧으로 변환하여 []byte 로 리턴합니다
func Struct2JSONByte(s interface{}) []byte {
	t, e := ffjson.Marshal(s)
	if e != nil {
		logrus.Errorf("Struct2JSONByte Err : %s", e)
		return nil
	}
	return t
}

// Struct2JSONString 는 Struct를 JSON 포멧으로 변환하여 string 로 리턴합니다
func Struct2JSONString(s interface{}) string {
	return BytesToString(Struct2JSONByte(s))
}

// Map2Struct 는 map[string]interface{}를 struct 로 변환합니다
func Map2Struct(m map[string]interface{}, s interface{}) interface{} {
	t, e := ffjson.Marshal(m)
	if e != nil {
		return nil
	}
	e = ffjson.Unmarshal(t, s)
	if e != nil {
		return nil
	}
	return s
}

// JSONString2Struct 는 JSON 포멧 형태의 string을 struct 로 변환합니다
func JSONString2Struct(raw string, s interface{}) interface{} {
	return JSONByte2Struct(StringToBytes(raw), s)
}

// JSONByte2Struct 는 JSON 포멧 형태의 []byte를 struct 로 변환합니다
func JSONByte2Struct(raw []byte, s interface{}) interface{} {
	if e := ffjson.Unmarshal(raw, &s); e != nil {
		logrus.Errorf("JSONByte2Struct Err : %s", e)
		return nil
	}
	return s
}

// MsgPackByte2Struct 는 MsgPack 포멧 형태의 []byte를 struct 로 변환합니다
func MsgPackByte2Struct(raw []byte, s interface{}) interface{} {
	if e := msgpack.Unmarshal(raw, &s); e != nil {
		logrus.Errorf("MsgPackByte2Struct Err : %s", e)
		return nil
	}
	return s
}

// InterfaceSlice2StringSlice 는 Interface{} 타입의 슬라이스를 String 타입의 슬라이스로 변경합니다
func InterfaceSlice2StringSlice(tmp []interface{}) []string {
	res := []string{}
	for _, v := range tmp {
		switch v.(type) {
		case string:
			res = append(res, v.(string))
		case map[string]interface{}:
			v := Map(v.(map[string]interface{}))
			res = append(res, v.GetJSONString())
		}
	}
	return res
}

//
// from gramework
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToBytes(s string) []byte {
	strstruct := runtimer.StringStructOf(&s)
	return *(*[]byte)(unsafe.Pointer(&runtimer.SliceType2{
		Array: strstruct.Str,
		Len:   strstruct.Len,
		Cap:   strstruct.Len,
	}))
}
