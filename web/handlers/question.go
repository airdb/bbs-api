package handlers

import (
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// QueryQuestion godoc
// @Summary Query question
// @Description Query question
// @Tags question
// @Accept  json
// @Produce  json
// @Param req body vo.ListArticleReq true "Record"
// @Success 200 {object} vo.ListArticleResp
// @Router /question/getAll/:pid [get]
func QueryQuestion(c *gin.Context) {
	middlewares.SetResp(
		c,
		"list article",
	)
}
