package main

import (
	"errors"
	"sync"
)

type Store interface {
	Put(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

var ErrNoSuchKey = errors.New("no such key")

type simpleStore struct {
	sync.RWMutex
	store map[string]string
}

func NewSimpleStore() Store {
	return &simpleStore{
		store: make(map[string]string),
	}
}

func (s *simpleStore) Put(key, value string) error {
	s.Lock()
	defer s.Unlock()

	s.store[key] = value
	return nil
}

func (s *simpleStore) Get(key string) (string, error) {
	s.RLock()
	value, ok := s.store[key]
	s.RUnlock()

	if !ok {
		return "", ErrNoSuchKey
	}

	return value, nil
}

func (s *simpleStore) Delete(key string) error {
	s.Lock()
	defer s.Unlock()

	delete(s.store, key)

	return nil
}
