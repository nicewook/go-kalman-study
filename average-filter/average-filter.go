package main

import (
	"fmt"
	"math/rand"
)

var (
	lastAve           float32
	repeat            int
	notFirstOperation bool
)

func aveFilter(signalValue int) float32 {
	repeat++
	if !notFirstOperation {
		notFirstOperation = true
		lastAve = float32(signalValue)
		return lastAve
	}

	a := float32(1.0) / float32(repeat)
	fmt.Printf("%f  ", a)

	lastAve = (float32(1)-a)*lastAve + a*(float32(signalValue))
	return lastAve

}

func main() {
	loopCount := 10000
	generator := func() <-chan int {
		signalStream := make(chan int, loopCount)
		for i := 0; i < loopCount; i++ {
			signalStream <- rand.Intn(55)
		}
		return signalStream
	}

	signal := generator()

	for i := 0; i < loopCount; i++ {
		tmp := <-signal
		fmt.Printf("%f, \tsignal %d\n", aveFilter(tmp), tmp)
	}

}
