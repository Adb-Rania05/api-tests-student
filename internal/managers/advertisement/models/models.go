package models

// CreateAdvertisementRequest — структура запроса на создание объявления.
type CreateAdvertisementRequest struct {
	Title       *string
	Description *string
	Price       *float64
	Quantity    *int
	Photos      []string
}
