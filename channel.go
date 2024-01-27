package main

import "fmt"

// 管道模式：协程+管道 https://weread.qq.com/web/reader/6e832a20813ab752ag0154c9kd6432e00228d645920e3401

// 实现给定切片子项平方求和
// 1. 生成数组，2. 求平方，3. 求和

func GenerateArr(max int) <-chan int {
	out := make(chan int, max)
	go func() {
		for i := 0; i < max; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func Square(arr <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for v := range arr {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func Sum(arr <-chan int) <-chan int {
	out := make(chan int, 1)
	go func() {
		var sum = 0
		for v := range arr {
			sum += v
		}
		out <- sum
		close(out)
	}()

	return out
}

func RunPipe() {
	arr := GenerateArr(4)
	arrSquare := Square(arr)
	sum := <-Sum(arrSquare)
	fmt.Println(sum)
	//fmt.Println(<-arr)
}
