package channel

import (
	"fmt"
	"time"
)

type Request interface {
}

func handle(r Request) {
	fmt.Println(r.(int))
}

const RateLimitPeriod = time.Minute

const RateLimit = 100

func handleRequests(requests <-chan Request) {
	quotas := make(chan time.Time, RateLimit)

	go func() {
		tick := time.NewTicker(RateLimitPeriod / RateLimit)
		defer tick.Stop()
		for t := range tick.C {
			select {
			case quotas <- t:
			default:

			}
		}
	}()

	for request := range requests {
		<-quotas
		go handle(request)
	}
}

func RunMain() {
	requests := make(chan Request)
	go handleRequests(requests)

	for i := 0; i < 500; i++ {
		requests <- i
	}
}

func Calous1() {
	for i := 0; i < 10; i++ {
		func() {
			fmt.Println(i)
		}()
	}
}

func Calous2() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func Calous3() {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second * 2)
}

func Calous4() {
	for i := 0; i < 100; i++ {
		go func(index int) {
			fmt.Println(index)
		}(i)
	}
	time.Sleep(time.Second * 2)
}
