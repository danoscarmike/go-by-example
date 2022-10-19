package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	timer1 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer1.C
		elapsed := time.Now().Sub(start)
		fmt.Println("Timer 1 fired", elapsed)
	}()

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(3 * time.Second)
}
