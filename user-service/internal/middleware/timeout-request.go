package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	. "github.com/hadanhtuan/go-sdk/common"
)

func TimeoutMiddleware(perSecond int) gin.HandlerFunc {
	return func(c *gin.Context) {
		res := APIResponse{Status: APIStatus.Timeout}

		convertToSecond := time.Duration(perSecond) * time.Second

		ctx, cancel := context.WithTimeout(c.Request.Context(), convertToSecond)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		done := make(chan bool, 1)
		go func() {
			c.Next()
			done <- true
		}()

		select {
		case <-done:
			return
		case <-ctx.Done():
			res.Message = "Request timeout"
			c.AbortWithStatusJSON(int(APIStatus.Timeout), res)
		}
	}
}
