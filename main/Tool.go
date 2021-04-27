package net

import (
	"regexp"
	"strings"
)

//var url="https://api.bilibili.com/x/article/list/articles?id=99196&jsonp=jsonp"
var url = "https://api.bilibili.com/x/article/list/articles?id=61150&jsonp=jsonp"

func GetPart() []string {
	html1 := GetHtml(url)
	list1 := regexp.MustCompile(",\\{\"id\":.*?,\"title\":").FindAllString(html1, -1)
	list2 := make([]string, 0)
	for _, s := range list1 {
		s = strings.Split(s, ":")[1]
		s = strings.Split(s, ",")[0]
		list2 = append(list2, s)
	}
	return list2
}
