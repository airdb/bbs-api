package handlers

import (
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// QueryPerm godoc
// @Summary Query area
// @Description Query area
// @Tags perm
// @Accept  json
// @Produce  json
// @Param req body vo.ListArticleReq true "Record"
// @Success 200 {object} vo.ListArticleResp
// @Router /perm/query/:pid [get]
func QueryDataPerms(c *gin.Context) {
	middlewares.SetResp(
		c,
		"list article",
	)
}
