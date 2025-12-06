package lookup

import (
	"iter"
	"maps"
)

// set is a Lookup that uses a map to look up values.
type set[T comparable] struct {
	// inner is the map that backs the storage.
	inner map[T]struct{}
}

// newSetLookup creates a Lookup that is always backed by a set.
func newSetLookupWithSize[T comparable](size int, elements ...T) *set[T] {
	inner := make(map[T]struct{}, size)
	for _, element := range elements {
		inner[element] = struct{}{}
	}
	return &set[T]{
		inner: inner,
	}
}

// Has implements Lookup.
func (s *set[T]) Has(value T) bool {
	_, ok := s.inner[value]
	return ok
}

// Add implements Lookup by adding another entry to the map. It returns true if the
// value already exists in the slice.
func (s *set[T]) Add(value T) bool {
	if s.Has(value) {
		return true
	}
	s.inner[value] = struct{}{}
	return false
}

// Remove implements Lookup. It removes a matching entry from the map.
func (s *set[T]) Remove(value T) bool {
	if !s.Has(value) {
		return false
	}
	delete(s.inner, value)
	return true
}

// Len implements Lookup. It returns the length of the map.
func (s *set[T]) Len() int {
	return len(s.inner)
}

// Iter implements Lookup. It returns an iterator over the keys in the map.
func (s *set[T]) Iter() iter.Seq[T] {
	return maps.Keys(s.inner)
}
