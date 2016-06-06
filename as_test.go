package as

import (
	"reflect"
	"testing"
)

func slice1() []string {
	return []string{"one", "two", "three", "one", "two"}
}

func TestMustMap(t *testing.T) {
	in := slice1()
	want := []string{`"one"`, `"two"`, `"three"`, `"one"`, `"two"`}
	quote := func(s string) string { return `"` + s + `"` }
	got := MustMap(in, quote)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MustMap(%v) == %v, want %v", in, got, want)
	}
}

func TestIndicesOf(t *testing.T) {
	in := slice1()
	want := []int{1, 4}
	got := IndicesOf(in, "two")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("IndicesOf(%v, two) == %v, want %v", in, got, want)
	}
}

func TestIndicesOfPattern(t *testing.T) {
	in := slice1()
	want := []int{1, 2, 4}
	got := IndicesOfPattern(in, "^t")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("IndicesOfPattern(%v, two) == %v, want %v", in, got, want)
	}
}

func TestIndicesOfPattern2(t *testing.T) {
	in := slice1()
	want := []int{0, 2, 3}
	got := IndicesOfPattern(in, "^on|^th")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("IndicesOfPattern2(%v) == %v, want %v", in, got, want)
	}
}

func TestSplitAt(t *testing.T) {
	in := slice1()
	want := [][]string{[]string{"one", "two"}, []string{"three", "one"}, []string{"two"}}
	got := SplitAt(in, []int{2, 4})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SplitAt(%+v) = %+v, want %+v", in, got, want)
	}
}
