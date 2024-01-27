package channel

import "fmt"

func UsageChannel() {

	var changelist = make(chan int, 10)

	go func() {

		for i := range changelist {
			fmt.Println(i)
		}

		//if len(changelist) < cap(changelist) {
		//	fmt.Println(<-changelist)
		//}
	}()

	for i := 0; i < 100; i++ {
		changelist <- i
	}
	fmt.Println(len(changelist), cap(changelist))
}
