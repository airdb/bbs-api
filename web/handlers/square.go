package handlers

import (
	"fmt"

	"github.com/airdb/bbs-api/model/vo"
	"github.com/airdb/sailor/enum"
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// ListSquare godoc
// @Summary List square
// @Description List square
// @Tags square
// @Accept  json
// @Produce  json
// @Param req body vo.ListArticleReq true "Record"
// @Success 200 {object} vo.ListArticleResp
// @Router /square/list [get]
func ListSquare(c *gin.Context) {
	fmt.Print(enum.AirdbSuccess)
	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		vo.ListSquareResp{
			ID:         3,
			TagTypeID:  2,
			Sort:       1,
			Status:     1,
			CreateUser: 2,
			Flag:       1,
			Name:       "前端",
			Title:      "前端开发",
			Image:      "https://www.lyh.red/file/首页轮播_20190418155210_g6fk/20190418160520_8hee.png",
		},
	)
}

/*

id: 3,
tag_type_id: 2,
sort: 1,
status: 1,
create_user: 2,
create_time: "2019-05-29T05:47:42.000Z",
update_user: 2,
update_time: "2019-10-25T08:28:46.000Z",
delete_user: null,
delete_time: null,
flag: 1,
name: "前端",
image: "https://www.lyh.red/file/首页轮播_20190418155210_g6fk/20190418160520_8hee.png",
title: "前端开发"
*/

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
	fmt.Print(enum.AirdbSuccess)
	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		vo.ListSquareResp{
			ID:    3,
			Name:  "前端",
			Title: "前端开发",
			Image: "https://www.lyh.red/file/首页轮播_20190418155210_g6fk/20190418160520_8hee.png",
		},
	)
}
