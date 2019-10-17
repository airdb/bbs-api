package handlers

import (
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// QueryTag godoc
// @Summary Query tag
// @Description Query tag
// @Tags tag
// @Accept  json
// @Produce  json
// @Param req body vo.ListArticleReq true "Record"
// @Success 200 {object} vo.ListArticleResp
// @Router /tag/getAll/:pid [get]
func QueryTag(c *gin.Context) {
	middlewares.SetResp(
		c,
		"list article",
	)
}
