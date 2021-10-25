package main

import (
    "fmt"
    "time"
)


func service1(c chan string) {
    // time.Sleep(3 * time.Second)
    c <- "Hello from service 1 \n"
}

func service2(c chan string) {
    // time.Sleep(5 * time.Second)
    c <- "Hello from service 2 \n"
}

func main() {
    start := time.Now()
    fmt.Println("main() started", time.Since(start))

    chan1 := make(chan string)
    chan2 := make(chan string)

    go service1(chan1)

    go service2(chan2)


    select {
    case res := <-chan1:
        fmt.Println("Response from service 1 \n", res, time.Since(start))
    case res := <-chan2:
        fmt.Println("Response from service 2 \n", res, time.Since(start))
    }

    fmt.Println("main() stopped", time.Since(start))
}