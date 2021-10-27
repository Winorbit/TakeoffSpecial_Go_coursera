package main
import (
        "fmt"
        "sync"
        )
 
var counter int = 0             //  общий ресурс

func main() {
    chann := make(chan bool)       
    var mutex sync.Mutex        // определяем мьютекс
    for i := 1; i < 5; i++{
        go work(i, ch, &mutex)
    }
     
    for i := 1; i < 5; i++{
        <-ch
    }
     
    fmt.Println("The End")
}
func increase_counter (number int, ch chan bool, mutex *sync.Mutex){
    mutex.Lock()    // блокируем доступ к переменной counter
    counter = 0
    for value := 1; value <= 5; value++{
        counter++
        fmt.Println("Goroutine", number, "-", counter)
    }
    mutex.Unlock()  // деблокируем доступ
    ch <- true
}