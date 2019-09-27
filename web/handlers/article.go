package handlers

import (
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// ListArticle godoc
// @Summary List Articles
// @Description List Articles
// @Tags article
// @Accept  json
// @Produce  json
// @Param req body vo.ListArticleReq true "Record"
// @Success 200 {object} vo.ListArticleResp
// @Router /article/list [get]
func ListArticle(c *gin.Context) {
	middlewares.SetResp(
		c,
		"list article",
	)
}

// QueryAritcle godoc
// @Summary Query article
// @Description Query article
// @Tags article
// @Accept  json
// @Produce  json
// @Param req body vo.QueryArticleReq true "Record"
// @Success 200 {object} vo.QueryArticleResp
// @Router /article/list [get]
func QueryArticle(c *gin.Context) {
	middlewares.SetResp(
		c,
		"query article",
	)
}
