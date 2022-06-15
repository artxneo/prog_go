package main

import (
	"fmt"
	"math"
)

const APPLE_PRICE float64 = 5.99
const PEAR_PRICE int = 7
const MONEY int = 23

func main() {

	applesToBuy := 9
	pearsToBuy := 8

	fmt.Printf("Скільки грошей потрібно, щоб купити 9 яблук та 8 груш? \n %.2f грн \n", float64(applesToBuy)*APPLE_PRICE+float64(pearsToBuy*PEAR_PRICE))
	fmt.Printf("Скільки груш ми можемо купити? \n %v груш \n", MONEY/PEAR_PRICE)
	fmt.Printf("Скільки яблук ми можемо купити? \n %v яблок \n", math.Floor(MONEY/APPLE_PRICE))
	applesToBuy = 2
	pearsToBuy = 2
	response := ""
	if float64(MONEY) > (float64(applesToBuy)*APPLE_PRICE + float64(pearsToBuy*PEAR_PRICE)) {
		response = "Так"
	} else {
		response = "Ні"
	}
	fmt.Printf("Чи ми можемо купити 2 грущі та 2 яблока? \n %v \n", response)
}
