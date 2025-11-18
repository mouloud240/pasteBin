package main

import (
	"pasteBin/internal/crons"
	"pasteBin/internal/database"
	. "pasteBin/internal/routes"
   swaggerfiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"

	docs "pasteBin/cmd/api/docs"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

// @title Pastebin API
// @version 1.0
// @description A simple pastebin service API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@pastebin.local

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey CookieAuth
// @in cookie
// @name auth

// setupRouter initializes the Gin router
// Returns:
//   *gin.Engine: The initialized Gin router
func setupRouter() *gin.Engine {
	r := gin.Default()
  

	return r
}

func main() {
	 docs.SwaggerInfo.BasePath = "/"
	addr:=":8080"
	db,err:=database.InitDB()
	if  err!=nil{
		panic("Failed to connect to database:" +err.Error())
	}

		  
	sqlDb,err:=db.DB()
	if err!=nil{
		panic("Failed to get database instance: "+err.Error())
	}
	defer sqlDb.Close()

	r := setupRouter()
	  r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	c:=cron.Cron{}

	SetupRoutes(r,db)
	crons.NewCronsManager(db,&c).InitCrons()

	// Don't trust any proxies for now until nginx is setup
	r.SetTrustedProxies(nil)
	// Listen and Server in 0.0.0.0:8080
	r.Run(addr);

}


