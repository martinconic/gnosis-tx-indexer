package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martinconic/gnosis-tx-indexer/pkg/constants"
	"github.com/martinconic/gnosis-tx-indexer/pkg/indexer"
)

func GetTransactions(c *gin.Context) {
	address := c.Param("address")

	trans, _ := indexer.GetTransactionByAddress(address, service.viper.GetString(constants.ApiKey))

	service.store.Open(constants.FilePath + address)

	go service.store.Put(trans)

	c.JSON(http.StatusOK, gin.H{
		"Response": trans,
	})

}
