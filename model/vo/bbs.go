package vo

import "github.com/airdb/bbs-api/model/po"

func GetBBSArticles() (preForumPost []po.PreForumPost) {
	return po.GetBBSArticles()
}
