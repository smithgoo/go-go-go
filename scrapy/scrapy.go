package scrapy

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)


// 爬取网页内容
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

	// 查找网页中的元素
	var result []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		result = append(result, text)
	})

	// 返回爬取结果
	c.JSON(http.StatusOK, gin.H{"data": result})
}
