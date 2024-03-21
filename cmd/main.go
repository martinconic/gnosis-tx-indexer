package main

import (
	"fmt"
	"os"

	"github.com/martinconic/gnosis-tx-indexer/cmd/config"
	"github.com/martinconic/gnosis-tx-indexer/pkg/api"
)

func main() {

	v, err := config.NewViper()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	api.StartServer(v)

}
