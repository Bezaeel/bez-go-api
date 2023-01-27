package bootstrap

import (
	"bez-go-api/api"
	"bez-go-api/db"
	"bez-go-api/internal/app"
	"bez-go-api/internal/core"
	"bez-go-api/pkg/httpserver"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Run() {

	// init DB
	db, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	repo := core.NewContainer(db)
	service := app.NewContainer(repo)

	handler := gin.New()
	api.NewContainer(handler, service)
	httpServer := httpserver.New(handler, httpserver.Port("8080"))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Fatal("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Fatal(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	fmt.Println(service.ProductService.GetProducts())
}
