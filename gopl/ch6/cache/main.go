package main

import "sync"

var (
	mu      sync.Mutex
	mapping = make(map[string]string)
)

func LookUp(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

func main() {

}
