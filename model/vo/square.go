package vo

type ListSquareReq struct {
}

type ListSquareResp struct {
	ID         uint   `json:"id"`
	TagTypeID  uint   `json:"tag_type_id"`
	Status     int    `json:"status"`
	Sort       uint   `json:"sort"`
	CreateUser uint   `json:"create_user"`
	Title      string `json:"Title"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	Flag       bool   `json:"flag"`
}
