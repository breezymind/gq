package gq

import (
	"encoding/json"

	"github.com/sirupsen/logrus"
)

// Struct2JSONByte 는 Struct를 JSON 포멧으로 변환하여 []byte 로 리턴합니다
func Struct2JSONByte(s interface{}) []byte {
	t, e := json.Marshal(s)
	if e != nil {
		return nil
	}
	return t
}

// Struct2JSONString 는 Struct를 JSON 포멧으로 변환하여 string 로 리턴합니다
func Struct2JSONString(s interface{}) string {
	return string(Struct2JSONByte(s))
}

// Map2Struct 는 map[string]interface{}를 struct 로 변환합니다
func Map2Struct(m map[string]interface{}, s interface{}) interface{} {
	t, e := json.Marshal(m)
	if e != nil {
		return nil
	}
	e = json.Unmarshal(t, s)
	if e != nil {
		return nil
	}
	return s
}

// JSONString2Struct 는 JSON 포멧 형태의 string을 struct 로 변환합니다
func JSONString2Struct(raw string, s interface{}) interface{} {
	return JSONByte2Struct([]byte(raw), s)
}

// JSONByte2Struct 는 JSON 포멧 형태의 []byte를 struct 로 변환합니다
func JSONByte2Struct(raw []byte, s interface{}) interface{} {
	if e := json.Unmarshal(raw, &s); e != nil {
		logrus.Errorf("SetByte Err : %s", e)
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
