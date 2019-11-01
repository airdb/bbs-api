package po

import (
	"github.com/airdb/sailor/dbutils"
	"github.com/jinzhu/gorm"
)

type Square struct {
	gorm.Model
	Name string `json:"name"`
}

func ListSquare() (squares []*Square) {
	dbutils.DefaultDB().Table("square_tab").Debug().Find(&squares)
	return
}
