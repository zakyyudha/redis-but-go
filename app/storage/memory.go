package storage

import "time"

type MemoryStorage struct {
	data map[string]ValueWithExpiry
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]ValueWithExpiry),
	}
}

func (kv *MemoryStorage) Get(key string) (string, bool) {
	valueWithExpiry, ok := kv.data[key]
	if !ok {
		return "", false
	}

	if valueWithExpiry.IsExpired() {
		delete(kv.data, key)
		return "", false
	}

	return valueWithExpiry.value, true
}

func (kv *MemoryStorage) Set(key, value string) {
	kv.data[key] = ValueWithExpiry{
		value: value,
	}
}

func (kv *MemoryStorage) SetWithExpiry(key, value string, expiry time.Duration) {
	kv.data[key] = ValueWithExpiry{
		value:     value,
		expiresAt: time.Now().Add(expiry),
	}
}

type ValueWithExpiry struct {
	value     string
	expiresAt time.Time
}

func (v ValueWithExpiry) IsExpired() bool {
	if v.expiresAt.IsZero() {
		return false
	}

	return v.expiresAt.Before(time.Now())
}
