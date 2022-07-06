package main

import "fmt"

type animal interface {
	fmt.Stringer
	weighter
	foodNeeded
}

type weighter interface {
	Weight() float64
}

type foodNeeded interface {
	calculateFood() float64
}

func main() {
	var totalFoodNeeded float64
	animals := []animal{
		dog{weight: 5},
		dog{weight: 14},
		cat{weight: 3},
		cat{weight: 8},
		cow{weight: 714},
		cow{weight: 278},
	}

	for _, v := range animals {
		fmt.Printf("%s, weight: %.1fkg, food neede for 1 month: %.1fkg\n", v.String(), v.Weight(), v.calculateFood())
		totalFoodNeeded += v.calculateFood()
	}
	fmt.Printf("Total amount of food needed for all animals on the farm: %.1fkg\n", totalFoodNeeded)
}
