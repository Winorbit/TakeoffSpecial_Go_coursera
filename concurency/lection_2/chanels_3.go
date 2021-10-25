package main

import ("fmt"
        "time")

func main() {
  outcomeMessages := make(chan string)
  incomeMessages := make(chan string)

  // Закидываем в горутину функцию, обрабатывающую каналы. Она сработает только после заполнения его заполнения.
  go processMessage(incomeMessages, outcomeMessages)

  // Здесь запускаем процесс наполнения канала с исходящими сообщениями в еще одной горутине
  go func() {
    incomeMessages <- "Login"
    incomeMessages <- "Logout"
    incomeMessages <- "42"
    incomeMessages <- "VVVVVVV"
  }()

  // А вот тут, уже в самой функции main будем извлекать сообщения из output-канала
  fmt.Println(<-outcomeMessages)
  fmt.Println(<-outcomeMessages)
  fmt.Println(<-outcomeMessages)
  fmt.Println(<-outcomeMessages)
}

func processMessage(in <-chan string, out chan<- string) {
  fmt.Println("Initializing goroutine...")
  for {
    lastMessage := <-in
    timestamp := time.Now().Format(time.RFC850)
    resultMessage := fmt.Sprintf("Message {%v} was sent in %v", lastMessage, timestamp)
    out <- resultMessage
  }
}