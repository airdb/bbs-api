package handlers

import (
	"github.com/airdb/bbs-api/model/vo"
	"github.com/airdb/sailor/enum"
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

/*
// QueryCarousel godoc
// @Summary Query carousel
// @Description Query carousel
// @Tags carousel
// @Accept  json
// @Produce  json
// @Param req body vo.ListCarouselReq true "Record"
// @Success 200 {object} vo.ListCarouselResp
// @Router /carousel/list [get]

*/
func ListCarousel(c *gin.Context) {
	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		vo.ListCarousel(),
	)
}
