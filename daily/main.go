package main

import (
	"github.com/cch123/supermonkey"
)

//go:noinline
func test1() {
	supermonkey.PatchByFullSymbolName(
		"123",
		func() string {
			return "111"
		})
}
func main() {
	test1()
}
