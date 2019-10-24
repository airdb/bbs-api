package handlers

import (
	"github.com/airdb/sailor/enum"
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// QueryColumn godoc
// @Summary Query area
// @Description Query area
// @Tags column
// @Accept  json
// @Produce  json
// @Param req body vo.ListArticleReq true "Record"
// @Success 200 {object} vo.ListArticleResp
// @Router /column/query/:pid [get]
func QueryColumn(c *gin.Context) {
	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		"list article",
	)
}
