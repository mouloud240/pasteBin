package main

import (
	"pasteBin/internal/initializers"
	. "pasteBin/internal/routes"

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
	err:=initializers.InitDb()
	if err!=nil{
		panic("Failed to connect to database: "+err.Error())
	}

	r := setupRouter()

	SetupRoutes(r)

	
	// Don't trues any proxies for now until nginx is setup
	r.SetTrustedProxies(nil)
	// Listen and Server in 0.0.0.0:8080
	r.Run(addr);

}


