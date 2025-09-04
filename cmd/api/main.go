package main

import (

	"github.com/gin-gonic/gin"
)
// setupRouter initializes the Gin router
// Returs:
//   *gin.Engine: The initialized Gin router
func setupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func main() {
	addr:=":8080"
	r := setupRouter()
	// Don't trues any proxies for now until nginx is setup
	r.SetTrustedProxies(nil)
	// Listen and Server in 0.0.0.0:8080
	r.Run(addr);}

