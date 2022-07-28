package main

import (
	"02_BasicTask/etime"
	"flag"
	"fmt"
	"os"
)

func main() {
	var host string
	flag.StringVar(&host, "h", "time.apple.com", "set host address")
	flag.Parse()
	tm, err := etime.GetExactTime(host)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(tm.String())
}
