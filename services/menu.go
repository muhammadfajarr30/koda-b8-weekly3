package services

import (
	"encoding/json"
	"fmt"
	"koda-b8-weekly3/models"
	"os"
)

func GetMenu() []models.Menu {
	var menu []models.Menu
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
