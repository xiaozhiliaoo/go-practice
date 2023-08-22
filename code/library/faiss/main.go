package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

func main() {

	d := 5  // dimension
	nb := 5 // database size
	nq := 2 // number of queries

	xb := make([]float32, d*nb)
	xq := make([]float32, d*nq)

	for i := 0; i < nb; i++ {
		for j := 0; j < d; j++ {
			xb[i*d+j] = rand.Float32()
		}
		xb[i*d] += float32(i) / 1000
	}

	for i := 0; i < nq; i++ {
		for j := 0; j < d; j++ {
			xq[i*d+j] = rand.Float32()
		}
		xq[i*d] += float32(i) / 1000
	}

	fmt.Print("database:%d\n", nb)
	fmt.Printf("query:%d\n", nq)
	fmt.Printf("database array:%s\n", jsonString(xb))
	fmt.Printf("query array:%s\n", jsonString(xq))
}

func jsonString(data []float32) string {
	marshal, _ := json.Marshal(data)
	return string(marshal)
}
