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

func Run() {
	config.Init()
	fmt.Printf("Env: %s, Port: %s\n", config.GetEnv(), config.GetPort())
	err := NewRouter().Run("0.0.0.0:" + config.GetPort())
	if err != nil {
		fmt.Println("error: ", err)
	}
}

// @title BBS API
// @description This is the API server of BBS.
// @BasePath /apis/bbs/v1
func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// router.NoRoute(middlewares.Default404)

	v0API := router.Group("/apis/bbs/v0")
	swaggerHandlers.RegisterSwagger(v0API)
	RobotAPI := v0API.Group("/robot")
	RobotAPI.GET("query", handlers.QueryBBS)

	v1API := router.Group("/apis/bbs/v1")
	v1API.Use(
		middlewares.Jsonifier(),
	)
	swaggerHandlers.RegisterSwagger(v1API)

	// Area APIs.
	area(v1API)

	homeGroup(v1API)

	charts(v1API)
	column(v1API)
	count(v1API)
	dataPerms(v1API)
	draft(v1API)
	file(v1API)
	folder(v1API)
	log(v1API)
	menu(v1API)
	question(v1API)
	role(v1API)
	role_relation(v1API)
	tag(v1API)
	tagRelation(v1API)
	topic(v1API)
	user(v1API)

	return router
}

func area(api *gin.RouterGroup) {
	r := api.Group("/area")
	r.GET("list", handlers.ListArea)
	r.GET("query", handlers.QueryArea)
	r.GET("update", handlers.UpdateArea)

}

func homeGroup(api *gin.RouterGroup) {
	// Artciles APIs.
	artilceAPI := api.Group("/article")
	artilceAPI.GET("list", handlers.ListArticle)
	artilceAPI.GET("query", handlers.QueryArticle)
	artilceAPI.GET("create", handlers.QueryArticle)

	squareAPI := api.Group("/square")
	squareAPI.GET("list", handlers.ListSquare)

	noticeAPI := api.Group("/notice")
	noticeAPI.GET("list", handlers.ListNotice)

	// Carousels APIs.
	carouselAPI := api.Group("/carousel")
	carouselAPI.GET("list", handlers.ListCarousel)
}

func user(api *gin.RouterGroup) {
	userAPI := api.Group("/user")
	userAPI.POST("login", handlers.LoginUser)
	userAPI.GET("query", handlers.QueryUser)
	userAPI.OPTIONS("login", handlers.LoginUser)
}

func charts(api *gin.RouterGroup) {
	r := api.Group("/charts")
	r.GET("query", handlers.QueryCharts)
}

func column(api *gin.RouterGroup) {
	r := api.Group("/column")
	r.GET("query", handlers.QueryColumn)
}

func count(api *gin.RouterGroup) {
	r := api.Group("/count")
	r.GET("query", handlers.QueryCount)
}

func dataPerms(api *gin.RouterGroup) {
	r := api.Group("/data_perms")
	r.GET("query", handlers.QueryDataPerms)
}

func draft(api *gin.RouterGroup) {
	r := api.Group("/draft")
	r.GET("query", handlers.QueryDraft)
}
func file(api *gin.RouterGroup) {
	r := api.Group("/file")
	r.GET("query", handlers.QueryFile)
}
func folder(api *gin.RouterGroup) {
	r := api.Group("/folder")
	r.GET("query", handlers.QueryFolder)
}

func log(api *gin.RouterGroup) {
	r := api.Group("/log")
	r.GET("query", handlers.QueryLog)
}

func menu(api *gin.RouterGroup) {
	r := api.Group("/menu")
	r.GET("query", handlers.QueryMenu)
}

func question(api *gin.RouterGroup) {
	r := api.Group("/question")
	r.GET("query", handlers.QueryQuestion)
}
func role(api *gin.RouterGroup) {
	r := api.Group("/role")
	r.GET("query", handlers.QueryRole)
}
func role_relation(api *gin.RouterGroup) {
	r := api.Group("/role_relation")
	r.GET("query", handlers.QueryRoleRelation)
}

func tag(api *gin.RouterGroup) {
	r := api.Group("/tag")
	r.GET("query", handlers.QueryTag)
}

func tagRelation(api *gin.RouterGroup) {
	r := api.Group("/tag_relation")
	r.GET("query", handlers.QueryTagRelation)
}

func topic(api *gin.RouterGroup) {
	r := api.Group("/topic")
	r.GET("query", handlers.QueryTopic)
}
