package main

import (
	"fmt"
	"sync"
	"time"
)

type Philospher struct { // struct describing the philopher
	name      string
	leftFork  int
	rightFork int
}

var Philosphers = []Philospher{ // slice of philosphers
	{"Plato", 4, 0},
	{"Aristotle", 0, 1},
	{"Gais", 1, 2},
	{"Socrates", 2, 3},
	{"Shashwat", 3, 4},
}

var hungry = 3
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second

var orderMutex sync.Mutex // Mutex to handle order of eating
var order []string        // slice of string to store order of eating

func dine() {
	wg := &sync.WaitGroup{} // Waitgroup to synchronise the entire flow
	wg.Add(len(Philosphers))
	orderMutex = sync.Mutex{}

	seated := &sync.WaitGroup{} // Waitgroup to sync the seating of philosphers
	seated.Add(len(Philosphers))

	var forks = make(map[int]*sync.Mutex)

	for i := 0; i < len(Philosphers); i++ {
		forks[i] = &sync.Mutex{} //Each fork has a mutex on it
	}

	for i := 0; i < len(Philosphers); i++ {
		go diningPhiloshers(Philosphers[i], wg, forks, seated) // Each Philopher can eat conncurrently
	}
	wg.Wait() // The main goroutine thread will wait here, till the wg Waitgroup had been decremented.
}

func diningPhiloshers(philospher Philospher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done() // To mark the philospher done eating / gproutine ending
	fmt.Printf("%s is seated at the Table.\n", philospher.name)
	seated.Done() // To mark the current philospher as seated
	seated.Wait() // All sub gorutines will be stopped here until every philospher is seated

	for i := hungry; i > 0; i-- {
		// The below if else is required to avoid deadlock
		if philospher.leftFork > philospher.rightFork {
			forks[philospher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philospher.name)
			forks[philospher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philospher.name)
		} else {
			forks[philospher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philospher.name)
			forks[philospher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philospher.name)
		}
		fmt.Printf("\t%s has taken both the forks and is eating.\n", philospher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", philospher.name)
		time.Sleep(thinkTime)

		// When done we should unlock the forks
		forks[philospher.leftFork].Unlock()
		forks[philospher.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", philospher.name)
	}
	fmt.Println(philospher.name, "is satisified.")
	fmt.Println(philospher.name, "has left the table.")

	orderMutex.Lock()
	order = append(order, philospher.name)
	orderMutex.Unlock()
}

func main() {
	fmt.Println("Dining Philospher problem")
	fmt.Println("The table is empty.")
	dine()
	fmt.Println("The table is empty.")
	fmt.Println("The order of eating was :-")
	for i := 0; i < len(order); i++ {
		fmt.Printf("%s ", order[i])
	}
}
