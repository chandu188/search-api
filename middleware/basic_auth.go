package middleware

import "github.com/gin-gonic/gin"

func BasicAuth() gin.HandlerFunc {
	users := make(map[string]string)
	users["chandu"] = "password"
	return gin.BasicAuth(users)
}
