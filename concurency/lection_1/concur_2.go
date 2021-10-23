package main

import ("fmt"
        "time"
        "runtime")
/*
С помощью этого GOMAXPROCS мы запрашиваем переход приложения на несколько ядер. 
И ключевые слова go, добавляющиеся перед исполнением функции, 
могут исполняться уже отдельно на разных ядрах, увеличивая производительность приложения.
*/

func main() {
  // Тут-то мы добавляем вместе с конкурентностью и параллелизм. 

  // runtime.GOMAXPROCS(10)
  // runtime.GOMAXPROCS(2)
  runtime.GOMAXPROCS(5)
  start := time.Now()
  go func() {
  // func() {
    for i:=0; i < 3; i++ {
      fmt.Println(i)
    }
  }()

  go func() {
    for i:=4; i < 10; i++ {
      fmt.Println(i)
    }
  }()

  elapsedTime := time.Since(start)

  fmt.Println("Total Time For Execution: " + elapsedTime.String())
  // go fmt.Println("Total Time For Execution: " + elapsedTime.String())

  time.Sleep(time.Second)
}