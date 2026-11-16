package models

<<<<<<< HEAD
=======
// CreateAdvertisementRequest — структура запроса на создание объявления.
>>>>>>> b13fc60 (final)
type CreateAdvertisementRequest struct {
	Title       *string
	Description *string
	Price       *float64
	Quantity    *int
	Photos      []string
}
