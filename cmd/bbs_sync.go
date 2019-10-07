package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("vim-go")
		time.Sleep(2 * time.Second)

	}
}

/*
// 去掉html中所有标签
func trimHtml(src string) string {

	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	re, _ = regexp.Compile("\\[align=left\\]")
	// src = re.ReplaceAllString(src, "\n\\[align=left\\]")
	src = re.ReplaceAllString(src, "\n")

	// //去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	// [attach]
	// src = strings.Replace(src, "[attach]", "[attach]图片ID：", -1)
	src = strings.Replace(src, "&nbsp;", "", -1)
	src = strings.Replace(src, "· ", "\r", -1)

	// 去掉[]标签, 如 [color=#000000]
	re, _ = regexp.Compile("\\[[\\S\\s]+?\\]")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	// src = re.ReplaceAllString(src, "\n")
	src = re.ReplaceAllString(src, "")

	// //去除连续的换行符
	// re, _ = regexp.Compile("\\s{4,}")
	// src = re.ReplaceAllString(src, "\n")

	return strings.TrimSpace(src)
}

func formatTime(tstr string) (tm time.Time, err error) {
	// timeFormats := []string{"2006-01-02 15:04:05", "2006-01-0215:04:05", "2006-01", "2006-1-02", "2006-01-2", "2006-1-2", "2006-01-02", "2006-0102", "2006年1月2日", "2006年1月02日", "2006年01月2日", "2006年01月02日", "2006年01月**日", "2006年1月", "2006年"}
	// timeFormats := []string{"2006-01-02 15:04:05", "2006-01-02150405", "2006-01-0215:04:05", "2006-01", "2006-1-02", "2006-01-2", "2006-1-2", "2006-01-02", "2006-0102", "2006--01-02", "2006年1月2日", "2006年1月02日", "2006年01月2日", "2006年01月02日", "2006年01月**日", "2006年1月", "2006年", "2006年01月02", "2006年01月02月", "2006月01月02日", "2006年01年02月", "2006年01年02日", "2006年1月02", "2006"}
	timeFormats := []string{"2006-1月", "2006-01月", "2006年01", "2006年1", "006-1-2", "06年1月2", "20060102", "2006年1月2", "06>年1月2日", "2006", "2006-01-02 15:04:05", "2006-01-0215:04:05", "2006-01", "2006-1-02", "2006-01-2", "2006-1-2", "2006-01-02", "2006-0102", "2006--01-02", "2006年1月2日", "2006年1月02日", "2006年01月2日", "2006年01月02日", "2006年01月**日", "2006年1月", "2006年", "2006年01月02", "2006年01月02月", "2006月01月02日", "2006年01年02月", "2006年01年02日", "2006年1月02"}

	for _, timeFormat := range timeFormats {
		tm, err = time.ParseInLocation(timeFormat, tstr, time.Local)
		if err == nil {
			break
		}
	}
	return
}

func parseHtml(datafrom, title, msg string) (article models.Article) {
	article.DataFrom = datafrom
	article.Title = title
	article.Subject = title

	var err error

	infoList := strings.Split(msg, "\r")
	if len(infoList) <= 1 {
		// beego.Error("=======fail", datafrom, "==========")
		infoList = strings.Split(msg, "\n")
	}

	var start bool
	var end bool
	var detailFlag bool
	// 因为详细描述多行都是value, 所以要在外层定义
	var key string
	var value string
	for _, info := range infoList {
		info = strings.TrimLeft(info, "\n")

		if "" == info || strings.Contains(info, "本帖最后由") {
			if strings.Contains(info, "本帖最后由") {
				start = false
			}
			if !start {
				end = false
			}
			continue
		}

		if !start && (strings.Contains(info, "：") || strings.Contains(info, ":")) {
			start = true
		}

		// 如果没有标记为开始, 就已经标记为结束了，则退出
		if !start || end {
			continue
		}
		if strings.Contains(info, "站务电话") || strings.Contains(info, "注册时间") {
			detailFlag = false
		}
		if !detailFlag {
			exp := regexp.MustCompile(`.*：|.*:`)
			// 最短匹配 key, 去除 key 的部分就是 value 或是 value 的一部分
			key = exp.FindString(info)
			value = strings.Replace(info, key, "", -1)

			// 处理key 和 value
			// 删除key中的非汉字和空格
			exp = regexp.MustCompile(`(\[.+?\])|(\^M)|：|:|(\s)|[^\p{Han}]| `)
			key = exp.ReplaceAllString(key, "")
		} else {
			value = info
		}
		beego.Error(key, detailFlag, value)

		// tmexp := regexp.MustCompile(`[（|(].*[）|)]|农历|阳历|不准确|大概|日期不确定|不确定|约|X日|份左右|期不详。|。|具体.+?|失踪|宋振彪|宋振邦2008年|.*身份证日期|左右|号|阴历|某天|夏.*|年底|腊月.*|九.*|八.*|天已记不清楚|冬月.*|正.*|初.*|下午1点半|大|生`)
		tmexp := regexp.MustCompile(`[（|(].*[）|)]|农历|公历|古历|阳历|旧历|不准确|大概|日期不确定|不确定|约|X日|份左右|期不
详。|。|具体.*|失踪.*|宋振彪|宋振邦2008年|.*身份证日期|左右|号|阴历|某天|夏.*|年底|腊月.*|九.*|八.*|天已记不清楚|冬月.*|正.*|初.*|下>午.*|大|0{4,}|生.*|出.*|元月.*|上午.*|晚上.*|十.*|七.*|五.*|一.*|二.*|三.*|四.*|~.*|、.*|或.*|&.*|失|————|到.*|暑假|是|和.*|《.*|（.*|冬天|之间|早上.*|深秋|底|春.*|秋.*|●|·`)

		switch key {
		case "寻亲类别", "类别":
			article.Category = value
		case "宝贝回家编号", "编号", "寻亲编号":
			article.Babyid = value
		case "姓名":
			article.Nickname = value
		case "性别":
			// 值为1时是男性，值为2时是女性，值为0时是未知
			if "女" == value {
				// 1 --> 2
				article.Gender = 2
			} else if "男" == value {
				// 0 -> 1
				article.Gender = 1
			}
		case "失踪时身高":
			article.Height = value
		case "失踪地点", "失踪地址", "地址":
			article.MissedAddress = value
		case "失踪者特征描述":
			article.Characters = value
		case "失踪人户籍所在地", "失踪人所在省", "籍贯":
			article.BirthedProvince = value
		case "采血情况":
		case "出生日期":
			value = tmexp.ReplaceAllString(value, "")
			value = strings.Replace(value, "—", "", -1)
			article.BirthedAt, err = formatTime(value)
			if err != nil {
				beego.Error("出生日期", value, err)
			}
			// beego.Info(key, value, ", 格式化:", article.BirthedAt)
		case "失踪日期", "失踪时间":
			value = tmexp.ReplaceAllString(value, "")
			value = strings.Replace(value, "—", "", -1)
			article.MissedAt, err = formatTime(value)
			// tm = tm.Format("2006-01-02 15:04:05")
			if err != nil {
				beego.Error("失踪日期", value, err)
			}
			// beego.Info(key, value, ", 格式化:", article.MissedAt)
		case "注册时间", "站务电话":
			end = true
			detailFlag = false
			break
		case "其他资料", "其他情况", "共同经历资料":
			detailFlag = true
			article.Details += value
		default:
			// beego.Error("this is default", key)
		}
	}
	beego.Info("Babyid:", article.Babyid, ", 数据来源:", article.DataFrom )
	return
}

func syncFrombbs() {
	// preForumPosts := models.GetLastBBSInfo()
	preForumPosts := models.GetAllBBSInfo()

	for i, preForumPost := range preForumPosts {
		datafrom := "https://bbs.baobeihuijia.com/thread-"
		datafrom += strconv.FormatInt(int64(preForumPost.Tid), 10) + "-1-1.html"
		beego.Info(i, datafrom, preForumPost.Subject)
		var art models.Article
		art.Subject = preForumPost.Subject
		art.DataFrom = datafrom

		models.AddArticleDataFrom(art)
		continue
		// beego.Info("====", datafrom)

		// beego.Error(preForumPost.Tid, preForumPost.Pid)
		msg := trimHtml(preForumPost.Message)
		article := parseHtml(datafrom, preForumPost.Subject, msg)
		if article.Babyid == "" {
			beego.Critical("update datafrom only, this babyid is null.", article.DataFrom)
			article.SyncStatus = -1
			models.AddArticleDataFrom(article)
			continue
		}

		article.UUID = uuid.New().String()
		models.AddArticle(article)
		models.SyncPictureFromBbs(preForumPost.Tid, preForumPost.Pid, article.Babyid, article.UUID)
		if beego.BConfig.RunMode == "prod" {
		}
		return
	}
}
*/
