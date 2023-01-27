package features

import (
	"bez-go-api/internal/core/entity"
	"bez-go-api/internal/core/repository"
)

type ProductDTO struct{}

type IProductService interface {
	GetProducts() ([]entity.Product, error)
}

type ProductService struct {
	ProductRepo repository.IProductRepository
}

func (p *ProductService) GetProducts() ([]entity.Product, error) {
	return p.ProductRepo.FindAll()
}
