package core

import (
	"bez-go-api/internal/core/repository"
	"database/sql"
)

type IContainer interface {
	NewProductRepository() repository.IProductRepository
}

type Container struct {
	// db here
	_db *sql.DB
}

func NewContainer(db *sql.DB) IContainer {
	return &Container{
		// set db here
		_db: db,
	}
}

func (c *Container) NewProductRepository() repository.IProductRepository {
	return &repository.ProductRepository{
		DB: c._db,
	}
}
