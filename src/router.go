package main

import (
	"decrypt-api/src/functions"
	"github.com/gin-gonic/gin"
)

func init_router(eng *gin.Engine) {
	problemGroup := eng.Group("/api")

	problemGroup.GET("problem1", problems.Problem1)
}
