package gq

import (
	"sync"
)

type SyncMap struct {
	sync.RWMutex
	data *Map
}

func NewSyncMap() *SyncMap {
	return &SyncMap{data: &Map{}}
}

func NewSyncMapByJSONByte(raw []byte) *SyncMap {
	if res := NewMapByJSONByte(raw); res != nil {
		return &SyncMap{data: res}
	}
	return nil
}

func NewSyncMapByStruct(raw interface{}) *SyncMap {
	if res := NewMapByStruct(raw); res != nil {
		return &SyncMap{data: NewMapByStruct(raw)}
	}
	return nil
}

// func NewSyncMapByInterface(raw interface{}) *SyncMap {
// 	return &SyncMap{data: NewMapByInterface(raw)}
// }

func (t *SyncMap) GetJSONPretty() string {
	t.RLock()
	defer t.RUnlock()
	return t.data.GetJSONPretty()
}

func (t *SyncMap) GetJSONString() string {
	t.RLock()
	defer t.RUnlock()
	return t.data.GetJSONString()
}

func (t *SyncMap) GetJSONByte() []byte {
	t.RLock()
	defer t.RUnlock()
	return t.data.GetJSONByte()
}

func (t *SyncMap) SetJSONString(raw string) *SyncMap {
	t.Lock()
	t.data.SetJSONString(raw)
	t.Unlock()
	return t
}

func (t *SyncMap) SetJSONByte(raw []byte) *SyncMap {
	t.Lock()
	t.data.SetJSONByte(raw)
	t.Unlock()
	return t
}

func (t *SyncMap) SetStruct(raw interface{}) *SyncMap {
	t.Lock()
	t.data.SetStruct(raw)
	t.Unlock()
	return t
}

// func (t *SyncMap) SetInterface(raw interface{}) *SyncMap {
// 	t.Lock()
// 	t.data.SetInterface(raw)
// 	t.Unlock()
// 	return t
// }

func (t *SyncMap) SetAttr(k string, v interface{}) *SyncMap {
	t.Lock()
	t.data.SetAttr(k, v)
	t.Unlock()
	return t
}

func (t *SyncMap) SetAttrJSONString(k, v string) *SyncMap {
	t.Lock()
	t.data.SetAttrJSONByte(k, []byte(v))
	t.Unlock()
	return t
}

func (t *SyncMap) SetAttrJSONByte(k string, v []byte) *SyncMap {
	t.Lock()
	t.data.SetAttrJSONByte(k, v)
	t.Unlock()
	return t
}

func (t *SyncMap) IsExistAttr(k string) bool {
	t.RLock()
	t.RUnlock()
	return t.data.IsExistAttr(k)
}

func (t *SyncMap) GetAttrMap(k string) *Map {
	t.RLock()
	defer t.RUnlock()
	return t.data.GetAttrMap(k)
}

func (t *SyncMap) SetAttrQuery(k string, v interface{}) *SyncMap {
	t.Lock()
	t.data.SetAttrQuery(k, v)
	t.Unlock()
	return t
}

func (t *SyncMap) GetAttrQuery(k string) interface{} {
	t.RLock()
	defer t.RUnlock()
	return t.data.GetAttrQuery(k)
}

func (t *SyncMap) GetAttrInterface(k string) interface{} {
	t.RLock()
	defer t.RUnlock()
	return t.data.GetAttrInterface(k)
}

func (t *SyncMap) GetMapInterface() map[string]interface{} {
	t.RLock()
	defer t.RUnlock()
	return t.data.GetMapInterface()
}

func (t *SyncMap) GetAttrInt(k string) int {
	t.RLock()
	defer t.RUnlock()
	return t.data.GetAttrInt(k)
}

func (t *SyncMap) GetAttrString(k string) string {
	t.RLock()
	defer t.RUnlock()
	return t.data.GetAttrString(k)
}

func (t *SyncMap) GetAttrSlice(k string) []interface{} {
	t.RLock()
	defer t.RUnlock()
	return t.data.GetAttrSlice(k)
}

func (t *SyncMap) GetAttrStringSlice(k string) []string {
	t.RLock()
	defer t.RUnlock()
	return t.data.GetAttrStringSlice(k)
}

func (t *SyncMap) DelAttr(k string) *SyncMap {
	t.Lock()
	t.data.DelAttr(k)
	t.Unlock()
	return t
}

func (t *SyncMap) Keys() []string {
	t.RLock()
	defer t.RUnlock()
	return t.data.Keys()
}

func (t *SyncMap) Values() []interface{} {
	t.RLock()
	defer t.RUnlock()
	return t.data.Values()
}

func (t *SyncMap) Clone() *SyncMap {
	t.RLock()
	defer t.RUnlock()
	return &SyncMap{data: t.data.Clone()}
}
