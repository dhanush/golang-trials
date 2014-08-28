package main

import (
	"fmt"
)

/* 
	A program that will chat with you until you enter X.
	Demostrates the use of go routines
*/


func server (ch chan string, msg string) {
	fmt.Println("Message Entered", msg)
	ch <- "Hello " +msg
}

func client(ch chan string){
	serverMsg := <-ch
	fmt.Println("Messsage from Server ", serverMsg)
}

func main(){
	var ch chan string  = make(chan string)
	var input string
	fmt.Println("Say Hi to the Chat Server. X will exit the Program")
	for i:=0;;i++ {
		fmt.Scanln(&input)
		if input=="X" { break}
		go server(ch, input)
		go client(ch)
	}



}


