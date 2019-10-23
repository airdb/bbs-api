package handlers

import (
	"github.com/airdb/bbs-api/model/vo"
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
// @Router /square/query [get]
func QuerySquare(c *gin.Context) {
	middlewares.SetResp(
		c,
		vo.ListSquareResp{
			ID: 3,
			Name:"前端",
			Title: "前端开发",
			Image: "https://www.lyh.red/file/首页轮播_20190418155210_g6fk/20190418160520_8hee.png",
		},
	)
}
