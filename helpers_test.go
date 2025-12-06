package lookup

import (
	"iter"
	"testing"
)

// AssertSliceCount asserts that the needle appears in the haystack the expected amount of
// times.
func AssertSliceCount[V comparable](t *testing.T, haystack []V, needle V, expected int) {
	t.Helper()

	count := 0
	for _, hay := range haystack {
		if hay == needle {
			count++
		}
	}

	if count != expected {
		t.Errorf(`%v appears in %v %d times, want %d times`, needle, haystack, count, expected)
	}
}

// AssertMapCount asserts that the needle appears in the haystack the expected amount of
// times.
func AssertMapCount[V comparable, Ignored any](t *testing.T, haystack map[V]Ignored, needle V, expected int) {
	t.Helper()

	count := 0
	for hay := range haystack {
		if hay == needle {
			count++
		}
	}

	if count != expected {
		t.Errorf(`%v appears in %v %d times, want %d times`, needle, haystack, count, expected)
	}
}

// AssertIterProduces assert that the iterator produces the wanted values. want should
// map the value to the expected number of times it should appear.
func AssertIterProduces[V comparable](t *testing.T, seq iter.Seq[V], want map[V]int) {
	t.Helper()
	got := make(map[V]int, len(want))

	for produced := range seq {
		got[produced]++
	}

	for produced, got := range got {
		want := want[produced]

		if got != want {
			t.Errorf(`received %v %d times, wanted %d times`, produced, got, want)
		}
	}

	for wanted := range want {
		if _, ok := got[wanted]; !ok {
			t.Errorf(`wanted to receive %v, but didn't receive it`, wanted)
		}
	}
}
