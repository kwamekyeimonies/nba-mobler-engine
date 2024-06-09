package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/nba-mobler-engine/config"
)

func GinRestAPIServer() error {
	serverURL := config.WithConfigInjector.KeystoreConfig.ServerAddress

	if len(serverURL) == 0 {
		return errors.New("unable to get server url/address")
	}

	ginServer := gin.Default()
	apiGroup := ginServer.Group("/api/v1.0")
	teamRouter.TeamRoutes(apiGroup)

	fmt.Println("server running on port: ", serverURL)

	err := ginServer.Run(serverURL)
	if err != nil {
		return err
	}

	return nil
}
