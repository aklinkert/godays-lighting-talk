package main

import (
	"log"
	"time"
)

// RenderFrames defines the smallest render unit of a bar. Has to be multiplier of 4.
const RenderFrames uint8 = 64

// BarParams are a reusable informational struct on how fast and in what scheme something should be played
type BarParams struct {
	NoteValue uint8  `json:"noteValue" yaml:"noteValue"`
	NoteCount uint8  `json:"noteCount" yaml:"noteCount"`
	Speed     uint16 `json:"speed" yaml:"speed"`
}

// BarChange describes the changes of tempo and notes during a song
type BarChange struct {
	BarParams
	At uint64 `json:"at" yaml:"at"`
}

// CalcRenderSpeed calculates the render speed of a BarChange to a time.Duration
func CalcRenderSpeed(bc *BarChange) time.Duration {
	return time.Minute / time.Duration(bc.Speed*uint16(bc.NoteValue)/4) / time.Duration(RenderFrames/bc.NoteValue)
}

func main() {
	bc := &BarChange{
		At: 0,
		BarParams: BarParams{
			NoteValue: 32,
			NoteCount: 32,
			Speed:     195,
		},
	}

	renderSpeed := CalcRenderSpeed(bc)
	log.Printf("Rendering at %s timer duration (%v FPS)", renderSpeed, int(time.Second/renderSpeed))
}
