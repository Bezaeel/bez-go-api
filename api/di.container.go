package api

import (
	"bez-go-api/api/controllers"
	"bez-go-api/internal/app"

	"github.com/gin-gonic/gin"
)

type Container struct {
	NewProductController *gin.Engine
}

func NewContainer(handler *gin.Engine, _service *app.Container) {
	h := handler.Group("/")
	{
		controllers.NewProductController(h, _service)
	}
}
