package vo

import (
	"fmt"

	"github.com/airdb/bbs-api/model/po"
)

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

func FromPoSquare(square *po.Square) *Square {
	return &Square{
		ID:   square.ID,
		Name: square.Name,
	}
}

func FromPoSquares(pSquares []*po.Square) (vSquares []*Square) {
	for _, square := range pSquares {
		vSquares = append(vSquares, FromPoSquare(square))
	}

	return
}

func ListSquare() []*Square {
	aa := FromPoSquares(po.ListSquare())
	for _, a := range aa {
		fmt.Println("xxxx", a.Name)
	}
	return aa
}
