package template

import "fmt"

type CaffeineBeverageInterface interface {
	Brew()
	AddCondiments()
	CustomerWantsCondiments() bool
}

type CaffeineBeverage struct{}

func (c *CaffeineBeverage) PrepareRecipe(beverage CaffeineBeverageInterface) {
	c.boilWater()
	beverage.Brew()
	c.pourInCup()

	if beverage.CustomerWantsCondiments() {
		beverage.AddCondiments()
	}
}

func (c *CaffeineBeverage) boilWater() {
	fmt.Println("Hirviendo el agua")
}

func (c *CaffeineBeverage) pourInCup() {
	fmt.Println("Vertiendo en la taza")
}

func (c *CaffeineBeverage) CustomerWantsCondiments() bool {
	return true
}
