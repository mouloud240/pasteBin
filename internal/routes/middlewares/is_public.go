package middlewares

import "github.com/gin-gonic/gin"
func PublicSetter() gin.HandlerFunc {
	  
	return SetMiddleware("isPublic", true);
}
func PublicPrivate() gin.HandlerFunc {
	
	return  SetMiddleware("isPublicPrivate",true)
	
}
