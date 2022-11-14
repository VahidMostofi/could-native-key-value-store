package main

import "errors"

type Store interface {
	Put(key string, value string) error
	Get(key string) (string, error)
	Delete(key string) error
}

var ErrNoSuchKey = errors.New("no such key")

type simpleStore struct {
	store map[string]string
}

func NewSimpleStore() Store {
	return &simpleStore{
		store: make(map[string]string),
	}
}

func (s *simpleStore) Put(key, value string) error {
	s.store[key] = value
	return nil
}

func (s *simpleStore) Get(key string) (string, error) {
	value, ok := s.store[key]
	if !ok {
		return "", ErrNoSuchKey
	}

	return value, nil
}

func (s *simpleStore) Delete(key string) error {
	delete(s.store, key)

	return nil
}
