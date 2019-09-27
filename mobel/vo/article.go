package vo

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
