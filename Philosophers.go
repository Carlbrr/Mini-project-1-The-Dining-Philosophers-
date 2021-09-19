package main

import (
	"time"
)

type Philosopher struct {
	timesEaten int
	left       *Fork
	right      *Fork
	outgoing   chan int
	incoming   chan int
}

func (p Philosopher) eat() {
	for {

		p.right.locker.Lock()
		p.left.locker.Lock()

		// 1 betyder at nøglen skal starte
		// Se note i Forks.go filen
		p.right.incoming <- 1
		p.left.incoming <- 1

		p.timesEaten++
		p.outgoing <- p.timesEaten

		time.Sleep(time.Millisecond * 2000)

		// Antallet af gange, der er spist kan aldrig være -1
		// -1 betyder derfor at filosoffen er færdig med at spise
		p.outgoing <- -1

		p.right.locker.Unlock()
		p.left.locker.Unlock()

		time.Sleep(time.Millisecond * 500)
	}
}
