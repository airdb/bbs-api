package handlers

import (
	"github.com/airdb/sailor/enum"
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// QueryFile godoc
// @Summary Query file
// @Description Query file
// @Tags file
// @Accept  json
// @Produce  json
// @Param req body vo.ListArticleReq true "Record"
// @Success 200 {object} vo.ListArticleResp
// @Router /file/query/:pid [get]
func QueryFile(c *gin.Context) {
	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		"list article",
	)
}
