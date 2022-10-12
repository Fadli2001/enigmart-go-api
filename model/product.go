package model 

type Product struct {
	Id string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Price int `json:"price" binding:"required"`
	Stock int	`json:"stock" binding:"required"`
}