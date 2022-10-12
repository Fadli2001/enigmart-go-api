package delivery

import (
	"github.com/gin-gonic/gin"
	"enigmart-api/usecase"
	"enigmart-api/repository"
	"enigmart-api/delivery/controller"
)

type appServer struct {
	prodUseCase usecase.ProductUseCase
	engine      *gin.Engine
	host        string
}

func Server() *appServer {
	ginEngine := gin.Default()
	prodRepo := repository.NewProductRepository()
	prodUseCase := usecase.NewProductUseCase(prodRepo)
	host := ":8181"
	return &appServer{
		prodUseCase: prodUseCase,
		engine:      ginEngine,
		host:        host,
	}
}

func (a *appServer) initHandlers() {
	controller.NewProductController(a.engine, a.prodUseCase)
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
