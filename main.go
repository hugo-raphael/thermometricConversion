package main

import "fmt"

const (
	tempK float64 = 227.80
	zeroK float64 = 273.15
)

func main() {
	tempC := tempK - zeroK

	fmt.Printf("%.2f°K em graus Celsius é %.2f°C\n", tempK, tempC)

}
