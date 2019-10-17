package handlers

import (
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// QueryLog godoc
// @Summary Query log
// @Description Query log
// @Tags log
// @Accept  json
// @Produce  json
// @Param req body vo.ListArticleReq true "Record"
// @Success 200 {object} vo.ListArticleResp
// @Router /log/getAll/:pid [get]
func QueryLog(c *gin.Context) {
	middlewares.SetResp(
		c,
		"list article",
	)
}
