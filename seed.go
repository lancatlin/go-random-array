package main

import (
	"crypto/rand"
)

func getSeed() int64 {
	data := make([]byte, 8)
	_, err := rand.Read(data)
	if err != nil {
		panic(err)
	}
	var sum int64
	for i, v := range data {
		sum += int64(v) << (int64(i) * 8)
	}
	return sum
}
