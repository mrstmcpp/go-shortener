package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimit() gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(5), 10)

	return func(ctx *gin.Context) {
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests",
			})
			return
		}
		ctx.Next()
	}

}
