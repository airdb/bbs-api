package handlers

import (
	"github.com/airdb/bbs-api/mobel/vo"
	"github.com/gin-gonic/gin"
)

// QueryBBS godoc
// @Summary for QQ robot query article
// @Description Query article
// @Tags article
// @Accept  json
// @Produce  json
// @Param req body vo.QueryBBSReq true "Message"
// @Success 200 {object} vo.QueryBBSResp
// @Router /article/list [get]
func QueryBBS(c *gin.Context) {
	var req vo.QueryBBSReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.String(400, "不合法的请求参数")
		return
	}

	var msg string
	if req.Message == "" {
		// help
		msg += "bbhj 机器人使用帮助\n"
		msg += "示例1：bbhj 4407\n"
		msg += "示例2：bbhj 山西 4407\n"
		msg += "示例3：bbhj 山西 运城 张\n"
		msg += "\n"
		msg += "说明：bbhj命令支持最多3个关键字的查询; 命令及各关键字只能以空格分隔。"
	} else {
		articles := vo.QueryBBSByKeyword(req.Message)
		if len(articles) == 0 {
			msg += "您的查询的信息，暂时无结果，可能是后台同步论坛数据失败。\n"
		} else {
			for _, article := range articles {
				msg += article.Subject + "\n" + article.DataFrom + "\n"
			}
		}
	}

	msg += "(出处: 宝贝回家论坛)\n"
	c.String(200, msg)
}
