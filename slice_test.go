package lookup

import "testing"

func TestSliceHas(t *testing.T) {
	tests := []struct {
		name  string
		slice *slice[string]
		value string
		want  bool
	}{
		{
			name: "It has the element",
			slice: &slice[string]{
				inner: []string{"one", "two", "three"},
			},
			value: "two",
			want:  true,
		},
		{
			name: "It does not have the element",
			slice: &slice[string]{
				inner: []string{"one", "two", "three"},
			},
			value: "four",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.slice.Has(tt.value); got != tt.want {
				t.Fatalf(`Has(%v) = %v, want %v`, tt.value, got, tt.want)
			}
		})
	}
}

func TestSliceAdd(t *testing.T) {
	tests := []struct {
		name     string
		slice    *slice[int]
		add      int
		elements []int
		want     bool
	}{
		{
			name: "adds an existing element",
			slice: &slice[int]{
				inner: []int{1, 2, 3},
			},
			add:      3,
			elements: []int{1, 2, 3},
			want:     true,
		},
		{
			name: "adds a new element",
			slice: &slice[int]{
				inner: []int{1, 2, 3},
			},
			add:      4,
			elements: []int{1, 2, 3, 4},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.slice.Add(tt.add); got != tt.want {
				t.Errorf(`Add(%v) = %v, want %v`, tt.add, got, tt.want)
			}
			for _, element := range tt.elements {
				AssertSliceCount(t, tt.slice.inner, element, 1)
			}
		})
	}
}

func TestSliceRemove(t *testing.T) {
	tests := []struct {
		name     string
		slice    *slice[int]
		remove   int
		elements []int
		want     bool
	}{
		{
			name: "removes an existing element",
			slice: &slice[int]{
				inner: []int{1, 2, 3},
			},
			remove:   3,
			elements: []int{1, 2},
			want:     true,
		},
		{
			name: "removes an element that isn't present",
			slice: &slice[int]{
				inner: []int{1, 2, 3},
			},
			remove:   4,
			elements: []int{1, 2, 3},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.slice.Remove(tt.remove); got != tt.want {
				t.Errorf(`Remove(%v) = %v, want %v`, tt.remove, got, tt.want)
			}
			AssertSliceCount(t, tt.slice.inner, tt.remove, 0)
			for _, element := range tt.elements {
				AssertSliceCount(t, tt.slice.inner, element, 1)
			}
		})
	}
}

func TestSliceLen(t *testing.T) {
	slice := &slice[int]{
		inner: []int{1, 2, 3},
	}
	want := 3
	if got := slice.Len(); got != want {
		t.Fatalf(`Len() = %d, want %d`, got, want)
	}
}

func TestSliceIter(t *testing.T) {
	slice := &slice[int]{
		inner: []int{1, 2, 3},
	}
	want := map[int]int{
		1: 1,
		2: 1,
		3: 1,
	}

	AssertIterProduces(t, slice.Iter(), want)
}
