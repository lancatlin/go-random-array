package main

import (
	"math/rand"
	"flag"
	"fmt"
)

func shuffle(src []int) {
	seed := getSeed()
	fmt.Printf("Random seed: %X\n", seed)
	rand.Seed(seed)
	rand.Shuffle(len(src), func(i, j int) {
		src[i], src[j] = src[j], src[i]
	})
}

func generateList(begin, end int) []int {
	src := make([]int, end - begin + 1)
	for i := begin; i <= end; i ++ {
		src[i-begin] = i
	}
	return src
}

func main() {
	var begin, end int
	flag.IntVar(&begin, "b", 1, "begin number")
	flag.IntVar(&end, "e", 100, "end number")
	flag.Parse()
	if begin > end {
		fmt.Println("begin number cannot bigger than end number")
		return
	}
	src := generateList(begin, end)
	shuffle(src)
	fmt.Println(src)
}
