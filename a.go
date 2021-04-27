package main

import (
	net "bilibili/main"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"io/ioutil"
	"strings"
	"time"
)

type text struct {
	p     string
	title string
}

func main() {
	string := ""
	texts := getContent()
	for _, t := range texts {
		string += t.title + t.p
	}
	err := ioutil.WriteFile("C:\\Users\\wodeq\\Desktop\\新建文本文档.txt", []byte(string), 0777)
	if err != nil {
		panic(err)
	}

}
func getContent() []text {
	texts := make([]text, 0)

	list := net.GetPart()
	for _, s := range list {
		url := "https://www.bilibili.com/read/cv"
		text := text{}
		url += s
		html := net.GetHtml(url)
		duration := time.Duration(1) * time.Second
		time.Sleep(duration)
		title, p := GetPartDate(html)
		text.title = title
		text.p = p
		texts = append(texts, text)
	}
	return texts
}
func GetPartDate(html1 string) (string, string) {
	note, _ := html.Parse(strings.NewReader(html1))
	dom := goquery.NewDocumentFromNode(note)
	title := dom.Find("h1[class=title]").Text() + "\n\n"
	title = "\t\t\t" + title
	p := dom.Find("p[class!=title]").Text()
	p = strings.ReplaceAll(p, "。", "。\n")
	p = strings.ReplaceAll(p, title, "")
	p = strings.ReplaceAll(p, "申请成为专栏UP主查看专栏使用说明", "")
	return title, p
}
