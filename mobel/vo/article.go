package vo

import "github.com/airdb/bbs-api/mobel/po"

type ListArticleReq struct {
}

type ListArticleResp struct {
	Article []*string `json:"Article"`
}

type QueryArticleReq struct {
}

type QueryArticleResp struct {
	Article []*string `json:"Article"`
}

type QQRobotQueryReq struct {
	QQ      uint   `form:"qq"`
	Group   uint   `form:"group"`
	Command string `form:"command"`
	Message string `form:"message"`
}

type QQRobotQueryResp string

func QueryBBSByKeyword(keyword string) []*po.Article {
	return po.QueryBBSByKeywords(keyword)
}
