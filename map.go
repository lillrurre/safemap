package safemap

import "sync"

type SafeMap[T comparable, V any] struct {
	mut sync.RWMutex
	m   map[T]V
}

// NewMap create a new map with the given key-value types.
// Key must be a comparable value.
func NewMap[T comparable, V any]() *SafeMap[T, V] {
	return &SafeMap[T, V]{m: make(map[T]V)}
}

// Store safely stores a value for the given key.
// It does not take previous values into account.
func (s *SafeMap[T, V]) Store(key T, value V) {
	s.mut.Lock()
	defer s.mut.Unlock()
	s.m[key] = value
}

// Delete uses the builtin delete function to delete from the map, but safely.
func (s *SafeMap[T, V]) Delete(key T) {
	s.mut.Lock()
	defer s.mut.Unlock()
	delete(s.m, key)
}

// Load get the value from the map with the given key.
// This works like map[key] -> value, so zero or nil values may occur.
// Use LoadBool for additional safety.
func (s *SafeMap[T, V]) Load(key T) (value V) {
	s.mut.RLock()
	defer s.mut.RUnlock()
	return s.m[key]
}

// LoadBool works like: val, ok := map[key] but is thread safe.
func (s *SafeMap[T, V]) LoadBool(key T) (value V, ok bool) {
	s.mut.RLock()
	defer s.mut.RUnlock()
	value, ok = s.m[key]
	return value, ok
}

// Range ranges over the key values in the map and runs the provided function with
// the key and the value as parameters.
func (s *SafeMap[T, V]) Range(f func(T, V) bool) {
	s.mut.RLock()
	defer s.mut.RUnlock()
	for k, v := range s.m {
		if !f(k, v) {
			break
		}
	}
}

// RangeValue loops over the values in the map and runs the function.
// Works like Range, but it does only use the value as parameter.
func (s *SafeMap[T, V]) RangeValue(f func(V) bool) {
	s.mut.RLock()
	defer s.mut.RUnlock()
	for _, v := range s.m {
		if !f(v) {
			break
		}
	}
}

// Swap changes the value for a key. It returns the old value and a boolean indicating if
// the previous value was available in the map.
func (s *SafeMap[T, V]) Swap(key T, value V) (previous V, loaded bool) {
	previous, loaded = s.LoadBool(key)
	s.Store(key, value)
	return previous, loaded
}
