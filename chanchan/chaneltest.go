// package main
package chanchan

import (
	"fmt"
	"strconv"
	"time"
)

func main2() {
	///////////////////////////////////////////////
	// for i := 0; i < 10; i++ {
	// 	go f(i)
	// }
	// var input string
	// fmt.Scanln(&input)
	///////////////
	//채널은 큐다
	var ch chan string = make(chan string)
	go pinger(ch)
	go ponger(ch)
	go printer(ch)
	var input string
	fmt.Scanln(&input)
}

//	func f(n int) {
//		for i := 0; i < 10; i++ {
//			fmt.Println(n, ":", i)
//			amt := time.Duration(rand.Intn(250))
//			time.Sleep(time.Millisecond * amt)
//		}
//	}
//
// func pinger(c chan<- string)  //넣는거만 가능
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping" + strconv.Itoa(i)
	}
}
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong" + strconv.Itoa(i)
	}
}

// func pinger(c chan<- string)  //빼기만 가능
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 1000)
	}
}
