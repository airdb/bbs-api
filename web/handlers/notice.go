package handlers

import (
	"github.com/airdb/bbs-api/model/vo"
	"github.com/airdb/sailor/enum"
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// ListNotice godoc
// @Summary List notice
// @Description List notice
// @Tags notice
// @Accept  json
// @Produce  json
// @Param req body vo.ListNoticeReq true "Record"
// @Success 200 {object} vo.ListNoticeResp
// @Router /notice/list [get]
func ListNotice(c *gin.Context) {
	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		vo.ListNotice(),
	)
}
