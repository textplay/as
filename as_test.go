package as

import (
	"reflect"
	"testing"
)

func TestMustMap(t *testing.T) {
	in := []string{"a", "bb", "ccc"}
	want := []string{`"a"`, `"bb"`, `"ccc"`}
	quote := func(s string) string { return `"` + s + `"` }
	got := MustMap(in, quote)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MustMap(%s) == %s, want %s", in, got, want)
	}
}
