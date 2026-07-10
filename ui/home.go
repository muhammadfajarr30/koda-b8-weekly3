package ui

import (
	"fmt"
	"koda-b8-weekly3/services"
	"koda-b8-weekly3/utils"
	"os"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func Home() {

	for {
		utils.Clear()
		fmt.Println("========SELAMAT DATANG DI GACOAN ========")
		fmt.Println("1. Order Makanan")
		fmt.Println("2. Keranjang")
		fmt.Println("Checkout")
		fmt.Println("------------------------------------------")
		fmt.Println("0 Exit")

		input := utils.Input("Pilih Menu: ")
		switch input {
		case "1":
			displayCategory()
		case "2":
			displayCart()
		case "0":
			os.Exit(0)
		}
	}
}

func displayCart() {
	p := message.NewPrinter(language.Indonesian)
	utils.Clear()
	fmt.Println("=========== KERANJANG ===========")
	cart := services.GetOrderdItems()

	if len(cart) == 0 {
		fmt.Println("keranjang kamu masih kosong silakan pesan menu dulu")
	} else {

		for idx, viewCart := range cart {
			p.Printf("Pesanan anda:\n %d. %s  (%d) - Rp. %d\n", idx+1, viewCart.Menu.Name, viewCart.Quantity, viewCart.Menu.Price)
		}
	}
	utils.Input("\nTekan Enter Untuk Kembali... ")
}

func displayCategory() {
	p := message.NewPrinter(language.Indonesian)
	utils.Clear()
	fmt.Println("===== SILAKAN PILIH KATEGORI =====")
	fmt.Println("1. MAKANAN")
	fmt.Println("2. MINUMAN")
	fmt.Println("3. DIMSUM")
	fmt.Println("------------------------------------")
	fmt.Println("====================================")
	input := utils.Input("Masukkan Input: ")

	switch input {
	case "1":
		utils.Clear()
		fmt.Println("======= MAKANAN =======")
		foods := services.GetMenuByCategory("Makanan")
		for idx, menu := range foods {
			if menu.Category == "Makanan" {
				p.Printf("%d. %s - Rp.%d\n", idx+1, menu.Name, menu.Price)
			}
		}
		inputMenu := utils.Input("Pilih menu : ")
		menuIndex, err := strconv.Atoi(inputMenu)
		if err != nil {
			fmt.Println("Input harus berupa angka")
			return
		}
		if menuIndex < 1 || menuIndex > len(foods) {
			fmt.Println("Pilh sesuai dengan menu yang ada")
		}
		selectedMenu := foods[menuIndex-1]
		qtyInput := utils.Input("Jumlah: ")
		qty, err := strconv.Atoi(qtyInput)
		if err != nil {
			fmt.Print("Jumlah harus berupa angka")
		}
		services.AddToCart(selectedMenu, qty)
	case "2":
		utils.Clear()
		fmt.Println("======= MINUMAN =======")
		drinks := services.GetMenuByCategory("Minuman")
		for idx, menu := range drinks {
			if menu.Category == "Minuman" {
				p.Printf("%d. %s - Rp.%d\n", idx+1, menu.Name, menu.Price)
			}
		}
		inputMenu := utils.Input("Pilih menu : ")
		menuIndex, err := strconv.Atoi(inputMenu)
		if err != nil {
			fmt.Println("Input harus berupa angka")
			return
		}
		if menuIndex < 1 || menuIndex > len(drinks) {
			fmt.Println("Pilh sesuai dengan menu yang ada")
		}
		selectedMenu := drinks[menuIndex-1]
		qtyInput := utils.Input("Jumlah: ")
		qty, err := strconv.Atoi(qtyInput)
		if err != nil {
			fmt.Print("Jumlah harus berupa angka")
		}
		services.AddToCart(selectedMenu, qty)
	case "3":
		utils.Clear()
		fmt.Println("======= DIMSUM =======")
		sideDish := services.GetMenuByCategory("Dimsum")
		for idx, menu := range sideDish {
			if menu.Category == "Dimsum" {
				p.Printf("%d. %s - Rp.%d\n", idx+1, menu.Name, menu.Price)
			}
		}
		inputMenu := utils.Input("Pilih menu : ")
		menuIndex, err := strconv.Atoi(inputMenu)
		if err != nil {
			fmt.Println("Input harus berupa angka")
			return
		}
		if menuIndex < 1 || menuIndex > len(sideDish) {
			fmt.Println("Pilh sesuai dengan menu yang ada")
		}
		selectedMenu := sideDish[menuIndex-1]
		qtyInput := utils.Input("Jumlah: ")
		qty, err := strconv.Atoi(qtyInput)
		if err != nil {
			fmt.Print("Jumlah harus berupa angka")
		}
		services.AddToCart(selectedMenu, qty)

	default:
		fmt.Println("Pilih sesuai menu")
	}
}
