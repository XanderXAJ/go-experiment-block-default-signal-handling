package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	for {
		s := <-c
		log.Println("Signal received:", s)
	}
}
