package main

import (
	"log"
	"time"
)

func main() {
	go updater()
	sendLoop()
}

// new local time wrapping artnet.Address
type Address struct {
	Net      int
	lastSent time.Time // time this address was last updated
	stale    bool      // indicates whether this address data has not yet been sent on the wire

	data [512]byte
}

// this is the max fps in interval
var interval = time.Second / 44

// this is the minimum update interval as designed in the spec
var minInterval = 300 * time.Millisecond

// convenience time far in the past
var epoch = time.Unix(0, 0)

// a temp array holding some addresses
var addresses = []*Address{
	{

		Net:   1,
		stale: true,
	},
	{
		Net:   2,
		stale: true,
	},
	{
		Net: 3,
	},
	{
		Net: 4,
	},
}

// to replace the dmxUpdate logic
func sendLoop() {
	tick := time.NewTicker(interval)
	for range tick.C {
		now := time.Now()
		for _, a := range addresses {
			if a.stale && a.lastSent.Before(now.Add(-interval)) {
				log.Printf("fps tick for %d", a.Net)
				a.lastSent = now
				a.stale = false
				continue
			}
			if a.lastSent.Before(now.Add(-minInterval)) {
				log.Printf("periodic tick for %d", a.Net)
				a.lastSent = now
				a.stale = false
				continue
			}
		}
	}
}

// send some input a lot faster then max fps
func updater() {
	tick := time.NewTicker(10 * time.Millisecond)
	for range tick.C {
		addresses[1].lastSent = epoch
		addresses[1].stale = true
	}
}
