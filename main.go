package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Oddhobo/hello-world/hello"
	"github.com/Oddhobo/hello-world/printer"
)

func generateNMessages(n int) []printer.Message {
	toReturn := make([]printer.Message, n)
	colors := []printer.Color{
		printer.Red,
		printer.BoldRed,
		printer.Blue,
		printer.BoldBlue,
		printer.Green,
		printer.BoldGreen,
		printer.Cyan,
		printer.BoldCyan,
		printer.Magenta,
		printer.BoldMagenta,
	}
	for i := 0; i < n; i++ {
		ci := rand.Intn(len(colors))
		toReturn[i] = printer.Message{
			Payload: fmt.Sprintf(hello.Hello()+" %v", i),
			Color:   colors[ci],
		}
	}
	return toReturn
}

func makeNoise(delay time.Duration, runTime time.Duration, m printer.Message, ch chan<- printer.Message) {
	stop := time.Now().Add(runTime)
	for i := time.Now(); i.Before(stop); i = time.Now() {
		ch <- m
		time.Sleep(delay)
	}

}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	messages := generateNMessages(6)
	messageChannel := make(chan printer.Message)
	go printer.PrintMessages(messageChannel)
	for _, v := range messages {
		pause := rand.Intn(4)*500 + 500
		v.Payload = fmt.Sprintf(v.Payload+" delay: %v", pause)
		go makeNoise(time.Duration(pause)*time.Millisecond, 10*time.Second, v, messageChannel)
	}
	time.Sleep(12 * time.Second)
	close(messageChannel)
	time.Sleep(10 * time.Millisecond)
}
