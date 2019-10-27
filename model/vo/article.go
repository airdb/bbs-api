package vo

import (
	"github.com/airdb/bbs-api/model/po"
)

type ListArticleReq struct {
}

type ListArticleResp struct {
	// Article []*string `json:"Article"`
	Result   []*Article
	CurPage  uint `json:"curPage"`
	PageSize uint `json:"pageSize"`
	Totals   uint `json:"totals"`
}

type Article struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Flag uint `json:"flag"`
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

func QueryBBSByKeyword(keyword string) []*po.MinaArticle {
	return po.QueryBBSByKeywords(keyword)
}

func FromPoArticle(pArticle *po.Article) *Article {
	return &Article{
		ID:      pArticle.ID,
		Content: pArticle.Content,
		Flag: pArticle.Flag,
	}

}

func FromPoArticles(pArticles []*po.Article) (articles []*Article) {
	for _, pArticle := range pArticles {
		articles = append(articles, FromPoArticle(pArticle))
	}
	return
}

func ListArticle() *ListArticleResp {
	resp := ListArticleResp{
		Result:   FromPoArticles(po.ListArticle()),
		CurPage:  1,
		PageSize: 10,
		Totals:   10,
	}

	return &resp
}
