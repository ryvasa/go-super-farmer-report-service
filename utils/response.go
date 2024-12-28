package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func NewResponse(status int, success bool, message string, data interface{}, errors interface{}) Response {
	return Response{
		Status:  status,
		Success: success,
		Message: message,
		Data:    data,
		Errors:  errors,
	}
}

func ErrorResponse(c *gin.Context, err error) {
	statusCode := GetStatusCode(err)
	response := NewResponse(statusCode, false, "failed", nil, err)
	c.JSON(statusCode, response)
}

func SuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	response := NewResponse(statusCode, true, "success", data, nil)
	c.JSON(statusCode, response)
}
