package po

type PreForumPost struct {
	// gorm.Model
	Message  string `json:"message"`
	Subject  string `json:"subject"`
	Useip    string `json:"useip"`
	Pid      int64  `json:"pid"`
	Tid      int64  `json:"tid"`
	Authorid int64  `json:"authorid"`
}
