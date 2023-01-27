package repository

import (
	"bez-go-api/internal/core/entity"
	"database/sql"
	"fmt"
)

type IProductRepository interface {
	FindAll() ([]entity.Product, error)
	Create(product entity.Product) (*entity.Product, error)
}

type ProductRepository struct {
	DB *sql.DB
}

// Create implements IProductRepository
func (*ProductRepository) Create(product entity.Product) (*entity.Product, error) {
	panic("unimplemented")
}

func (o *ProductRepository) FindAll() ([]entity.Product, error) {
	fmt.Println(o.DB)
	orders := []entity.Product{
		{
			Id:   1,
			Name: "grocery",
		}, {
			Id:   2,
			Name: "kitchen",
		},
		{
			Id:   3,
			Name: "furniture",
		},
	}

	return orders, nil
}
