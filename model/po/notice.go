package po

import (
	"github.com/airdb/sailor/dbutils"
	"github.com/jinzhu/gorm"
)

type Notice struct {
	gorm.Model
	ID      uint
	Href    string
	Content string
	Name    string `json:"name"`
	Flag    uint
	Status  uint
}

func ListNotice() (notices []*Notice) {
	dbutils.DefaultDB().Table("notice_tab").Find(&notices)
	return
}
