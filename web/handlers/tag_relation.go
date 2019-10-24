package handlers

import (
	"github.com/airdb/sailor/enum"
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// QueryTagRelation godoc
// @Summary Query tag_relation
// @Description Query tag_relation
// @Tags tag_relation
// @Accept  json
// @Produce  json
// @Param req body vo.ListArticleReq true "Record"
// @Success 200 {object} vo.ListArticleResp
// @Router /tag_relation/getAll/:pid [get]
func QueryTagRelation(c *gin.Context) {
	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		"list article",
	)
}
