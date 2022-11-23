package middleware

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Content-Type, Content-Length, Accept-Encoding, accept, origin, Cache-Control, X-Requested-With, token
		//content-type,Version,Device,Unique-Device,Drives,Web-Token
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Version,Device,Unique-Device,Drives,Web-Token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
	}
}
