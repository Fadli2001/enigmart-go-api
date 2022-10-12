package delivery

import (
	"enigmart-api/config"
	"enigmart-api/delivery/controller"
	"enigmart-api/manager"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func Server() *appServer {
	ginEngine := gin.Default()
	config := config.NewConfig()
	infra := manager.NewInfraManager(config)
	repo := manager.NewRepositoryManager(infra)
	usecase := manager.NewUseCaseManager(repo)
	return &appServer{
		useCaseManager: usecase,
		engine:         ginEngine,
		host:           config.Url,
	}
}

func (a *appServer) initHandlers() {
	controller.NewProductController(a.engine, a.useCaseManager.ProductUseCase())
}

func (a *appServer) Run() {
	a.initHandlers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
