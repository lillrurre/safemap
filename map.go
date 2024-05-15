package safemap

import "sync"

type SafeMap[T comparable, V any] struct {
	mut sync.RWMutex
	m   map[T]V
}

func NewMap[T comparable, V any]() *SafeMap[T, V] {
	return &SafeMap[T, V]{m: make(map[T]V)}
}

func (s *SafeMap[T, V]) Store(key T, value V) {
	s.mut.Lock()
	defer s.mut.Unlock()
	s.m[key] = value
}

func (s *SafeMap[T, V]) Delete(key T) {
	s.mut.Lock()
	defer s.mut.Unlock()
	delete(s.m, key)
}

func (s *SafeMap[T, V]) Load(key T) (value V) {
	s.mut.RLock()
	defer s.mut.RUnlock()
	return s.m[key]
}

func (s *SafeMap[T, V]) LoadBool(key T) (value V, ok bool) {
	s.mut.RLock()
	defer s.mut.RUnlock()
	value, ok = s.m[key]
	return value, ok
}

func (s *SafeMap[T, V]) Range(f func(T, V) bool) {
	s.mut.RLock()
	defer s.mut.RUnlock()
	for k, v := range s.m {
		if !f(k, v) {
			break
		}
	}
}

func (s *SafeMap[T, V]) RangeValue(f func(V) bool) {
	s.mut.RLock()
	defer s.mut.RUnlock()
	for _, v := range s.m {
		if !f(v) {
			break
		}
	}
}

func (s *SafeMap[T, V]) Swap(key T, value V) (previous V, loaded bool) {
	s.mut.Lock()
	defer s.mut.Unlock()
	previous, loaded = s.m[key]
	s.m[key] = value
	return previous, loaded
}
