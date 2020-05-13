package main

import (
	"fmt"
	"time"
)

func main() {
 //1
	//go paralel()
	//time.Sleep(2)
	//fmt.Println("d")
//2
	//	go multiple.Numbers()
	//	go multiple.Alphabets()
	//	time.Sleep(3000 * time.Millisecond)
	//	fmt.Println("main terminated")
//3
//	s := []int{7, 2, 8, -9, 4, 0}
//
//	c := make(chan int)
//	go channels.Sum(s[:len(s)/2], c)
//	go channels.Sum(s[len(s)/2:], c)
//	x, y := <-c, <-c
//
//	fmt.Println(x, y, x+y)

//4
	//	done := make(chan bool)
	//	go hello(done)
	//	<- done
	//	fmt.Println(<- done)
//5
	//	done := make(chan bool)
	//	fmt.Println("Main going to call hello go goroutine")
	//	go hello2(done)
	//	<-done
	//	fmt.Println("Main received data")
//6
//	number := 589
//	sqrch := make(chan int)
//	cubech := make(chan int)
//	go calcSquares(number, sqrch)
//	go calcCubes(number, cubech)
//	squares, cubes := <-sqrch, <-cubech
//	fmt.Println("Final output", squares+cubes)
	//7
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v,"from ch")
		time.Sleep(2 * time.Second)

	}
}

func paralel(){
	fmt.Println("paralel")
}

func hello(done chan bool) {
	fmt.Println("Hello world goroutine")
	done <- true
}

func hello2(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeop <- sum
}


func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}