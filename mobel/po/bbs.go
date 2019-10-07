package po

import (
	"github.com/airdb/sailor/config"
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