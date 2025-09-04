package main

import "fmt"

func main() {
	startExchange()
}

func startExchange() {
	var currency string
	for {
		fmt.Println("Введите валюту которую вы хотите обменять\nНапример EUR, USD или RUB")
		fmt.Scanln(&currency)

		if currency == "EUR" || currency == "USD" || currency == "RUB" {
			break
		}

		fmt.Println("Такой валюты нет, попробуйте снова из предложенного списка")
	}

	var currencyAmount int64
	for {
		fmt.Printf("Введите количество валюты которую вы хотите обменять: ")
		fmt.Scanln(&currencyAmount)

		if currencyAmount > 0 {
			break
		}
		fmt.Println("Количество должно быть больше 0, попробуйте снова")
	}

	getCurrencyToExchange(currency, currencyAmount)
}

func getCurrencyToExchange(userCurrency string, amount int64) {
	var currencyToExchange string
	for {
		fmt.Println("Введите валюту на которую вы хотите обменять\nНапример EUR, USD или RUB")
		fmt.Scanln(&currencyToExchange)

		if currencyToExchange == userCurrency {
			fmt.Println("Введите отличное от вашей валюты значение")
			continue
		}

		if currencyToExchange != "EUR" && currencyToExchange != "USD" && currencyToExchange != "RUB" {
			fmt.Println("Такой валюты нет, попробуйте снова из предложенного списка")
			continue
		}

		break
	}

	exhangedCurrency := exchangeCurrency(amount, userCurrency, currencyToExchange)
	fmt.Printf("После обмена %s на %s вы получили %.2f\n", userCurrency, currencyToExchange, exhangedCurrency)
}

func exchangeCurrency(amount int64, currentCurrency string, currencyToExchange string) float64 {
	const USDTOEUR float64 = 0.8566
	const USDTORUB float64 = 80.85

	switch currentCurrency {
	case "RUB":
		switch currencyToExchange {
		case "USD":
			return float64(amount) / USDTORUB
		case "EUR":
			return float64(amount) / USDTORUB * USDTOEUR
		}
	case "EUR":
		switch currencyToExchange {
		case "USD":
			return float64(amount) / USDTOEUR
		case "RUB":
			return float64(amount) / USDTOEUR * USDTORUB
		}
	case "USD":
		switch currencyToExchange {
		case "EUR":
			return float64(amount) * USDTOEUR
		case "RUB":
			return float64(amount) * USDTORUB
		}
	}

	return 0
}
