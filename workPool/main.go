package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker(taskch <-chan int, resultch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range taskch {
		result := process(n)
		resultch <- result
	}
}

func main() {

	var wg sync.WaitGroup
	//taskch
	taskch := make(chan int)
	//resultch
	resultch := make(chan int)

	//done
	done := make(chan struct{})

	//记录程序开始时间
	startTime := time.Now()

	//open goroutine and input the tasks into the task channel
	go func() {
		for i := 0; i < 1000; i++ {
			taskch <- i
		}
		close(taskch)
	}()

	//create worker goroutines for processing the data
	for range 100 {
		wg.Add(1)
		go Worker(taskch, resultch, &wg)
	}

	//print the value in the result channel
	go func() {
		for v := range resultch {
			fmt.Println("current value is :", v)
		}
		close(done)
	}()

	wg.Wait()
	close(resultch)

	<-done

	//记录程序结束时间并计算运行时长
	elapsedTime := time.Since(startTime)
	fmt.Printf("Program ran for: %v\n", elapsedTime)
}

func process(num int) int {
	time.Sleep(2 * time.Second)
	return num * 2
}
