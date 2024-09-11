package models

type Item struct {
	Item_ID     int     `json:"item_id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Quantity    int     `json:"descrption"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
}

func NewItem(item_id int, category, name, description string, unit_price float64, quantity int) *Item {
	total_price := unit_price * float64(quantity)

	return &Item{
		Item_ID:     item_id,
		Category:    category,
		Name:        name,
		Description: description,
		UnitPrice:   unit_price,
		TotalPrice:  total_price,
		Quantity:    quantity,
	}
}
