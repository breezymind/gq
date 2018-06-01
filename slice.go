package gq

import (
	"encoding/json"
	"sort"

	"github.com/sirupsen/logrus"
)

// KV 타입은 Key-Value 데이터 형태의 struct 이며, K 속성은 Key(string), V속성은 Value(interface{}) 를 의미 합니다.
type KV struct {
	K string      `json:"k"`
	V interface{} `json:"v"`
}

// KVSlice 타입은 "KV" 타입의 Slice 타입이며, "Map" 타입과 상호 전환하여 사용할 수 있습니다.
type KVSlice []*KV

// Clone 는 `KVSlice 타입` 을 복제하여 리턴합니다.
func (t KVSlice) Clone() KVSlice {
	tmplist := make([]*KV, len(t))
	copy(tmplist, t)
	return KVSlice(tmplist)
}

// SortByValue 는 KVSlice 를 V값 기준으로 역정렬하여 리턴합니다.
func (t KVSlice) RSortByValue() KVSlice {
	sort.Slice(t, func(i, j int) bool {
		res := true
		switch x, y := t[i].V, t[j].V; x.(type) {
		case float64:
			res = x.(float64) > y.(float64)
		default:
			res = x.(int) > y.(int)
		}
		return res
	})
	return t
}

// GetJSONPretty 는 `KVSlice 타입`에 정의된 KV 데이터셋을 JSON string 으로 보기좋게 리턴합니다
func (t KVSlice) GetJSONPretty() string {
	res, e := json.MarshalIndent(t, "", "\t")
	if e != nil {
		logrus.Errorf("GetPretty Err : %s", e)
		return ""
	}
	return string(res)
}

// GetJSONString 는 `KVSlice 타입`에 정의된 KV 데이터셋을 JSON 포멧(string) 으로 리턴합니다
func (t KVSlice) GetJSONString() string {
	res := t.GetJSONByte()
	if res != nil {
		return string(res)
	}
	return ""
}

// GetJSONByte 는 `KVSlice 타입`에 정의된 KV 데이터셋을 JSON 포멧(byte) 으로 리턴합니다
func (t KVSlice) GetJSONByte() []byte {
	res, e := json.Marshal(t)
	if e != nil {
		logrus.Errorf("GetByte Err : %s", e)
		return nil
	}
	return res
}
