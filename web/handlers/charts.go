package handlers

import (
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// QueryCarousel godoc
// @Summary Query area
// @Description Query area
// @Tags area
// @Accept  json
// @Produce  json
// @Param req body vo.ListArticleReq true "Record"
// @Success 200 {object} vo.ListArticleResp
// @Router /area/getAll/:pid [get]
func QueryCharts(c *gin.Context) {
	middlewares.SetResp(
		c,
		"list article",
	)
}
