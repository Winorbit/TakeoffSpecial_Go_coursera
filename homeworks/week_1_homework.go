package main

import ("fmt"
		"strings"
		"strconv")


func main(){
	fmt.Print("Enter integers, divided by coma, like 3,15,16,65: ")

    inputData := scanInput()
    convertedNumbers := stringsToNumbers(inputData)

	fmt.Println(arr)
	BubbleSort(arr)
	fmt.Println(arr)
}


func Swap(v []int, i int){
	v[i], v[i+1] = v[i+1], v[i]
}

func BubbleSort(v []int){
	n := len(v)
	for i := 0; i < n - 1; i++{
		for j := 0; j < n - i - 1; j++{
			if v[j] > v[j + 1]{
				Swap(v, j)
			}
		}
	}
}


func scanInput() []string {
  var input string
  fmt.Scanln(&input)

  numbers := strings.Split(input, ",")
  return numbers
}


func stringsToNumbers(numbers []string) []int {
  var stringedNumbers = []int{}

    for _, i := range numbers {
        j, err := strconv.Atoi(i)
        if err != nil {
            panic(err)
        }
        stringedNumbers = append(stringedNumbers, j)
    }
    return stringedNumbers
}