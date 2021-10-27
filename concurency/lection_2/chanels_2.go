package main

import "fmt"

func main(){
  messagesChan := make(chan string)

  go sendMessageToChanel("Hello!", messagesChan)
  go sendMessageToChanel("Hello! 2", messagesChan)
  go sendMessageToChanel("Hello! 3", messagesChan)
  
  fmt.Println(<-messagesChan)
  fmt.Println(<-messagesChan)
  fmt.Println(<-messagesChan)
}

// func sendMessageToChanel(message string, out chan<- string){
func sendMessageToChanel(message string, out chan string){
  out <- message
 }
