package main

import (
	"fmt"
	"log"
	"os"
	"stone-challenge/internal/app"
	"stone-challenge/internal/model"
)

func printMap(result map[string]int) {
	fmt.Println("Map result:")
	for key, value := range result {
		fmt.Printf("Key: %v Value: %v\n", key, value)
	}

}

func main() {
	exitCode := 0
	var cartExample = []model.ShoppingCart{

		model.ShoppingCart{
			Quantity: 1,
			Item: model.Item{
				Name:  "Rice",
				Price: 1537,
			},
		},
		model.ShoppingCart{
			Quantity: 5,
			Item: model.Item{
				Name:  "Sugar",
				Price: 1252,
			},
		},
		model.ShoppingCart{
			Quantity: 10,
			Item: model.Item{
				Name:  "Coke",
				Price: 1200,
			},
		},
		model.ShoppingCart{
			Quantity: 7,
			Item: model.Item{
				Name:  "Bacon",
				Price: 3500,
			},
		},
	}
	var customersExample = []model.Customer{
		model.Customer{
			Email: "user0@stone.com.br",
		},
		model.Customer{
			Email: "user1@stone.com.br",
		},
		model.Customer{
			Email: "user2@stone.com.br",
		},
	}
	result, err := app.CalculateValues(cartExample, customersExample)

	if err != nil {
		exitCode = 1
		log.Fatal(err)
	} else {
		printMap(result)
	}
	os.Exit(exitCode)
}
