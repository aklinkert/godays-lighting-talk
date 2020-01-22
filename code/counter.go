package main

import (
	"flag"
	"fmt"
	"sync"
)

type storage struct {
	sync.Mutex
	counter uint64
}

func (s *storage) IncDeferred() {
	s.Lock()
	defer s.Unlock()
	s.counter++
}

func (s *storage) IncSync() {
	s.Lock()
	s.counter++
	s.Unlock()
}

const max = 10000000 // 10 million

func main() {
	var useDefer = flag.Bool("defer", false, "Use defered function call")
	flag.Parse()

	store := &storage{}

	if *useDefer {
		for i := 0; i < max; i++ {
			store.IncDeferred()
		}
	} else {
		for i := 0; i < max; i++ {
			store.IncSync()
		}
	}

	fmt.Println(store.counter)
}
