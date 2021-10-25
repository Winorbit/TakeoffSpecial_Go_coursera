package main
import "fmt"
 
func main() {
    messages := make(chan string, 2)
    // messages := make(chan string)  
    messages <- "Login"
    messages <- "Go to cabinet"
    close(messages)
    // messages <- "Go to main menu"       // ошибка - канал уже закрыт
    fmt.Println(<-messages)  // Добавить в канал, после закрытия, ничего нельзя, а читать добавленные ранее сообщения - можно. 
    fmt.Println(<-messages)  
    fmt.Println(<-messages)  

    val, opened:= <-messages
    fmt.Println(val, opened)
}