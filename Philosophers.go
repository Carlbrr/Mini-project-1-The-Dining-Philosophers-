package main

import (
	"time"
)

type Philosopher struct {
	timesEaten int
	left       *Fork
	right      *Fork
	isEating   bool
	outgoing   chan int
	incoming   chan int
}

func (p *Philosopher) eat() {
	for {
		p.right.locker.Lock()
		p.left.locker.Lock()

		p.isEating = true

		p.timesEaten++

		time.Sleep(time.Millisecond * 500)

		// 1 betyder at nøglen skal starte
		// Se note i Forks.go filen
		p.right.incoming <- 0
		p.left.incoming <- 0

		time.Sleep(time.Millisecond * 1500)

		p.isEating = false

		p.right.locker.Unlock()
		p.left.locker.Unlock()

		time.Sleep(time.Millisecond * 500)
	}
}

func (p *Philosopher) philReceiver() {

	for {
		var call = <-p.incoming

		if call == 1 {
			if p.isEating {
				p.outgoing <- p.timesEaten
			} else {
				p.outgoing <- -1
			}
		}
	}
}
