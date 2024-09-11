package models

type Warehouse struct {
	Location        string `json:"location"`
	CurrentCapacity int    `json:"current_capacity"`
	TotalCapacity   int    `json:"total_capacity"`
}

func NewWarehouse(location string, current_capacity, total_capacity int) *Warehouse {

	return &Warehouse{
		Location:        location,
		CurrentCapacity: current_capacity,
		TotalCapacity:   total_capacity,
	}
}
