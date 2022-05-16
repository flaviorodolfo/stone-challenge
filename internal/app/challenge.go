package app

import (
	"stone-challenge/internal/model"
)

func calculateCartValues(list []model.ShoppingCart) int {
	total := 0
	for _, cart := range list {
		total += cart.Item.Price * cart.Quantity
	}
	return total
}

func divideTotalValue(totalValue int, totalPersons int) []int {

	personValue := totalValue / totalPersons
	remainder := totalValue % totalPersons
	var valuesList []int
	for i := 0; i < totalPersons; i++ {
		valuesList = append(valuesList, personValue)
		if remainder > 0 {
			valuesList[i]++
			remainder--
		}
	}
	return valuesList
}

func CalculateValues(cart []model.ShoppingCart, customer []model.Customer) map[string]int {
	totalValue := calculateCartValues(cart)
	valuesList := divideTotalValue(totalValue, len(customer))
	resultMap := make(map[string]int)
	for index, value := range valuesList {
		resultMap[customer[index].Email] = value
	}
	return resultMap
}
