package main

import (
	"l2/develop/dev05/gogrep"
	"testing"
)

var tests = []struct {
	testName string
	flags    gogrep.Flags
	lines    []string
	pattern  string
	result   string
}{
	{"common_test", gogrep.Flags{}, []string{"test\n", "not\n", "radical\n", "do\n", "random\n", "yeet\n"}, "ra", "radical\nrandom\n"},
	{"ignore_test", gogrep.Flags{IgnoreCase: true}, []string{"test\n", "not\n", "radical\n", "do\n", "rAndom\n", "yeet\n"}, "RA", "radical\nrAndom\n"},
	{"after_test", gogrep.Flags{After: 3}, []string{"test\n", "not\n", "radical\n", "do\n", "random\n", "yeet\n"}, "test", "test\nnot\nradical\ndo\n"},
	{"before_test", gogrep.Flags{Before: 3}, []string{"test\n", "not\n", "radical\n", "do\n", "random\n", "yeet\n"}, "do", "test\nnot\nradical\ndo\nrandom\n"},
	{"context_test", gogrep.Flags{Context: 2}, []string{"test\n", "not\n", "radical\n", "do\n", "random\n", "yeet\n"}, "radical", "test\nnot\nradical\ndo\nrandom\n"},
	{"count_test", gogrep.Flags{Count: true}, []string{"test\n", "not\n", "radical\n", "do\n", "random\n", "yeet\n"}, "t", "3\n"},
	{"invert_test", gogrep.Flags{Invert: true}, []string{"test\n", "not\n", "radical\n", "do\n", "random\n", "yeet\n"}, "t", "radical\ndo\nrandom\n"},
	{"fixed_test", gogrep.Flags{Fixed: true}, []string{"test\n", "not\n", "radical\n", "do\n", "random\n", "yeet\n"}, "do\n", "do\n"},
	{"lineNum_test", gogrep.Flags{LineNum: true}, []string{"test\n", "not\n", "radical\n", "do\n", "random\n", "yeet\n"}, "do", "4:do\n5:random\n"},
}

func TestSortRun(t *testing.T) {
	for _, tt := range tests {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			//t.Parallel()
			str := gogrep.Grep(tt.lines, tt.pattern, tt.flags)
			if str != tt.result {
				t.Fatalf("strings are not equal: %+#v %+#v", str, tt.result)
			}
		})
	}
}
