package main

import "fmt"

func readLogs(c chan string) {
    for {
        log := <-c
        fmt.Println(log)
    }
}

func main() {
    fmt.Println("main() started")
    logsChan := make(chan string, 5)

    go readLogs(logsChan)

    logsChan <- "login"
    logsChan <- "error"
    logsChan <- "warning"
    logsChan <- "login"

    fmt.Println("main() stopped")
}