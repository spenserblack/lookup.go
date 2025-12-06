package lookup

import (
	"iter"
	"slices"
)

// slice is a Lookup that uses a slice to look up values.
type slice[T comparable] struct {
	// inner is the slice that backs the storage.
	inner []T
}

// newSliceLookup creates a Lookup that is always backed by a slice.
//
// The slice will be a slice of unique elements. This function will attempt to make the
// slice of elements unique.
func newSliceLookup[T comparable](elements []T) *slice[T] {
	unique := slices.Compact(elements)
	return &slice[T]{
		inner: unique,
	}
}

// newSliceLookupWithSize creates a Lookup that is always backed by a slice with the
// size provided.
//
// The slice will be a slice of unique elements. This function will attempt to make the
// slice of elements unique.
func newSliceLookupWithSize[T comparable](size int, elements ...T) *slice[T] {
	inner := make([]T, 0, size)
	unique := slices.Compact(elements)
	inner = append(inner, unique...)
	return &slice[T]{
		inner: inner,
	}
}

// Has implements Lookup. It uses slices.Contains to check for the value.
func (s *slice[T]) Has(value T) bool {
	return slices.Contains(s.inner, value)
}

// Add implements Lookup. It appends to the slice. It returns true if the value already existed in the slice.
func (s *slice[T]) Add(value T) bool {
	if s.Has(value) {
		return true
	}
	s.inner = append(s.inner, value)
	return false
}

// Remove implements Lookup. It removes a matching item from the slice. It is assumed
// that all elements in the slice are unique.
func (s *slice[T]) Remove(value T) bool {
	index := slices.Index(s.inner, value)
	if index < 0 {
		return false
	}
	// TODO Since only one element will ever be deleted, perhaps a custom implementation
	//      could be more efficient.
	s.inner = slices.Delete(s.inner, index, index+1)
	return true
}

// Len implements Lookup. It returns the length of the slice.
func (s *slice[T]) Len() int {
	return len(s.inner)
}

// Iter implements Lookup. It returns an iterator over the elements in the slice.
func (s *slice[T]) Iter() iter.Seq[T] {
	return slices.Values(s.inner)
}
