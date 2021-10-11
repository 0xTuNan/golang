package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get 获取待爬取网站的html
func Get(url string,agent string) string {
	ua := agent
	res,err := http.Get(url)
	res.Header.Set("User-Agent",ua)
	res.Header.Set("X-Forwarded-For","127.0.0.1")
	res.Header.Set("Referer","www.google.com")
	if err !=nil {
		fmt.Println("http get err ",err)
	}
	defer res.Body.Close()
	Body,err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read error", err)
	}
	doc := string(Body)
	return doc
}