package main

import "fmt"

func main() {
	sAnimal := []Animal{
		cat{name: "murzik", weight: 2},
		cat{name: "barsik", weight: 3},
		dog{name: "sharik", weight: 7},
		dog{name: "baton", weight: 10},
		cow{name: "manya", weight: 100},
	}

	fmt.Printf("\nTotal food needs for farm for 1 month: %.2f kg\n", calculateFood(sAnimal))
}

func calculateFood(animal []Animal) float64 {
	var foodNeedle float64
	for _, animal := range animal {
		foodNeedle += animal.monthFoodsNeed()
		fmt.Printf("\nName: %s , weight: %.1f kg, need food for 1 month: %.1f kg", animal.String(), animal.getWeight(), animal.monthFoodsNeed())
	}

	return foodNeedle
}
