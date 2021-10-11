package main

import "fmt"

func main(){
	// Присвоение функции переменной
	fmt.Println(helloFunc)
	fmt.Println(helloFunc())

	funcVar := helloFunc
	funcVar()

	// Анонимные функции

	//anon := func(){fmt.Println("Hello, i'm anon func!")}
	//anon
	//anon()

	// Функции возвращающие и принимающие функции
	// returnFunc()
	returnFuncWithArg()("Egor")

	// decorator
	printSomething("hello world")

	// with decorator
	decorator(printSomething)("yolo")

}


func helloFunc() string {
  return "Hello"
}


func returnFunc() func() {
  somfunc := func(){fmt.Println("Hello, i'am inside func")}
  return somfunc
}


func returnFuncWithArg() func(string) {
  somfunc := func(name string){fmt.Println(name)}
  return somfunc
}



// DECORATORS

func printSomething(s string) {
	fmt.Println(s)
}

func decorator(incomeFunc func(arg string)) func(string) {
	return func(arg string) {
		fmt.Println("Hello, i print before func is run!")
		incomeFunc(arg)
		fmt.Println("I'm - after!")
	}
}


