package main

import "fmt"
import "time"

func nowMilliseconds() int {
	return int(1e-6 * float64(time.Now().UnixNano()))
}

func log(index, lifetime, referenceTime int) {
	now := nowMilliseconds()
	startTime := now - referenceTime - lifetime
	endTime := now - referenceTime

	fmt.Printf(
		"[%v]: (%.2fs, %.2fs)\n",
		index,
		0.001*float64(startTime),
		0.001*float64(endTime),
	)
}

func compute(index, lifetime int, notify chan int) {
	startTime := nowMilliseconds()
	elapsedTime := 0

	for elapsedTime < lifetime {
		elapsedTime = nowMilliseconds() - startTime
	}

	notify <- index
}

func main() {
	referenceTime := nowMilliseconds()
	lifetimes := []int{5000, 4000, 3000, 2000, 1000}
	notify := make(chan int)

	for i := 0; i < len(lifetimes); i++ {
		index := i
		lifetime := lifetimes[index]
		go compute(index, lifetime, notify)
	}

	for i := 0; i < len(lifetimes); i++ {
		index := <-notify
		log(index, lifetimes[index], referenceTime)
	}

	fmt.Println("Done")
}

/*
	[4]: (0.00s, 1.00s)
	[3]: (0.01s, 2.01s)
	[2]: (0.00s, 3.00s)
	[1]: (0.00s, 4.00s)
	[0]: (0.00s, 5.00s)
	Done
*/
