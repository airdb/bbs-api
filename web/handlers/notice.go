package handlers

import (
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// QueryNotice godoc
// @Summary Query notice
// @Description Query notice
// @Tags notice
// @Accept  json
// @Produce  json
// @Param req body vo.ListArticleReq true "Record"
// @Success 200 {object} vo.ListArticleResp
// @Router /notice/getAll/:pid [get]
func QueryNotice(c *gin.Context) {
	middlewares.SetResp(
		c,
		"list article",
	)
}
