package services

import (
	"koda-b8-weekly3/models"
)

var CartItems []models.Cart

func AddToCart(menu models.Menu, qty int) {
	for i := range CartItems {
		if CartItems[i].Menu.ID == menu.ID {
			CartItems[i].Quantity += qty
			return
		}
	}
	item := models.Cart{
		Menu:     menu,
		Quantity: qty,
	}
	CartItems = append(CartItems, item)
}

func GetOrderdItems() []models.Cart {
	return CartItems
}
