package main

import (
	"fmt"
	"sync"
)

func concurrentSquare(nums []int) {
	wg := &sync.WaitGroup{}

	for _, v := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num * num)
		}(v)
	}

	wg.Wait()
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	concurrentSquare(nums)
}
