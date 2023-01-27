package controllers

import (
	"bez-go-api/internal/app"

	"github.com/gin-gonic/gin"
)

type IProductController interface {
	GetAllProducts() string
}

type ProductController struct {
	_service *app.Container
}

func NewProductController(handler *gin.RouterGroup, service *app.Container) {
	p := &ProductController{
		_service: service,
	}
	h := handler.Group("product")
	{
		h.GET("/all", p.GetAllProducts)
	}
}

func (p *ProductController) GetAllProducts(c *gin.Context) {
	c.JSON(200, "Ask Talabi..")
}
