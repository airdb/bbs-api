package handlers

import (
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

// QueryAritcle godoc
// @Summary Query article
// @Description Query article
// @Tags article
// @Accept  json
// @Produce  json
// @Param req body vo.QueryArticleReq true "Record"
// @Success 200 {object} vo.QueryArticleResp
// @Router /article/list [get]
func QQRobotQuery(c *gin.Context) {
	// qq=137780017&group=311401898&command=bbhj&message=
	// var rmessage models.QQRobotReciveMessage
	// rmessage.QQ = u.GetString("qq")
	// rmessage.Groupid= u.GetString("group")
	// rmessage.Command = u.GetString("command")
	// rmessage.Message = u.GetString("message")

	/*
	keyword := u.GetString("message")
	logs.Info("the reqeust keyword message: ",  keyword)
	msg := ""
	if ("" != keyword) {
		for _, info := range models.GetArticleByKeyword(keyword) {
			msg += info.Subject +"\n" + info.DataFrom  + "\n"
		}
	}
	if "" == msg {
		if ("" != keyword) {
			msg += "您的查询的信息，暂时无结果，可能是后台同步论坛数据失败。\n"
		}

		msg += "bbhj 机器人使用帮助\n"
		msg += "示例1：bbhj 4407\n"
		msg += "示例2：bbhj 山西 4407\n"
		msg += "示例3：bbhj 山西 运城 张\n"
		msg += "\n"
		msg += "说明：bbhj命令支持最多3个关键字的查询; 命令及各关键字只能以空格分隔。"
	}
	msg += "(出处: 宝贝回家论坛)\n"

	u.Ctx.WriteString(msg)
	*/
	middlewares.SetResp(
		c,
		"query article",
	)
}
