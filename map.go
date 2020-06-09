package hashmap

import (
	"hash/crc32"
	"log"
	"sync"
)

type HashMap struct {
	num  int
	list []hashMapCore
}

type hashMapCore struct {
	sync.RWMutex
	data map[string]interface{}
}

func Make(num int) *HashMap {
	h := &HashMap{
		num: num,
	}
	for i := 0; i < num; i++ {
		h.list = append(h.list, hashMapCore{
			RWMutex: sync.RWMutex{},
			data:    make(map[string]interface{}),
		})
	}
	return h
}

func (h *HashMap) Set(key string, val interface{}) {
	id := h.getHashId(key)
	log.Println("写入第", id, "个map")
	h.list[id].set(key, val)
}

func (h *HashMap) Get(key string) interface{} {
	id := h.getHashId(key)
	return h.list[id].get(key)
}

func (h *HashMap) getHashId(key string) int {
	return int(crc32.ChecksumIEEE([]byte(key))) % h.num
}

func (hc *hashMapCore) set(key string, val interface{}) {
	defer hc.Unlock()
	hc.Lock()

	hc.data[key] = val
}

func (hc *hashMapCore) get(key string) interface{} {
	defer hc.RUnlock()
	hc.RLock()

	val, ok := hc.data[key]
	if !ok {
		return nil
	}
	return val
}
