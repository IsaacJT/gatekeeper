package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/isaacjt/gatekeeper-server/api/v1"
)

func initV1API(router *gin.Engine) {
	apiV1 := router.Group("/v1")
	{
		apiV1.GET("/info", v1.InfoEndpoint)
	}
}

func initAPI() (*gin.Engine, error) {
	router := gin.Default()

	if router == nil {
		return nil, fmt.Errorf("could not create router")
	}

	initV1API(router)

	return router, nil
}
