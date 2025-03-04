package dto

type CreateIcecreamRequest struct {
	Name           string  `json:"name" binding:"required,min=5"`
	Ingredients    string  `json:"ingredients" binding:"required"`
	ProductionDate string  `json:"production_date" binding:"required"`
	BestBefore     string  `json:"best_before" binding:"required"`
	Price          float64 `json:"price" binding:"required,gt=0"`
	Quantity       int     `json:"quantity" binding:"required,gt=0"`
}

type CreateIcecreamResponse struct {
	ID             int     `json:"id"`
	Name           string  `json:"name" binding:"required,min=5"`
	Ingredients    string  `json:"ingredients" binding:"required"`
	ProductionDate string  `json:"production_date" binding:"required"`
	BestBefore     string  `json:"best_before" binding:"required"`
	Price          float64 `json:"price" binding:"required,gt=0"`
	Quantity       int     `json:"quantity" binding:"required,gt=0"`
}
