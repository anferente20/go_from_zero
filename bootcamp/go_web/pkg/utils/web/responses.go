package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Data interface{} `json:"data"`
}

func SetErrorResponse(errorResponse ErrorResponse, ctx *gin.Context) {
	ctx.IndentedJSON(errorResponse.Status, errorResponse)
}

func SetResponse(response Response, ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, response)
}
