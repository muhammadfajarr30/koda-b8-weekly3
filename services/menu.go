package services

import (
	"encoding/json"
	"fmt"
	"koda-b8-weekly3/models"
	"os"
)

var menu []models.Menu

func GetMenu() []models.Menu {
	data, err := os.ReadFile("data/data.json")
	if err != nil {
		fmt.Println("gagal membaca data", err)
	}
	err = json.Unmarshal(data, &menu)
	if err != nil {
		fmt.Println("gagal parse JSON", err)
	}
	return menu
}

func GetMenuByCategory(category string) []models.Menu {
	dataMenu := GetMenu()
	var result []models.Menu
	for _, menu := range dataMenu {
		if menu.Category == category {
			result = append(result, menu)
		}
	}
	return result
}
