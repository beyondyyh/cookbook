package main

import (
	"fmt"
	"regexp"
	"time"

	"github.com/gocolly/colly/v2"
)

var visited = map[string]bool{}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.douban.com", "movie.douban.com"),
		colly.MaxDepth(1),
		// colly.Async(true), // 异步请求也抓
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36"),
		// colly.Debugger(&debug.LogDebugger{}),
	)

	// 匹配详情页
	detailRegex, _ := regexp.Compile(`/subject/\d+`)
	// 匹配列表页
	listRegex, _ := regexp.Compile(`/chart`)

	// 所有a标签上设置回调函数
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		// 不是想要的详情页 or 列表页
		if !detailRegex.Match([]byte(link)) && !listRegex.Match([]byte(link)) {
			// fmt.Printf("not match: %s\n", link)
			return
		}

		// 已访问过的url
		if visited[link] {
			// fmt.Printf("******visited link: %s\n", link)
			return
		}

		// time.Sleep(1 * time.Second)
		fmt.Printf("match: %s\n", link)

		visited[link] = true

		time.Sleep(2 * time.Millisecond)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL.String())
	})

	// After request save to xxxx
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Saving: ", r.Request.URL.String())
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://leetcode-cn.com/tag/tree/")
	// c.Visit("https://movie.douban.com/chart")

}
