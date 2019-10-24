package handlers

import (
	"github.com/airdb/sailor/enum"
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// QueryRole godoc
// @Summary Query role
// @Description Query role
// @Tags role
// @Accept  json
// @Produce  json
// @Param req body vo.ListArticleReq true "Record"
// @Success 200 {object} vo.ListArticleResp
// @Router /role/getAll/:pid [get]
func QueryRole(c *gin.Context) {
	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		"list article",
	)
}
