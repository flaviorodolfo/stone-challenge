package app

import (
	"errors"
	"stone-challenge/internal/model"
)

// calculateCartValues calculates the total price of all model.Item in a list of model.ShoppingCart
func calculateCartValues(list []model.ShoppingCart) int {
	total := 0
	for _, cart := range list {
		total += cart.Item.Price * cart.Quantity
	}
	return total
}

// divideTotalValue divide the given value(in cents) by the total number of persons.
func divideTotalValue(totalValue int, totalPersons int) []int {

	// by the integer division decimals will be truncated
	// eg: 15/4 will return 3 and not 3.75.
	personValue := totalValue / totalPersons
	//the remainder will represent the sum of truncated decimals
	remainder := totalValue % totalPersons
	var valuesList []int
	//create a list of values for each person
	// given that a mod b = c will always give b>c then there is no need to iterate above b
	for i := 0; i < totalPersons; i++ {
		valuesList = append(valuesList, personValue)
		//check if the remainder is bigger than 0 then add +1 in the current value
		//and reduce the remainder by -1
		if remainder > 0 {
			valuesList[i]++
			remainder--
		}
	}
	return valuesList
}

// CalculateValues uses calculateCartValues and divideTotalValue
// returns  (map[string]int,nil) or (nil,error) in cases of invalid lists
func CalculateValues(cart []model.ShoppingCart, customer []model.Customer) (map[string]int, error) {

	//checks the length of both lists to avoid 0/0 and N/0 division
	if len(cart) == 0 {
		return nil, errors.New("the list of items is empty")
	}
	if len(customer) == 0 {
		return nil, errors.New("the list of persons is empty")
	}
	totalValue := calculateCartValues(cart)
	valuesList := divideTotalValue(totalValue, len(customer))
	resultMap := make(map[string]int)
	//builds the map where the key is the customer's email and the value is the customer's value(price)
	for index, value := range valuesList {
		resultMap[customer[index].Email] = value
	}
	return resultMap, nil
}
