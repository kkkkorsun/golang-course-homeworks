package main

import "fmt"

func main() {
	const USDTOEURCOURSE float64 = 0.8566
	const USDTORUBCOURSE = 80.85

	var eurToExchange = 100
	var convertedValue = float64(eurToExchange) / USDTOEURCOURSE * USDTORUBCOURSE

	fmt.Printf("Converted: %.2f\n", convertedValue)
}

func getUserInput() float64 {
	var exchangeValue float64
	fmt.Println("Welcome to the currency exchange! Enter your currency value")
	fmt.Scanln(&exchangeValue)
	return exchangeValue
}

func exchangeCurrency(amount int, currentCurrency string, currencyToExchange string) {

}
