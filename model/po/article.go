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

type MinaArticle struct {
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

type Article struct {
	gorm.Model
	Name      string `json:"name"`
	Title     string `json:"title"`
	Flag      uint   `json:"flag"`
	Content   string `json:"content"`
	CreatedBy uint   `json:"created_by"`
}

func getDefaultDBName() (dbName string) {
	dbName = "mina"
	if !enum.IsLiveEnv(config.GetEnv()) {
		dbName = config.GetEnv() + "_" + dbName
	}

	return dbName
}

func QueryBBSByKeywords(keyword string) (articles []*MinaArticle) {
	dbName := getDefaultDBName()
	fmt.Println("db_name", dbName)
	keys := strings.Split(keyword, " ")

	if len(keys) == 3 {
		keys[0] = "%" + keys[0] + "%"
		keys[1] = "%" + keys[1] + "%"
		keys[2] = "%" + keys[2] + "%"
		dbutils.ReadDB(dbName).Table("mina_article").Where(
			"subject like ? and subject like ? and subject like ? ", keys[0], keys[1], keys[2],
		).Select("subject, data_from").Order("missed_at desc").Limit(5).Find(&articles)
	} else if len(keys) == 2 {
		keys[0] = "%" + keys[0] + "%"
		keys[1] = "%" + keys[1] + "%"
		dbutils.ReadDB(dbName).Table("mina_article").Where(
			"subject like ? and subject like ? ", keys[0], keys[1],
		).Select("subject, data_from").Order("missed_at desc").Limit(5).Find(&articles)
	} else {
		keys[0] = "%" + keys[0] + "%"
		// dbutils.ReadDB(dbName).Table("articles").Debug().Where(
		dbutils.ReadDB(dbName).Debug().Where(
			"subject like ?", keys[0],
		).Select("subject, data_from").Order("missed_at desc").Limit(5).Find(&articles)
	}

	return
}

func FirstOrCreateArticleDataFrom(article MinaArticle) (id uint) {
	fmt.Println("xxx", getDefaultDBName())
	// db := dbutils.WriteDB(getDefaultDBName()).Table("articles").Debug().FirstOrCreate(&article, MinaArticle{DataFrom: article.DataFrom})
	db := dbutils.WriteDB(getDefaultDBName()).Debug().FirstOrCreate(&article, MinaArticle{DataFrom: article.DataFrom})
	if db.Error == nil {
		id = article.ID
	}
	return
}

func UpdateArticle(article MinaArticle) {
	// dbutils.WriteDB(getDefaultDBName()).Table("articles").Debug().Save(&article)
	dbutils.WriteDB(getDefaultDBName()).Debug().Save(&article)
}

func ListArticle() (article []*Article) {
	dbutils.DefaultDB().Table("article_tab").Debug().Limit(10).Find(&article)
	return
}
