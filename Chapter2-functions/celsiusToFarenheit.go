package main

import "fmt"

func celsiusToFarenheit(celsius float64) float64 {
	return celsius*(9/5) + 32
}

func main() {
	var celsius float64
	_, err := fmt.Scan(&celsius)
	if err != nil {
		return
	}
	fahrenheit := celsiusToFarenheit(celsius)

	fmt.Printf("Temperature in Farenheit: %f \n", fahrenheit)
}
