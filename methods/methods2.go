package main

import "fmt"

func main() {
  egor := User{Name: "Egor", 
  						EnglishLevel: 2,
  						// Можно и вот так объявить функцию(метод)
  						someFunc: func(){fmt.Println("Hello, i'm anon func inside struct")}}

  egor.someFunc()
  // Можно сочетать оба варианта - и так, как someFunc, и так, как sayHello
  egor.sayHello()
}


type User struct{
	Name string
	EnglishLevel int
	someFunc func()
}

func (u User) sayHello(){
	fmt.Println("Hello")
}