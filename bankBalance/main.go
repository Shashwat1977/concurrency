package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	var bankBalance int
	mx := sync.Mutex{}

	incomes := []Income{
		{Source: "Main", Amount: 500},
		{Source: "Invesetment", Amount: 100},
	}
	wg.Add(len(incomes))
	for i, income := range incomes {
		go func(i int, income Income) {
			for week := 1; week <= 52; week++ {
				mx.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				mx.Unlock()
			}
			wg.Done()
		}(i, income)
	}
	wg.Wait()
	fmt.Printf("The final amount is %d\n", bankBalance)
}
