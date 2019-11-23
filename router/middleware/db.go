package middleware

import (
	"github.com/chuxinplan/gin-mvc/common/db"
	"github.com/gin-gonic/gin"
)

func GetDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := db.DB.NewSession()
		defer session.Close()

		c.Set("db", session)
		c.Next()
	}
}
