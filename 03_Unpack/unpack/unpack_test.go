package unpack

import (
	"testing"
)

type cmpResult struct {
	data, want string
}

var list = []cmpResult{
	{"qwe\\4\\5", "qwe45"},
	{"qwe\\45", "qwe44444"},
	{"qwe\\\\5", "qwe\\\\\\\\\\"},
	{"a4bc2d5e", "aaaabccddddde"},
	{"\\\\1\\0", "\\0"},
	{"\\\\10", "\\\\\\\\\\\\\\\\\\\\"},
	{"a5\\\\", "aaaaa\\"},
	{"\\10", ""},
	{"abcd", "abcd"},
	{"45", ""},
	{"", ""},
}

func TestUnpack(t *testing.T) {
	for i, v := range list {
		v.data = Unpack(&v.data)
		if v.data != v.want {
			t.Errorf("test N%d failed\nwanted: %s\nexpecter: %s\n", i, v.want, v.data)
			return
		}
	}
}
