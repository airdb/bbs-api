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
	homeGroup(v1API)

	apiGroup2(v1API)

	return router
}

func homeGroup(api *gin.RouterGroup) {
	// Artciles APIs.
	artilceAPI := api.Group("/article")
	artilceAPI.GET("list", handlers.ListArticle)
	artilceAPI.GET("query", handlers.QueryArticle)
	artilceAPI.GET("create", handlers.QueryArticle)

	areaAPI := api.Group("/area")
	areaAPI.GET("list", handlers.ListArea)
	areaAPI.GET("query", handlers.QueryArea)
	areaAPI.GET("update", handlers.UpdateArea)

	squareAPI := api.Group("/square")
	squareAPI.GET("list", handlers.ListSquare)

	noticeAPI := api.Group("/notice")
	noticeAPI.GET("list", handlers.ListNotice)

	// Carousels APIs.
	carouselAPI := api.Group("/carousel")
	carouselAPI.GET("list", handlers.ListCarousel)
}

func apiGroup2(api *gin.RouterGroup) {
	userAPI := api.Group("/user")
	userAPI.POST("login", handlers.LoginUser)
	userAPI.GET("query", handlers.QueryUser)
	userAPI.OPTIONS("login", handlers.LoginUser)

	chartAPI := api.Group("/charts")
	chartAPI.GET("query", handlers.QueryCharts)

	columnAPI := api.Group("/column")
	columnAPI.GET("query", handlers.QueryColumn)

	countAPI := api.Group("/count")
	countAPI.GET("query", handlers.QueryCount)

	dataPermsAPI := api.Group("/data_perms")
	dataPermsAPI.GET("query", handlers.QueryDataPerms)

	draftAPI := api.Group("/draft")
	draftAPI.GET("query", handlers.QueryDraft)

	fileAPI := api.Group("/file")
	fileAPI.GET("query", handlers.QueryFile)

	folderAPI := api.Group("/folder")
	folderAPI.GET("query", handlers.QueryFolder)

	logAPI := api.Group("/log")
	logAPI.GET("query", handlers.QueryLog)

	menuAPI := api.Group("/menu")
	menuAPI.GET("query", handlers.QueryMenu)

	questionAPI := api.Group("/question")
	questionAPI.GET("query", handlers.QueryQuestion)

	roleAPI := api.Group("/role")
	roleAPI.GET("query", handlers.QueryRole)

	roleRelationAPI := api.Group("/role_relation")
	roleRelationAPI.GET("query", handlers.QueryRoleRelation)

	tagAPI := api.Group("/tag")
	tagAPI.GET("query", handlers.QueryTag)

	tarRelationAPI := api.Group("/tag_relation")
	tarRelationAPI.GET("query", handlers.QueryTagRelation)

	topicAPI := api.Group("/topic")
	topicAPI.GET("query", handlers.QueryTopic)
}
