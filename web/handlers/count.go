package handlers

import (
	"github.com/airdb/sailor/enum"
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// QueryCount godoc
// @Summary Query count
// @Description Query count
// @Tags count
// @Accept  json
// @Produce  json
// @Param req body vo.ListArticleReq true "Record"
// @Success 200 {object} vo.ListArticleResp
// @Router /count/query/:pid [get]
func QueryCount(c *gin.Context) {
	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		"list article",
	)
}
