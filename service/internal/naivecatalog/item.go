package naivecatalog

import (
	"encoding/json"
)

type Item struct {
	Id               string `json:"productId"`
	Name             string `json:"productName"`
	Brand            string `json:"Brand"`
	ShortDescription string `json:"metaTagDescription"`
	Link             string `json:"link"`
	ImageUrl         string
	Json             string
}

func NewItem(payload string) (Item, error) {
	jsonData := []byte(payload)
	var item Item
	err := json.Unmarshal(jsonData, &item)
	if err == nil {
		item.Json = payload
	}
	return item, err
}
