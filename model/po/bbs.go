package po

import (
	"fmt"
	"github.com/airdb/sailor/config"
	"github.com/airdb/sailor/dbutils"
	"github.com/airdb/sailor/enum"
	"strings"
)

type PreForumPost struct {
	// gorm.Model
	Message  string `json:"message"`
	Subject  string `json:"subject"`
	Useip    string `json:"useip"`
	Pid      int64  `json:"pid"`
	Tid      int64  `json:"tid"`
	Authorid int64  `json:"authorid"`
}

func getBBSDBName() (dbName string) {
	env := strings.ToUpper(config.GetEnv())

	dbName = "bbs"
	if !enum.IsLiveEnv(env) {
		dbName = config.GetEnv() + "_" + dbName
	}

	return dbName
}

func GetBBSArticles() (preForumPost []PreForumPost) {
	dbName := getBBSDBName()
	fmt.Println("xxxx", dbName)

	sqltext := ""
	sqltext = "select * from pre_forum_post where  subject != '' "
	// sqltext += " and message like '%登记信息%宝贝回家编号%' "
	// sqltext += " and message like '%登记信息%编号%' "
	sqltext += " and message like '%本网站及志愿者提供的寻人服务均是免费%' "
	//sqltext += " and subject like '%3313%' "
	// sqltext += " and tid = 193856 "
	sqltext += " order by pid desc"
	sqltext += " limit 10 offset 0"
	dbutils.ReadDB(dbName).Raw(sqltext).Scan(&preForumPost)
	return
}
