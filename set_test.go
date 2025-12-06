package lookup

import "testing"

func TestSetHas(t *testing.T) {
	tests := []struct {
		name  string
		set   *set[string]
		value string
		want  bool
	}{
		{
			name: "It has the element",
			set: &set[string]{
				inner: map[string]struct{}{
					"one":   {},
					"two":   {},
					"three": {},
				},
			},
			value: "two",
			want:  true,
		},
		{
			name: "It does not have the element",
			set: &set[string]{
				inner: map[string]struct{}{
					"one":   {},
					"two":   {},
					"three": {},
				},
			},
			value: "four",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Has(tt.value); got != tt.want {
				t.Fatalf(`Has(%v) = %v, want %v`, tt.value, got, tt.want)
			}
		})
	}
}

func TestSetAdd(t *testing.T) {
	tests := []struct {
		name     string
		set      *set[int]
		add      int
		elements []int
		want     bool
	}{
		{
			name: "adds an existing element",
			set: &set[int]{
				inner: map[int]struct{}{
					1: {},
					2: {},
					3: {},
				},
			},
			add:      3,
			elements: []int{1, 2, 3},
			want:     true,
		},
		{
			name: "adds a new element",
			set: &set[int]{
				inner: map[int]struct{}{
					1: {},
					2: {},
					3: {},
				},
			},
			add:      4,
			elements: []int{1, 2, 3, 4},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Add(tt.add); got != tt.want {
				t.Errorf(`Add(%v) = %v, want %v`, tt.add, got, tt.want)
			}
			for _, element := range tt.elements {
				AssertMapCount(t, tt.set.inner, element, 1)
			}
		})
	}
}

func TestSetRemove(t *testing.T) {
	tests := []struct {
		name     string
		set      *set[int]
		remove   int
		elements []int
		want     bool
	}{
		{
			name: "removes an existing element",
			set: &set[int]{
				inner: map[int]struct{}{
					1: {},
					2: {},
					3: {},
				},
			},
			remove:   3,
			elements: []int{1, 2},
			want:     true,
		},
		{
			name: "removes an element that isn't present",
			set: &set[int]{
				inner: map[int]struct{}{
					1: {},
					2: {},
					3: {},
				},
			},
			remove:   4,
			elements: []int{1, 2, 3},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.set.Remove(tt.remove); got != tt.want {
				t.Errorf(`Remove(%v) = %v, want %v`, tt.remove, got, tt.want)
			}
			AssertMapCount(t, tt.set.inner, tt.remove, 0)
			for _, element := range tt.elements {
				AssertMapCount(t, tt.set.inner, element, 1)
			}
		})
	}
}

func TestSetLen(t *testing.T) {
	set := &set[int]{
		inner: map[int]struct{}{
			1: {},
			2: {},
			3: {},
		},
	}
	want := 3
	if got := set.Len(); got != want {
		t.Fatalf(`Len() = %d, want %d`, got, want)
	}
}

func TestSetIter(t *testing.T) {
	set := &set[int]{
		inner: map[int]struct{}{
			1: {},
			2: {},
			3: {},
		},
	}
	want := map[int]int{
		1: 1,
		2: 1,
		3: 1,
	}

	AssertIterProduces(t, set.Iter(), want)
}
