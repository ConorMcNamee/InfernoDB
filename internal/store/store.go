package store

import "sync"

type KVStore struct {
	lock  *sync.Mutex
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
	kv.store[key] = []byte(value)
	return true
}
