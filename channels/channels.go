package main

import (
	"fmt"
	"time"
)

func main() {
	primeirExemplo()
	bufferizacaoChannel()
	selectChannel()
}

func primeirExemplo() {
	mensagens := make(chan string)
	go func() { mensagens <- "ping" }()
	msg := <-mensagens
	fmt.Println(msg)
}

func bufferizacaoChannel() {
	mensagens := make(chan string, 2)
	mensagens <- "buffer1"
	mensagens <- "buffer2"

	fmt.Println(<-mensagens)
	fmt.Println(<-mensagens)
}

func selectChannel() {
	ch1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "resultado"
	}()

	select {
	case result := <-ch1:
		fmt.Println(result)
	case <-time.After(time.Second * 3):
		fmt.Println("Timeout")

	}

}
