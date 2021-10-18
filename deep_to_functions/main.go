package main

import "fmt"

func main(){
	// Присвоение функции переменной
	// fmt.Println(   helloFunc    )
	// fmt.Println(   helloFunc()  )

	// funcVar := helloFunc
	// funcVar()
	// fmt.Println(funcVar())

	// Анонимные функции

	// anon := func(){fmt.Println("Hello, i'm anon func!")}
	// func anon(){fmt.Println("Hello, i'm anon func!")}

	// fmt.Println(anon)
	// fmt.Println(anon())

	// Функции возвращающие и принимающие функции
	// returnFunc()()

	// returnFuncWithArg()("Egor")

	// decorator
	// printSomething("hello world")

	// with decorator
	//decorator(printSomething)("yolo")

	// defer
	checkDefer()

}


func helloFunc() string {
  fmt.Println("Hello")
  return "Hello"
}

// returnFunc()

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

// DEFERS
// defer - ключевое слово, которое говорит, что дейстие после него нужно выполнить ПОСЛЕ выполнения внешней функции

func checkDefer() {
	i := 1
	defer fmt.Println(i+1)
	i++
	fmt.Println("Hello!")
}
