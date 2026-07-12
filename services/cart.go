package services

import (
	"fmt"
	"koda-b8-weekly3/models"
	"sync"
	"time"
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
func CalculateTotal() int {
	total := 0
	for _, item := range CartItems {
		total += item.Menu.Price * item.Quantity
	}
	return total
}

func ClearCart() {
	CartItems = []models.Cart{}
}

var wg = sync.WaitGroup{}

func CookMeal() {
	for _, item := range CartItems {
		wg.Add(1)

		go func(item models.Cart) {
			defer wg.Done()

			fmt.Printf("Menyiapkan %s...\n", item.Menu.Name)
			time.Sleep(time.Duration(item.Quantity) * time.Second)

			fmt.Println()
			fmt.Printf("%s selesai!\n", item.Menu.Name)
		}(item)
	}

	wg.Wait()
	fmt.Println("\nSemua pesanan selesai.")
}
