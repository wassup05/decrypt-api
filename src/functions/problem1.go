package problems

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func factorial(num int) int {
	if num == 1 {
		return num
	}

	return num * factorial(num-1)
}

func Problem1(ctx *gin.Context) {

	type req struct {
		Input int `json:"input"`
	}

	var json req

	err := ctx.BindJSON(&json)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"output": factorial(json.Input)})
}
