package main

import (
	"sync"
)

type OwnSyncMap struct {
	sync.RWMutex
	cash map[int]interface{}
}

func InitializeOwnSyncMap() *OwnSyncMap {
	return &OwnSyncMap{cash: make(map[int]interface{})}
}

func (m *OwnSyncMap) Set(key int, value interface{}) {
	m.Lock()
	defer m.Unlock()
	m.cash[key] = value
}

func (m *OwnSyncMap) Get(key int) (interface{}, bool) {
	m.RLock()
	defer m.RUnlock()
	item, isHas := m.cash[key]
	if isHas {
		return item, true
	}
	return nil, false
}

func (m *OwnSyncMap) Delete(key int) {
	m.Lock()
	defer m.Unlock()
	if _, isHas := m.cash[key]; isHas {
		delete(m.cash, key)
	}
}

func main() {
	m := InitializeOwnSyncMap()
	for i := 0; i < 1000; i++ {
		go func(m *OwnSyncMap, i int) {
			m.Set(i, i)
		}(m, i)
	}
}
