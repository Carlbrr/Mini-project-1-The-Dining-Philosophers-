package main

import "sync"

type Fork struct {
	timesUsed int
	outgoing  chan int
	incoming  chan int
	locker    sync.Mutex
}

func (f Fork) forkReciever() {

	for {
		call := <-f.incoming

		//Dette virker ligegyldigt, da der kun kan være en værdi i denne channel.
		//Måske man ikke bør tage en værdi, men bare bede den om at inkremere timesUsed.

		if call == 1 {
			f.timesUsed++
			f.outgoing <- f.timesUsed
		}
	}
}
