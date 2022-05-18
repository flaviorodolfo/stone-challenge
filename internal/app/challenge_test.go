package app

import (
	"reflect"
	"stone-challenge/internal/model"
	"testing"
)

func TestCalculateValues(t *testing.T) {
	type args struct {
		cart     []model.ShoppingCart
		customer []model.Customer
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]int
		wantErr bool
	}{
		{name: "both empty lists",
			args: args{
				cart:     make([]model.ShoppingCart, 0),
				customer: make([]model.Customer, 0),
			},
			want:    nil,
			wantErr: true,
		},
		{name: "customer empty lists",
			args: args{
				cart: []model.ShoppingCart{
					model.ShoppingCart{
						Item: model.Item{
							Name:  "Bread",
							Price: 5,
						},
						Quantity: 1,
					},
				},
				customer: make([]model.Customer, 0),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "cart empty list",
			args: args{
				cart: make([]model.ShoppingCart, 0),
				customer: []model.Customer{
					model.Customer{
						Email: "user@stone.com",
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "both valid lists",
			args: args{
				cart: []model.ShoppingCart{
					model.ShoppingCart{
						Item: model.Item{
							Name:  "Beef",
							Price: 2451,
						},
						Quantity: 2,
					},
					model.ShoppingCart{
						Item: model.Item{
							Name:  "Mango",
							Price: 1200,
						},
						Quantity: 22,
					},
					model.ShoppingCart{
						Item: model.Item{
							Name:  "Salt",
							Price: 500,
						},
						Quantity: 3,
					},
				},
				customer: []model.Customer{
					model.Customer{
						Email: "user0@stone.com.br",
					},
					model.Customer{
						Email: "user1@stone.com.br",
					},
					model.Customer{
						Email: "user2@stone.com.br",
					},
					model.Customer{
						Email: "user3@stone.com.br",
					},
				},
			},
			want: map[string]int{
				"user0@stone.com.br": 8201,
				"user1@stone.com.br": 8201,
				"user2@stone.com.br": 8200,
				"user3@stone.com.br": 8200,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateValues(tt.args.cart, tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculateValues() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateCartValues(t *testing.T) {
	type args struct {
		list []model.ShoppingCart
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "list with one item",
			args: args{
				list: []model.ShoppingCart{

					model.ShoppingCart{
						Quantity: 1,
						Item: model.Item{
							Name:  "Rice",
							Price: 1000,
						},
					},
				},
			},
			want: 1000,
		},
		{
			name: "list with N items",
			args: args{
				list: []model.ShoppingCart{

					model.ShoppingCart{
						Quantity: 4,
						Item: model.Item{
							Name:  "Rice",
							Price: 1000,
						},
					},
					model.ShoppingCart{
						Quantity: 1,
						Item: model.Item{
							Name:  "Beef",
							Price: 5600,
						},
					},
					model.ShoppingCart{
						Quantity: 4,
						Item: model.Item{
							Name:  "Bread",
							Price: 500,
						},
					},
					model.ShoppingCart{
						Quantity: 3,
						Item: model.Item{
							Name:  "Bacon",
							Price: 2500,
						},
					},
				},
			},
			want: 19100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateCartValues(tt.args.list); got != tt.want {
				t.Errorf("calculateCartValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divideTotalValue(t *testing.T) {
	type args struct {
		totalValue   int
		totalPersons int
	}
	var tests = []struct {
		name string
		args args
		want []int
	}{
		{
			name: "both odd values",
			args: args{totalValue: 101,
				totalPersons: 3},
			want: []int{34, 34, 33},
		},
		{
			name: "both even values",
			args: args{totalValue: 124,
				totalPersons: 6},
			want: []int{21, 21, 21, 21, 20, 20},
		},
		{
			name: "even dividend odd divisor",
			args: args{totalValue: 3522,
				totalPersons: 11},
			want: []int{321, 321, 320, 320, 320, 320, 320, 320, 320, 320, 320},
		},
		{
			name: "odd dividend even divisor",
			args: args{totalValue: 7521,
				totalPersons: 8},
			want: []int{941, 940, 940, 940, 940, 940, 940, 940},
		},
		{
			name: "dividend and divisor with the same values",
			args: args{totalValue: 12,
				totalPersons: 12},
			want: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "dividend lower than divisor",
			args: args{totalValue: 2,
				totalPersons: 3},
			want: []int{1, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divideTotalValue(tt.args.totalValue, tt.args.totalPersons); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divideTotalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
