package main

import (
	"l2/develop/dev03/strsort"
	"testing"
)

type flags struct {
	Key     int
	Numeric bool
	Reverse bool
	Unique  bool
}

var tests = []struct {
	testName string
	flags    flags
	lines    []string
	result   []string
}{
	{"common_test", flags{0, false, false, false}, []string{"c", "a", "d", "b", "f", "e"}, []string{"a", "b", "c", "d", "e", "f"}},
	{"numeric_test", flags{0, true, false, false}, []string{"1", "3", "2", "91", "55", "4"}, []string{"1", "2", "3", "4", "55", "91"}},
	{"reverse_test", flags{0, false, true, false}, []string{"c", "a", "d", "b", "f", "e"}, []string{"f", "e", "d", "c", "b", "a"}},
	{"unique_test", flags{0, false, false, true}, []string{"c", "a", "d", "c", "a", "d"}, []string{"a", "c", "d"}},
	{"key_test", flags{1, false, false, false}, []string{"0 c", "0 a", "0 d", "0 b", "0 f", "0 e"}, []string{"0 a", "0 b", "0 c", "0 d", "0 e", "0 f"}},
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestSortRun(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			t.Parallel()
			strs := strsort.Sort(tt.lines, tt.flags.Key, tt.flags.Numeric, tt.flags.Reverse, tt.flags.Unique)
			if !stringSlicesEqual(strs, tt.result) {
				t.Fatalf("slices are not equal: %+#v %+#v", strs, tt.result)
			}
		})
	}
}
