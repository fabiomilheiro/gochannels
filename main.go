package main

import (
	"gochannels/utils"
	"log"
	"math/rand"
)

type Thing struct {
	Text  string
	Value int
}

var handlingCount map[string]int

func main() {
	handlingCount = map[string]int{}
	channel := make(chan Thing)

	go handleThings("A", channel)
	go handleThings("B", channel)
	go handleThings("C", channel)

	i := 0
	for {
		thing := Thing{
			Text:  utils.GenerateString(),
			Value: rand.Int(),
		}

		log.Printf("Sending message %v", thing.Text)

		channel <- thing

		i++

		if i%100 == 0 {
			log.Printf("Stats")
			log.Printf("Handler\tCount")
			for key, count := range handlingCount {
				log.Printf("%v\t%v", key, count)
			}
			utils.DelayLong()
		}

		utils.Delay()
	}
}

func handleThings(handlerName string, channel chan Thing) {
	for {
		thing := <-channel

		// TODO: Fix concurrent map assignment
		handlingCount[handlerName]++
		log.Printf("[Handler %v] Message received: %v, Value: %v", handlerName, thing.Text, thing.Value)
	}
}
