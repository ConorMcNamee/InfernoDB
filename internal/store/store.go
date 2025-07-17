package store

import (
	"database/internal/log"
	"sync"
)

type KVStore struct {
	lock  *sync.Mutex
	dbWAL *log.AppendOnlyLog
	store map[string][]byte
}

func NewKVStore() *KVStore {
	return &KVStore{
		store: make(map[string][]byte),
	}
}

func (kv *KVStore) Get(key string) ([]byte, bool) {
	return kv.store[key], true
}

func (kv *KVStore) Set(key string, value []byte) bool {
	err := kv.dbWAL.SaveToFile(key, value)
	if err != nil {
		return false
	}

	kv.store[key] = []byte(value)
	return true
}
