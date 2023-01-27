package app

import (
	"bez-go-api/internal/app/features"
	"bez-go-api/internal/core"
	"bez-go-api/internal/core/repository"
)

type IContainer interface {
	NewProductService(productRepo repository.ProductRepository) features.ProductService
}

type Container struct {
	repo           core.IContainer
	ProductService features.IProductService
}

func NewContainer(_repo core.IContainer) *Container {
	return &Container{
		repo:           _repo,
		ProductService: NewProductService(_repo.NewProductRepository()),
	}
}

func NewProductService(_repo repository.IProductRepository) features.IProductService {
	return &features.ProductService{
		ProductRepo: _repo,
	}
}
