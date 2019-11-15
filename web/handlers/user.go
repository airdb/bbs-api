package handlers

import (
	"github.com/airdb/bbs-api/model/vo"
	"github.com/airdb/sailor/enum"
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// QueryUser godoc
// @Summary Query user
// @Description Query user
// @Tags user
// @Accept  json
// @Produce  json
// @Param req body vo.ListArticleReq true "Record"
// @Success 200 {object} vo.ListArticleResp
// @Router /user/getAll/:pid [get]
func QueryUser(c *gin.Context) {
	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		"list article",
	)
}

// UserLogin godoc
// @Summary User login
// @Description User login
// @Tags user
// @Accept  json
// @Produce  json
// @Param req body vo.LoginReq true "user"
// @Success 200 {object} vo.LoginResp
// @Router /user/login [post]
func LoginUser(c *gin.Context) {
	var req vo.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.String(200, "OK")
		return
	}

	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		vo.Login(),
	)
}
