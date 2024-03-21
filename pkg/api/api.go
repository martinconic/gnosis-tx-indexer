package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/martinconic/gnosis-tx-indexer/pkg/constants"
	"github.com/martinconic/gnosis-tx-indexer/pkg/storage"
	"github.com/martinconic/gnosis-tx-indexer/pkg/storage/filestore"
	"github.com/spf13/viper"
)

type Service struct {
	router *gin.Engine
	viper  *viper.Viper
	store  storage.Store
}

var service *Service

func StartServer(v *viper.Viper) {
	service = &Service{viper: v}
	service.Initialize()
	service.Run()
}

func (server *Service) Initialize() {
	var err error
	server.router = gin.Default()

	server.store, _ = filestore.NewFileStore()
	if err != nil {
		log.Println(err)
	}

	service.InitializeRoutes()
}

func (service *Service) Run() {
	err := service.router.Run(":" + service.viper.GetString(constants.ApiServer))
	if err != nil {
		panic("unable to start server!")
	}
}
