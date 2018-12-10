package main

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
	"os"
)



func showMessage(messages chan string, signals chan bool){

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
				
			case <-signals:
			
				fmt.Println("Received kill signal ...")
				os.Exit(1)
				
					
		}
	}

}


func main() {
	messages := make(chan string,5)
	signals := make(chan bool)
	
	go showMessage(messages,signals)
	
	//Producer
	for{
		for i := 0; i < 20; i++ {
			r := rand.Intn(10)
			messages <- strconv.Itoa(r)
			time.Sleep(time.Duration(r) * time.Second)		
		}
	
		signals<-true
	}
}
