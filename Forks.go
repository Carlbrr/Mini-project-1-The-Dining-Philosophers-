package main

import "sync"

type Fork struct {
	timesUsed int
	inUse     bool
	outgoing  chan string
	incoming  chan string
	locker    sync.Mutex
}
