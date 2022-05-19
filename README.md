# Stone Challenge

### About the challenge
The challenge is a go project developed for Stone hiring test, the main goal of the project was to create a function that calculate the amount that one or more customers will pay (in cents) according to a shopping list.
With the objective of keeping easy understanding and future writing of the code some concepts of the book Código Limpo by Robert C. Martin were used.

### How to Start
To start the application you need execute the main function, for that you need to run the command ``go run main.go`` in the main file dir (``\stone-challenge\cmd\challenge``).

If you want to change the results you need to edit the input list, in order to do this you need to change the lists ``cartExample`` and ``customersExample`` inside the ``main.go`` file. The price of the Item must be give in cents, eg: if you want to put 5 reais you need to multiply by 100 to get 500.

Here is some examples of those lists:
```go
//list with one item
var cartExample = []model.ShoppingCart{

    model.ShoppingCart{
        Quantity: 7,
        Item: model.Item{
        Name:  "Bacon",
        Price: 3500,//price in cents
        },
    },
}
// list with two items

var cartExample = []model.ShoppingCart{

    model.ShoppingCart{
        Quantity: 7,
        Item: model.Item{
            Name:  "Rice",
            Price: 3500, //price in cents
        },
    },
    model.ShoppingCart{
        Quantity: 12,
		Item: model.Item{
        Name:  "Bread",
        Price: 4500, //price in cents
        },
    },
}
//list with two customers
var customersExample = []model.Customer{
    model.Customer{
        Email: "user0@stone.com.br",
	},
    model.Customer{
        Email: "user1@stone.com.br",
    },
}
//empty list
var customersExample []model.Customer
```

In order to run the tests you need to execute the command ``go test`` int the dir(``\stone-challenge\internal\app``) folder of the ``challenge_test.go`` file.

### License
MIT License

Copyright (c) 2022 Flávio Rodolfo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
IDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
