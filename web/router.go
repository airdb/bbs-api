package web

import (
	"fmt"
	_ "github.com/airdb/bbs-api/docs"
	"github.com/airdb/bbs-api/web/handlers"
	"github.com/airdb/sailor/config"
	swaggerHandlers "github.com/airdb/sailor/gin/handlers"
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// @title BBS API
// @description This is the API server of BBS.
// @BasePath /apis/bbs/v1
func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(
		middlewares.Jsonifier(""),
	)

	// router.NoRoute(middlewares.Default404)

	v0API := router.Group("/apis/bbs/v0")
	swaggerHandlers.RegisterSwagger(v0API)
	RobotAPI := v0API.Group("/robot")
	RobotAPI.GET("query", handlers.QueryBBS)

	v1API := router.Group("/apis/bbs/v1")
	swaggerHandlers.RegisterSwagger(v1API)
	// Area APIs.
	AreaAPI := v1API.Group("/area")
	AreaAPI.GET("list", handlers.ListArea)
	AreaAPI.GET("query", handlers.QueryArea)
	AreaAPI.GET("update", handlers.UpdateArea)

	// Artciles APIs.
	ArticleAPI := v1API.Group("/article")
	ArticleAPI.GET("list", handlers.ListArticle)
	ArticleAPI.GET("query", handlers.QueryArticle)
	ArticleAPI.GET("create", handlers.QueryArticle)

	return router
}

func Run() {
	config.Init()
	fmt.Printf("Env: %s, Port: %s\n", config.GetEnv(), config.GetPort())
	err := NewRouter().Run("0.0.0.0:" + config.GetPort())
	if err != nil {
		fmt.Println("error: ", err)
	}
}
