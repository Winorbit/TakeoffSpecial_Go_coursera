package main

import "fmt"

func main() {
	fmt.Println("Please enter the acceleration, init. velocity and init. displacement separated by spaces or newlines - like 10, 2, 1: ")

	var acceleration float64
	var velocity_init float64
	var displacement_init float64 
	var time float64
	
	fmt.Scan(&acceleration)
	fmt.Scan(&velocity_init)
	fmt.Scan(&displacement_init)
	
	fn := GenDisplaceFn(acceleration, velocity_init, displacement_init)
	
	fmt.Println("Please enter the time: ")
	fmt.Scan(&time)	
	fmt.Println("The final displacement is", fn(time), "units.")
}

func GenDisplaceFn(a, v_o, s_o float64) func (float64) float64 {
  calc := func (t float64) float64 { return (0.5 * a * t * t + v_o * t + s_o)}
  return calc
}
