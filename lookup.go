// Package lookup provides utilities for checking for the presence of a value.
package lookup

import "iter"

// Lookup is a type that can check for the presence of a value.
type Lookup[T comparable] interface {
	// Has returns true if the value exists in the lookup.
	Has(value T) bool
	// Add adds a value to the lookup, and returns true if it already existed.
	Add(value T) bool
	// Remove removes a matching value from the lookup if it exists. It returns true if
	// the value was found.
	Remove(value T) bool
	// Len gets the length of the collection.
	Len() int
	// Iter produces an iterator over all the values in the lookup.
	Iter() iter.Seq[T]
}

// New creates a new Lookup. The threshold is compared to the number of elements passed
// to determine how the elements will be stored.
func New[T comparable](threshold int, elements []T) Lookup[T] {
	size := len(elements)
	if size <= threshold {
		return newSliceLookup(elements)
	}
	return newSetLookupWithSize(size, elements...)
}

// WithSize creates a new Lookup. The threshold is compared to the size to determine
// how the elements will be stored. The backing slice or map will have a size set to
// size.
func WithSize[T comparable](threshold int, size int, elements ...T) Lookup[T] {
	if size <= threshold {
		return newSliceLookupWithSize(size, elements...)
	}
	return newSetLookupWithSize(size, elements...)
}
