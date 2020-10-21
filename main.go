package main

import (
	"git.qix.sx/hackathon/01-go/tank-rush/hackathon-2020/controllers"
	"git.qix.sx/hackathon/01-go/tank-rush/hackathon-2020/middleware"
	"github.com/gin-gonic/gin"
	redisCache "github.com/go-redis/cache/v8"
	"gorm.io/gorm"
)

var db *gorm.DB
var cache *redisCache.Cache

func main() {

	//Initializing all service things
	dotEnvInit()
	dbInit()

	//Initial csv-to-db import prompt
	initialSetup()

	cacheInit() //Late cache init, to prevent redis error on preparation input stage

	//Prepare JWT instance
	authMiddleware := middleware.AuthorizeJWT()
	r := gin.Default()

	//Get security token - make POST query {username: admin, password: pass}
	auth := r.Group("/api")

	auth.POST("/login", authMiddleware.LoginHandler)

	//Refresh token. simply go get query without any inputs
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)

	//If tokens are present and valid
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		//TODO: Add INSERT method TEST THIS!
		auth.POST("/put", func(c *gin.Context) { controllers.Put_info(c, db) })
		auth.GET("/v1/:product_id", func(c *gin.Context) { controllers.Get_info(c, db, cache) })
	}
	r.Run(":8081")
}
