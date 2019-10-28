package vo

import "github.com/airdb/bbs-api/model/po"

type ListNoticeReq struct {
}

type ListNoticeResp struct {
	Square []Square
}

type Notice struct {
	ID uint `json:"id"`
	Href string `json:"href"`
	Content string `json:"content"`
	Flag uint `json:"flag"`
	Status uint `json:"status"`
}

func FromPoNotice(pNotice *po.Notice) *Notice{
	return &Notice {
		ID: pNotice.ID,
		Href:pNotice.Href,
		Content:pNotice.Content,
		Flag:pNotice.Flag,
		Status:pNotice.Status,
	}
}

func FromPoNotices(pNotices []*po.Notice) (vNotices []*Notice) {
	for _, notice := range pNotices{
		vNotices = append(vNotices, FromPoNotice(notice))
	}

	return
}

func ListNotice() []*Notice{
	return FromPoNotices(po.ListNotice())
}
