package po

import (
	"fmt"
	"github.com/airdb/sailor/config"
	"github.com/airdb/sailor/dbutils"
	"github.com/airdb/sailor/enum"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
)

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

func getDefaultDBName() (dbName string) {
	config.Init()
	env := config.GetEnv()

	if enum.IsLiveEnv(env) {
		dbName = "mina"
	} else {
		dbName = config.GetEnv() + "_mina"
	}

	return dbName
}

func QueryBBSByKeywords(keyword string) (articles []*Article) {
	dbName := getDefaultDBName()
	fmt.Println("db_name", dbName)
	keys := strings.Split(keyword, " ")

	if len(keys) == 3 {
		keys[0] = "%" + keys[0] + "%"
		keys[1] = "%" + keys[1] + "%"
		keys[2] = "%" + keys[2] + "%"
		dbutils.ReadDB(dbName).Table("articles").Where(
			"subject like ? and subject like ? and subject like ? ", keys[0], keys[1], keys[2],
		).Select("subject, data_from").Order("missed_at desc").Limit(5).Find(&articles)
	} else if len(keys) == 2 {
		keys[0] = "%" + keys[0] + "%"
		keys[1] = "%" + keys[1] + "%"
		dbutils.ReadDB(dbName).Table("articles").Where(
			"subject like ? and subject like ? ", keys[0], keys[1],
		).Select("subject, data_from").Order("missed_at desc").Limit(5).Find(&articles)
	} else {
		keys[0] = "%" + keys[0] + "%"
		dbutils.ReadDB(dbName).Table("articles").Debug().Where(
			"subject like ?", keys[0],
		).Select("subject, data_from").Order("missed_at desc").Limit(5).Find(&articles)
	}

	return
}
