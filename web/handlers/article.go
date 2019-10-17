package handlers

import (
	"github.com/airdb/bbs-api/model/vo"
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// ListArticle godoc
// @Summary List Articles
// @Description List Articles
// @Tags article
// @Accept  json
// @Produce  json
// @Param req body vo.ListArticleReq true "Record"
// @Success 200 {object} vo.ListArticleResp
// @Router /article/list [get]
func ListArticle(c *gin.Context) {
	middlewares.SetResp(
		c,
		"list article",
	)
}

// QueryAritcle godoc
// @Summary Query article
// @Description Query article
// @Tags article
// @Accept  json
// @Produce  json
// @Param req body vo.QueryArticleReq true "Record"
// @Success 200 {object} vo.QueryArticleResp
// @Router /article/list [get]
func QueryArticle(c *gin.Context) {
	middlewares.SetResp(
		c,
		"query article",
	)
}

// QueryBBS godoc
// @Summary for QQ robot query article
// @Description Query article
// @Tags article
// @Accept  json
// @Produce  json
// @Param req body vo.QQRobotQueryReq true "Message"
// @Success 200 {string} string vo.QQRobotQueryResp
// @Router /robot/query [get]
func QueryBBS(c *gin.Context) {
	var req vo.QQRobotQueryReq
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
