package main

import (
	"pasteBin/internal/crons"
	"pasteBin/internal/database"
	"pasteBin/internal/initializers"
	. "pasteBin/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
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
	db,err:=database.InitDB()
	if  err!=nil{
		panic("Failed to connect to database:" +err.Error())
	}

		if err:=initializers.InitEnv(".env");err!=nil{
		panic("Failed to load env variables: "+err.Error())
	}
  


	r := setupRouter()
	c:=cron.Cron{}

	SetupRoutes(r,db)
	crons.NewCronsManager(db,&c).InitCrons()

	sqlDb,err:=db.DB()
	if err!=nil{
		panic("Failed to get database instance: "+err.Error())
	}
	defer sqlDb.Close()

	// Don't trues any proxies for now until nginx is setup
	r.SetTrustedProxies(nil)
	// Listen and Server in 0.0.0.0:8080
	r.Run(addr);
	if err!=nil{
		panic("Failed to get database instance: "+err.Error())
	}

}


