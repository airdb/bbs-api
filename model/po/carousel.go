package po

import (
	"github.com/airdb/sailor/dbutils"
	"github.com/jinzhu/gorm"
)

type Carousel struct {
	gorm.Model
	Image  string
	Click  uint
	Href   string
	Title  string
	Flag   uint
	Status uint
}

func ListCarousel() (carousels []*Carousel) {
	dbutils.DefaultDB().Table("carousel_tab").Debug().Find(&carousels)
	return
}
