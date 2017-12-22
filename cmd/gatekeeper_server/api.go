package main

import "github.com/gin-gonic/gin"

func initAPI() (*gin.Engine, error) {
	router := gin.Default()

	// Simple group: v1
	/*v1 := */
	router.Group("/v1")
	{
		/*v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)*/
	}

	return router, nil
}
