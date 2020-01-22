package main

import (
	"fmt"
)

const (
	secondInMS     float64 = 1000000 // ms
	breakPkg       float64 = 88      // ms
	markAfterBreak float64 = 12      // ms
	duration       float64 = 44      // ms
	stop           float64 = duration
	channels       float64 = 512
)

func main() {
	// [(120)+(12)+(44)+(CHL*44)+(0)+(50)] microsecs

	channelDuration := breakPkg + markAfterBreak + stop + channels*duration
	refreshRate := secondInMS / channelDuration

	fmt.Printf("Running on %v FPS refresh rate \n", refreshRate)

}
