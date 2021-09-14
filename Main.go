package main

func main() {

	var Forks = createForks()

	var phil1 = Philosopher{number: 1, left: &Forks[0], right: &Forks[1]}
	var phil2 = Philosopher{number: 2, left: &Forks[1], right: &Forks[2]}
	var phil3 = Philosopher{number: 3, left: &Forks[2], right: &Forks[3]}
	var phil4 = Philosopher{number: 4, left: &Forks[3], right: &Forks[4]}
	var phil5 = Philosopher{number: 5, left: &Forks[4], right: &Forks[0]}
}

func createForks() []Fork {
	Forks := make([]Fork, 5)
	for i := 0; i < 5; i++ {
		Forks[i] = Fork{timesUsed: 0, outgoing: make(chan string), incoming: make(chan string), inUse: false}
	}
	return Forks
}
