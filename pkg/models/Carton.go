package models

type Carton struct {
	Item     Item `json:"item"`
	Quantity int  `json:"quantity"`
}
