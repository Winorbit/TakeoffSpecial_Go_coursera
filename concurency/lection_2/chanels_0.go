package main

import "fmt"

func main(){
  var messagesChan chan string = make(chan string) // канал для данных типа string
  // messagesChanel := make(chan string) // канал для данных типа string

  go func(){
    fmt.Println("Go routine starts")
    messagesChan <- "Hello" // Анонинмная функция в горутине закидывает данные в канал messagesChan
    }()

  fmt.Println(<-messagesChan) // Здесь функция main блокируется до того момента, пока не из горутины не прилетит сообщение и его можно будет извлечь.
  fmt.Println("The End")
}
