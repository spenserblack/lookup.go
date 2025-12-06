package lookup

import (
	"math/rand"
	"testing"
)

func TestNew(t *testing.T) {
	threshold := 3

	t.Run("few elements creates a slice lookup", func(t *testing.T) {
		lookup := New(threshold, []int{1, 2, 3})
		slice, ok := lookup.(*slice[int])
		if !ok {
			t.Fatalf(`New() returned a %T, want *slice[int]`, lookup)
		}

		for _, want := range []int{1, 2, 3} {
			AssertSliceCount(t, slice.inner, want, 1)
		}
		// NOTE Assert no zero elements were added
		var zero int
		AssertSliceCount(t, slice.inner, zero, 0)
	})

	t.Run("many elements creates a set lookup", func(t *testing.T) {
		lookup := New(threshold, []int{1, 2, 3, 4})
		set, ok := lookup.(*set[int])
		if !ok {
			t.Fatalf(`New() returned a %T, want *set[int]`, lookup)
		}

		for _, want := range []int{1, 2, 3, 4} {
			AssertMapCount(t, set.inner, want, 1)
		}
		// NOTE Assert no zero elements were added
		var zero int
		AssertMapCount(t, set.inner, zero, 0)
	})
}

func TestWithSize(t *testing.T) {
	threshold := 3

	t.Run("small size creates a slice lookup", func(t *testing.T) {
		size := 3
		lookup := WithSize(threshold, size, 1)
		slice, ok := lookup.(*slice[int])
		if !ok {
			t.Fatalf(`WithSize() returned a %T, want *slice[int]`, lookup)
		}

		if got := cap(slice.inner); got != size {
			t.Errorf(`cap() = %d, want %d`, got, size)
		}

		for _, want := range []int{1} {
			AssertSliceCount(t, slice.inner, want, 1)
		}
		// NOTE Assert no zero elements were added
		var zero int
		AssertSliceCount(t, slice.inner, zero, 0)
	})

	t.Run("large size creates a set lookup", func(t *testing.T) {
		lookup := WithSize(threshold, 20, 1, 2)
		set, ok := lookup.(*set[int])
		if !ok {
			t.Fatalf(`WithSize() returned a %T, want *set[int]`, lookup)
		}

		for _, want := range []int{1, 2} {
			AssertMapCount(t, set.inner, want, 1)
		}
		// NOTE Assert no zero elements were added
		var zero int
		AssertMapCount(t, set.inner, zero, 0)
	})
}

func TestEmptyLookup(t *testing.T) {
	tests := []struct {
		name string
		elements []int
	}{
		{
			name: "empty slice elements",
			elements: []int{},
		},
		{
			name: "nil slice elements",
			elements: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lookup := New(3, tt.elements)
			if got := lookup.Has(rand.Int()); got {
				t.Errorf(`Has() = true, want false`)
			}
		})
	}
}
