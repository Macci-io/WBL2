package main

import (
	"fmt"
	"grep/pkg/grep"
)

func main() {
	fmt.Print(grep.NewGrep().Run())
}
