package handlers

import (
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// QuerySquare godoc
// @Summary Query square
// @Description Query square
// @Tags square
// @Accept  json
// @Produce  json
// @Param req body vo.ListArticleReq true "Record"
// @Success 200 {object} vo.ListArticleResp
// @Router /square/getAll/:pid [get]
func QuerySquare(c *gin.Context) {
	middlewares.SetResp(
		c,
		"list article",
	)
}
