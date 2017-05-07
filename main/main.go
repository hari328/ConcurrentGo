package main

import (
	"fmt"
	"time"
)

func main() {

	godur,_ := time.ParseDuration("5ms")

	go func (){
		for i := 0; i < 100; i++ {
			fmt.Println("hello")
			time.Sleep(godur)
		}
	}()

	go func (){
		for i := 0; i < 100; i++ {
			fmt.Println("go")
			time.Sleep(godur)
		}
	}()

	dur,_ := time.ParseDuration("5s")
	time.Sleep(dur)
}
