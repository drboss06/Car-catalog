package handler

import (
	"carDirectory/logger"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

var l = logger.Get()

// newErrorResponse logs an error message and aborts the request with a specified status code and error message.
//
// c *gin.Context, statusCode int, message string.
// None.
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	l.Error().Msg(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
