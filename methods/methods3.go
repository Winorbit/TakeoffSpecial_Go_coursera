package main

import "fmt"

func main() {
  egor := User{Name: "Egor", 
  						 EnglishLevel: 2}

  fmt.Println(egor.getName())
  fmt.Println(egor.getLevel())

  // fmt.Println(egor.getLevel())
  // egor.increaseLevel()
  // fmt.Println(egor.getLevel())

  // egor.decreaseLevel()
  // fmt.Println(egor.getLevel())
  // egor.decreaseLevel()
  // fmt.Println(egor.getLevel())
  // egor.decreaseLevel()
  // fmt.Println(egor.getLevel())
}


type User struct{
	Name string
	EnglishLevel int
	someFunc func()
}

func (u User) getName() string {
	return u.Name
}

func (u User) getLevel() int {
	return u.EnglishLevel
}


// func (u User) increaseLevel() {
// 	u.EnglishLevel = u.EnglishLevel+1
//  }


// Если мы не используем * - укзаатель, то обращаеся к струтуре не ПО АДРЕЕСУ, а ПО ЗНАЧЕНИЮ, то есть - к ее копии.
func (u *User) increaseLevel() {
	(*u).EnglishLevel = u.EnglishLevel+1
 }


func (u *User) decreaseLevel() {
	if u.EnglishLevel-1 > 0 {
	  (*u).EnglishLevel = u.EnglishLevel-1
  } else {
  	fmt.Printf("You level now is %v. There is nowhere to fall below :))", u.EnglishLevel)
  }
 }

// func (u User) updateLevel(newLevel int) string {
// 	return u.EnglishLevel