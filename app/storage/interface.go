package storage

import "time"

type Storage interface {
	Get(key string) (string, bool)
	Set(key, value string)
	SetWithExpiry(key, value string, expiry time.Duration)
}
