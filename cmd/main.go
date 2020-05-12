package main

import (
	"fmt"
	"time"
	"github.com/oktayalizada/goroutines"
)

func main() {
 //1
	//go paralel()
	//time.Sleep(2)
	//fmt.Println("d")
//2
	go multiple.Numbers()
	go multiple.Alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
	paralel()

}

func paralel(){
	fmt.Println("paralel")
	Numbers()
}



