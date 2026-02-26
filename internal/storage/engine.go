package storage

import (
	"sync"
	"time"
)

type Store struct {
	mu   sync.RWMutex
	data map[string]Item
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]Item),
	}
}

func (s *Store) Set(key string, value string, ttl time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var expiry time.Time
	if ttl > 0 {
		expiry = time.Now().Add(ttl)
	}

	s.data[key] = Item{
		Value:     value,
		ExpiresAt: expiry,
	}
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	item, exists := s.data[key]
	defer s.mu.RUnlock()

	if !exists {
		return "", false
	}

	if item.IsExpired() {
		go s.Delete(key)
		return "", false
	}

	return item.Value, true
}

func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}
