package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func ValidateRequest(scheme interface{}) gin.HandlerFunc {
	v := validator.New()
	return func(c *gin.Context) {
		if err := c.Bind(&scheme); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := v.Struct(scheme); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.Set("reqBody", scheme) // set the validated reqBody to the context for later use
		c.Next()
	}
}
