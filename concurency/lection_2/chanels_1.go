package main

import ("fmt"
        // "time"
      )

func main(){
  messages := []string{"Hello", "i am a log message", "Nothing important", "BBBB"}
  messagesChan := make(chan string) 
  // messagesChan := make(chan string, 2)

  
  go func(){messagesChan <- messages[0]}()
  go func(){messagesChan <- messages[1]}()
  go func(){messagesChan <- messages[2]}()
  go func(){messagesChan <- messages[3]}()

  fmt.Println(<-messagesChan)
  fmt.Println(<-messagesChan)
  fmt.Println(<-messagesChan)
  fmt.Println(<-messagesChan)


  /*
  for _, mess := range messages{
    go func(){messagesChan <- mess}()
  }

  fmt.Println(<-messagesChan)
  fmt.Println(<-messagesChan)
  fmt.Println(<-messagesChan)
  fmt.Println(<-messagesChan)
  */

  /*
  go func() {
    messagesChan <- messages[0]
    messagesChan <- messages[1]
    messagesChan <- messages[2]
    messagesChan <- messages[3]
  }()


  go fmt.Println(<-messagesChan)
  go fmt.Println(<-messagesChan)
  go fmt.Println(<-messagesChan)
  go fmt.Println(<-messagesChan)
  // go fmt.Println(<-messagesChan)

  time.Sleep(time.Second)
  */
}
