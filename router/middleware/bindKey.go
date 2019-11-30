package middleware

import (
	"github.com/chuxinplan/gin-mvc/common/db"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func BindKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		// bind db
		session := db.DB.NewSession()
		defer session.Close()
		c.Set("db", session)

		// bind request id
		requestId := c.Request.Header.Get("X-Request-Id")
		if requestId == "" {
			requestId = uuid.NewV4().String()
		}
		c.Set("X-Request-Id", requestId)
		c.Header("X-Request-Id", requestId)

		c.Next()
	}
}
