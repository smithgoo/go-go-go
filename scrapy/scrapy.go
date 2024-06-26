package scrapy

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

// 定义一个结构体来存储链接信息
type Link struct {
	Text string `json:"text"`
	Href string `json:"href"`
}

// 定义一个结构体来存储视频的详细信息
type VideoDetailInfo struct {
	Name               string   `json:"name"`
	NickName           string   `json:"nickname"`
	Address            []string `json:"address"`
	Title              string   `json:"title"`
	Cover              string   `json:"cover"`
	Content            string   `json:"content"`
	OtherName          string   `json:"otherName"`
	VideoDirector      string   `json:"videoDirector"`
	VideoMaincharacter string   `json:"videoMaincharacter"`
	VideoType          string   `json:"videoType"`
	VideoArea          string   `json:"videoArea"`
	VideoLanguage      string   `json:"videoLanguage"`
	VideoReleaseTime   string   `json:"videoReleaseTime"`
	VideoUpdate        string   `json:"videoUpdate"`
}

// 爬取网页上的列表
func FetchContent(c *gin.Context) {
	// 从查询参数中获取 URL
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL parameter is required"})
		return
	}

	// 发起 HTTP GET 请求
	res, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status)})
		return
	}

	// 使用 goquery 解析 HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 查找 ul/li/a 结构下的 a 标签
	var result []Link
	doc.Find("div.xing_vb ul li span.xing_vb4 a").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		href, exists := s.Attr("href")
		if exists {
			result = append(result, Link{
				Text: text,
				Href: href,
			})
		}
	})

	// 返回爬取结果
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// 获取当前页面的所有详细信息
func FetchCurrentVideoInfo(c *gin.Context) {
	// 从查询参数中获取 URL
	url := c.Query("url")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL parameter is required"})
		return
	}

	// 发起 HTTP GET 请求
	res, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status)})
		return
	}

	// 使用 goquery 解析 HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 创建一个 VideoDetailInfo 实例
	var videoInfo VideoDetailInfo

	// 查找 div.vodInfo div.vodh h2 结构下的文本内容
	doc.Find("div.vodInfo div.vodh h2").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		log.Printf(text)
		videoInfo.Title = text
	})

	// 查找 div.vodinfobox ul li 结构下的文本内容
	doc.Find("div.vodinfobox ul li").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		log.Printf(text)
		switch i {
		case 0:
			videoInfo.OtherName = text
		case 1:
			videoInfo.VideoDirector = text
		case 2:
			videoInfo.VideoMaincharacter = text
		case 3:
			videoInfo.VideoType = text
		case 4:
			videoInfo.VideoArea = text
		case 5:
			videoInfo.VideoLanguage = text
		case 6:
			videoInfo.VideoReleaseTime = text
		case 7:
			videoInfo.VideoUpdate = text
		}
	})

	// 查找 div.vodBox div.vodImg img.lazy 结构下的 src 属性
	doc.Find("div.vodBox div.vodImg img.lazy").Each(func(i int, s *goquery.Selection) {
		value, exists := s.Attr("src")
		if exists {
			videoInfo.Cover = value
		}
	})

	// 查找 div.vodplayinfo ul li input 结构下的 value 属性
	doc.Find("div.vodplayinfo ul li input").Each(func(i int, s *goquery.Selection) {
		value, exists := s.Attr("value")
		if exists {
			videoInfo.Address = append(videoInfo.Address, value)
		}
	})

	// 获取 pageIndex 参数
	pageIndexStr := c.Query("pageIndex")
	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil || pageIndex < 0 || pageIndex >= len(videoInfo.Address) {
		pageIndex = 0
	}

	// 返回爬取结果到 HTML 模板
	c.HTML(http.StatusOK, "videoPlayer.html", gin.H{
		"VideoInfo": videoInfo,
		"PageIndex": pageIndex,
	})
}

