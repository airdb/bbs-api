package vo

import "github.com/airdb/bbs-api/model/po"

type Square struct {
	ID         uint   `json:"id"`
	TagTypeID  uint   `json:"tag_type_id"`
	Status     int    `json:"status"`
	Sort       uint   `json:"sort"`
	CreateUser uint   `json:"create_user"`
	Title      string `json:"Title"`
	Name       string `json:"name"`
	Image      string `json:"image"`
	Flag       uint   `json:"flag"`
}

type ListSquareReq struct {
}

type ListSquareResp struct {
	Square []Square
}

func FromPoSquare() {

}

func ToPoSquare() {

}

func ListSquare() []*Square {
	po.ListSquare()
	var square []*Square

	s := Square{
		ID:         3,
		TagTypeID:  2,
		Sort:       1,
		Status:     1,
		CreateUser: 2,
		Flag:       1,
		Name:       "前端",
		Title:      "前端开发",
		Image:      "https://www.lyh.red/file/首页轮播_20190418155210_g6fk/20190418160520_8hee.png",
	}

	square = append(square, &s)

	return square
}