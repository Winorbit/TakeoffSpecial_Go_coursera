package main

import "fmt"

func main() {
  egor := User{Name: "Egor", EnglishLevel: 2}
  fmt.Println(egor.Name)
  egor.sayHello()
  egor.presntMe()
  egor.printWeather("Lviv")
  weatheMess := egor.weatherMessage("Lviv")
  fmt.Println(weatheMess)
}

type User struct{
	Name string
	EnglishLevel int
}

func (u User) sayHello(){
	fmt.Println("Hello")
}

// использум значения, заданные в Struct внутри нашего метода.
// Это его главное отличие от обычной функции - метод находиться в зоне видимости type, и имеет доступ к его значениям 
func (u User) presntMe(){
	fmt.Printf("Hello, my name is %v, and my english level is %v \n", u.Name, u.EnglishLevel)
}

// И, конечно, так, как это самая обычная функция, она принимает аргументы.
// Например, можно задать вот такой метод, который сообщает данные о погоде в заднном городе.
func (u User) printWeather(city string){
	fmt.Printf("Hello, %v, today in  %v fog all day \n", u.Name, city)
}


// Разумеется, и return так же работает.
func (u User) weatherMessage(city string) string{
	message := fmt.Sprintf("Hello, %v, today in  %v fog all day \n", u.Name, city)
	return message
}