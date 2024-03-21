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

	store storage.Store
}

var service *Service

func StartServer(v *viper.Viper) {
	service = &Service{}
	service.Initialize(v)
	service.Run(v.GetString(constants.ApiServer))
}

func (server *Service) Initialize(v *viper.Viper) {
	var err error
	server.router = gin.Default()

	server.store, _ = filestore.NewFileStore("test")
	if err != nil {
		log.Println(err)
	}

	service.InitializeRoutes()
}

func (service *Service) Run(addr string) {
	err := service.router.Run(":" + addr)
	if err != nil {
		panic("unable to start server!")
	}
}
