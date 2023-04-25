package main

import (
	"l2/develop/dev06/cut"
	"testing"
)

type flags struct {
	Delimiter string
	Fields    uint
	Separated bool
}

var tests = []struct {
	testName string
	flags    flags
	line     string
	result   string
}{
	{"default_test", flags{"\t", 2, false}, "common\ttest\tcase", "case"},
	{"colon_test", flags{":", 6, false}, "root:x:0:0::/root:/usr/bin/bash", "/usr/bin/bash"},
	{"separated_test", flags{"\t", 0, true}, "no_separator", ""},
	{"empty_test", flags{":", 4, false}, "root:x:0:0::/root:/usr/bin/bash", ""},
	{"space_delim_test", flags{" ", 5, false}, "drwxr-x---  10 root root  4096 Nov  2 21:12 root/", "Nov"},
}

func TestSortRun(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			t.Parallel()
			strs := cut.Cut(tt.line, tt.flags.Delimiter, tt.flags.Fields, tt.flags.Separated)
			if strs != tt.result {
				t.Fatalf("strings are not equal: %+#v %+#v", strs, tt.result)
			}
		})
	}
}
