package main

import (
	"fmt"
	"sync"
)

func printSomething(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(msg)
}

func main() {
	words := []string{"eit", "bit", "tres"}
	wg := sync.WaitGroup{}
	wg.Add(len(words))
	for _, v := range words {
		go printSomething(v, &wg)
	}
	wg.Wait()
}
