package main

import (
	"fmt"
	"strings"
)

func main() {

	phrase := "well then let it be let me take face the heat";

	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	close(ch)

	for msg := range ch {
		fmt.Println(msg)
	}
	//for {
	//	if msg, ok := <- ch; ok {
	//		fmt.Println(msg)
	//	}else {
	//		break
	//	}
	//}
}
