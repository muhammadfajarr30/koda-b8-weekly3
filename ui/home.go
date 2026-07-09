package ui

import (
	"fmt"
	"koda-b8-weekly3/services"
	"koda-b8-weekly3/utils"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Home() {
	p := message.NewPrinter(language.Indonesian)
	menus := services.GetMenu()
	fmt.Println("===== SELAMAT DATANG DI GACOAN =====")
	fmt.Println("1. MAKANAN")
	fmt.Println("2. MINUMAN")
	fmt.Println("3. DIMSUM")
	fmt.Println("====================================")
	input := utils.Input("Masukkan Input: ")

	switch input {
	case "1":
		utils.Clear()
		fmt.Println("======= MAKANAN =======")
		for idx, menu := range menus {
			if menu.Category == "Makanan" {
				p.Printf("%d. %s - Rp.%d\n", idx+1, menu.Name, menu.Price)
			}
		}

	case "2":
		utils.Clear()
		fmt.Println("======= MINUMAN =======")
		for idx, menu := range menus {
			if menu.Category == "Minuman" {
				p.Printf("%d. %s - Rp.%d\n", idx+1, menu.Name, menu.Price)
			}
		}

	case "3":
		utils.Clear()
		fmt.Println("======= DIMSUM =======")
		for idx, menu := range menus {
			if menu.Category == "Dimsum" {
				p.Printf("%d. %s - Rp.%d\n", idx+1, menu.Name, menu.Price)
			}
		}
	}
}
