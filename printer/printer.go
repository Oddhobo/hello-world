package printer

import "fmt"

type Message struct {
	Payload string
	Color   Color
}

//PrintMessages prints message entered in message channel
//ch
func PrintMessages(ch <-chan Message) {
	defer fmt.Printf(string(ResetColor) + "goodbye" + "\n")
	for message, open := <-ch; open; message, open = <-ch {
		fmt.Printf(string(message.Color) + message.Payload + "\n")
	}
}
