package model

import (
	"database/sql"
)

type Product struct {
	Id        string `json:"id"`
	Name      string `json:"name" binding:"required"`
	Price     int    `json:"price" binding:"required"`
	Stock     int    `json:"stock" binding:"required"`
	CreatedAt string `db:"created_at"`
	UpdatedAt sql.NullTime	`db:"updated_at"`
}
