package models

type Item struct {
	ID          int     `json:"item_id"`
	Category    string  `json:"category"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    int     `json:"descrption"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
}

func NewItem(id int, category, name, description string, unit_price float64, quantity int) *Item {
	total_price := unit_price * float64(quantity)

	return &Item{
		ID:          id,
		Category:    category,
		Name:        name,
		Description: description,
		UnitPrice:   unit_price,
		TotalPrice:  total_price,
		Quantity:    quantity,
	}
}
