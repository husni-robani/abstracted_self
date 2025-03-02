package response

import "github.com/gin-gonic/gin"

type response struct {
	Message string `json:"message,omitempty"`
	Data any `json:"data,omitempty"`
	Errors any `json:"errors,omitempty"`
}

func Success(c *gin.Context, httpCode int, message string, data any) {
	payload := response{
		Message: message,
		Data: data,
	}
	c.JSON(httpCode, payload)
}

func Error(c *gin.Context, httpCode int, message string, errors any) {
	payload := response{
		Errors: errors,
	}

	c.JSON(httpCode, payload)
}