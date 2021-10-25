package main

import "fmt"

func main(){
  logs := getLogsFromService()
  logsChanel := make(chan string) 
  go streamLogs(logs, logsChanel)
 
  for{log, opened := <- logsChanel    // получаем данные из потока
      if !opened{break}    // если поток закрыт, выход из цикла
      fmt.Println(log)
   }
}


func streamLogs(logs []string, ch chan string){
    defer close(ch)
    for _, log := range logs{
        ch <- log // наша функция разбирает влетевший массив логов и 
    }
}

/*
Представим, что эта функция берет логи из указанного сервиса за определнный период времени, 
и прокидывает их в брокер сообщений, а тот уже перенаправляет их еще куда-то.
*/
func getLogsFromService() []string {
  logs := []string{"login", "logout", "error", "warning", "warning", "login"}
  return logs
}