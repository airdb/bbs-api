package po

import (
	"github.com/airdb/sailor/dbutils"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

var DBName = "dev_mina"

type Article struct {
	gorm.Model
	UUID            string
	AvatarUrl       string
	Nickname        string
	Gender          uint
	Title           string
	Subject         string
	Characters      string
	Details         string
	DataFrom        string
	BirthedProvince string
	BirthedCity     string
	BirthedCountry  string
	BirthedAddress  string
	BirthedAt       time.Time `gorm:"type:datetime"`

	MissedCountry  string
	MissedProvince string
	MissedCity     string
	MissedAddress  string
	MissedAt       time.Time `gorm:"column:missed_at;type:datetime"`
	Handler        string
	Babyid         string
	Category       string
	Height         string
	SyncStatus     int `gorm:"column:syncstatus;default:0"`
}

func QueryBBSByKeywords(keyword string) (articles []*Article) {
	keys := strings.Split(keyword, " ")

	if len(keys) == 3 {
		keys[0] = "%" + keys[0] + "%"
		keys[1] = "%" + keys[1] + "%"
		keys[2] = "%" + keys[2] + "%"
		// conn.Debug().Where("subject like ? and subject like ? and subject like ? ", keys[0], keys[1], keys[2]).Select("subject, data_from").Order("missed_at desc").Limit(5).Find(&articles)
	} else if len(keys) == 2 {
		keys[0] = "%" + keys[0] + "%"
		keys[1] = "%" + keys[1] + "%"
		// conn.Debug().Where("subject like ? and subject like ? ", keys[0], keys[1]).Select("subject, data_from").Order("missed_at desc").Limit(5).Find(&articles)
	} else {
		keys[0] = "%" + keys[0] + "%"
		// conn.Debug().Where("subject like ?", keys[0]).Select("subject, data_from").Order("missed_at desc").Limit(5).Find(&articles)
		dbutils.ReadDB(DBName).Table("articles").Debug().Where("subject like ?", keys[0]).Select("subject, data_from").Order("missed_at desc").Limit(5).Find(&articles)
	}

	return

}
