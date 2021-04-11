package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// Item is Alfred's item struct.
type Item struct {
	Type     string `json:"type"`
	Icon     string `json:"icon"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}

// Menu is Alfred's menu struct.
type Menu struct {
	Items []Item `json:"items"`
}

func outputFormat(item Item) {
	var menu Menu
	item.Icon = "./icon.png"
	menu.Items = append(menu.Items, item)

	menuJSON, _ := json.Marshal(menu)
	fmt.Println(string(menuJSON))
}

func main() {
	uuidObj, _ := uuid.NewRandom()
	uuid := uuidObj.String()

	var item Item
	item.Title = uuid
	item.Arg = uuid

	outputFormat(item)
}
