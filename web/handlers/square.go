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
	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		vo.ListSquare(),
	)
}
