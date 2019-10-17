package handlers

import (
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// QueryTopic godoc
// @Summary Query topic
// @Description Query topic
// @Tags topic
// @Accept  json
// @Produce  json
// @Param req body vo.ListArticleReq true "Record"
// @Success 200 {object} vo.ListArticleResp
// @Router /topic/getAll/:pid [get]
func QueryTopic(c *gin.Context) {
	middlewares.SetResp(
		c,
		"list article",
	)
}
