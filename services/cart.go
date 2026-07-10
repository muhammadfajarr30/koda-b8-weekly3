package services

import "koda-b8-weekly3/models"

var CartItems []models.Cart

func AddToCart(menu models.Menu, qty int) {
	item := models.Cart{
		Menu:     menu,
		Quantity: qty,
	}
	CartItems = append(CartItems, item)
}

