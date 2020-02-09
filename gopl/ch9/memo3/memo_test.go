package memo_test

import (
	"go-practice/gopl/ch9/memo3"
	"go-practice/gopl/ch9/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

// cmd:go test -run=TestSequential -v go-practice/gopl/ch9/memo2
func TestSequential(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

// go test -run=TestConcurrent -race -v go-practice/gopl/ch9/memo2
func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
