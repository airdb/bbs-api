package handlers

import (
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
// @Param req body vo.ListSquareReq true "Record"
// @Success 200 {object} vo.ListSquareResp
// @Router /square/list [get]
func ListSquare(c *gin.Context) {
	/*
		square := vo.Square{
			ID:         3,
			TagTypeID:  2,
			Sort:       1,
			Status:     1,
			CreateUser: 2,
			Flag:       1,
			Name:       "前端",
			Title:      "前端开发",
			Image:      "https://www.lyh.red/file/首页轮播_20190418155210_g6fk/20190418160520_8hee.png",
		}

		var resp vo.ListSquareResp
		resp.Square = append(resp.Square, square)
	*/

	aa := vo.ListSquare()

	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		&aa,
	)
}
