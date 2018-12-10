package main

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
)



func showMessage(messages chan string){

	var buffer []string //Buffer de mensajeria

	for{
		//Consumer
		select{
			case msg:=<-messages:
				
				buffer=append(buffer,msg)
				fmt.Printf("[Msg]:%s [bufferSize]: %d\n",msg,len(buffer))
				
			case <-time.After(8 * time.Second):
				fmt.Println("\nNo message received for 8 seconds")
				buffer = buffer[:0]	
		}
	}

}


func main() {
	messages := make(chan string,5)
	
	go showMessage(messages)
	
	//Producer
	
	for{
		r := rand.Intn(10)
		messages <- strconv.Itoa(r)
		time.Sleep(time.Duration(r) * time.Second)
			
	}
}
