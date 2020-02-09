package memo_test

import (
	"go-practice/gopl/ch9/memo1"
	"go-practice/gopl/ch9/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

// cmd:go test -run=TestSequential -v go-practice/gopl/ch9/memo1
func TestSequential(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

// go test -run=TestConcurrent -race -v go-practice/gopl/ch9/memo1
func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
