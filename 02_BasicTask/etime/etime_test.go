package etime

import (
	"testing"
)

func TestGetExactTime(t *testing.T) {
	tm, err := GetExactTime("time.apple.com")
	if err != nil || tm == nil {
		t.Fatal("error:", err, "time:", tm)
	}
	tm, err = GetExactTime("0.ru.pool.ntp.org")
	if err != nil || tm == nil {
		t.Fatal("error:", err, "time:", tm)
	}
	tm2, err2 := GetExactTime("WRONG.time.apple.com")
	if err2 == nil || tm2 != nil {
		t.Fatal("error:", err, "time:", tm)
	}
}
