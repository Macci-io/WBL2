package etime

import (
	"github.com/beevik/ntp"
	"time"
)

//GetExactTime return current exact time
func GetExactTime(host string) (*time.Time, error) {
	tm, ok := ntp.Time(host)
	if ok != nil {
		return nil, ok
	}
	return &tm, nil
}
