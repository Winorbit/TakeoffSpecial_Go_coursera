package main

import (
    "fmt"
    "time"
    "runtime"
    "net/http"
    "encoding/json"
)

const rootURL = "https://www.nbrb.by/api/exrates/rates/"
var paramMod = "?parammode=2"

type Currency struct {
  Scale int `json:"Cur_Scale"`
  Rate float64 `json:"Cur_OfficialRate"`
  Abbr string `json:"Cur_Abbreviation"`
}

func getCurrency(abbr string) Currency {
  currencyURL := fmt.Sprintf("%v%v%v", rootURL, abbr, paramMod)

  res, err := http.Get(currencyURL)
  defer res.Body.Close()

  if err != nil {
    fmt.Println(err.Error())
  }
  
  var currency Currency
  decoder := json.NewDecoder(res.Body)
  decoder.Decode(&currency)
  return currency}

func createCurrencyMessage(curr Currency) string {
  currencyAbbr := curr.Abbr
  scale := curr.Scale
  rate := curr.Rate

  message := fmt.Sprintf("Курс %v на сегодня: %v бел.руб за %v %v.", currencyAbbr, rate, scale, currencyAbbr)
  return message}

func main() {
  // runtime.GOMAXPROCS(10)
  runtime.GOMAXPROCS(2)
  // runtime.GOMAXPROCS(5)
  abbrs:= []string{"usd", "aud", "jpy","AMD","BGN","UAH", "EUR", "Pln", "ISK", "CAD"}

  start := time.Now()
  for _, value := range abbrs{ 
    go func(val string){
    currency := getCurrency(val)
    currMessage := createCurrencyMessage(currency)
    // go fmt.Println(currMessage)
    fmt.Println(currMessage)
  }(value)


  }
  elapsedTime := time.Since(start)
  fmt.Println("Total Time For Execution: " + elapsedTime.String())
  time.Sleep(time.Second)

}