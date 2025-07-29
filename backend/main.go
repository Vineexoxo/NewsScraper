package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"github.com/go-rod/rod"
)



func fetchNews() []string {
	url := "https://newsdata.io/api/1/latest?apikey=pub_fa0e44ad2b334bb98df969e2386017cf&q=finance&language=en"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {


		fmt.Println("Error creating request:", err)
		return []string{}
	}
	fmt.Println("Request URL:", req.URL.String())
	fmt.Println("Response:", req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return []string{}
	}
	defer resp.Body.Close()


	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return []string{}
	}
	var result NewsAPIResponse
	err=json.Unmarshal(body, &result)
	fmt.Println("Unmarshalled response:", result.Results)
	var links []string
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: received non-200 response status:", resp.Status)
	}else{
		articles:= result.Results
		for _, article:= range articles {
			fmt.Println("Title:", article.Title)
			fmt.Println("Description:", article.Description)
			fmt.Println("URL:", article.Link)
			fmt.Println("--------------------")
			links = append(links, article.Link)
		}
		fmt.Println("Total articles fetched:", len(links))
	}
	return links
}

type ArticleCrawl struct {
	Title   string `json:"title"`
	Content string ""
	URL     string `json:"url"`
}

func CrawlTheArticles_Colly(links []string) []ArticleCrawl {
	c := colly.NewCollector()
	Articles := make([]ArticleCrawl, 0)	
	for _, link := range links {
		fmt.Println("Crawling article at URL:", link)
		c.Visit(link)
		var aCrawl ArticleCrawl
		aCrawl.URL = link
		c.OnHTML("article", func(e *colly.HTMLElement) {
			aCrawl.Title = e.ChildText("h1")
			aCrawl.Content = e.ChildText("p")
		})
		c.OnError(func(r *colly.Response, err error) {
			fmt.Println("Error occurred while crawling:", err)
			fmt.Println("Response status code:", r.StatusCode)
			fmt.Println("Response body:", string(r.Body))
		})
		Articles = append(Articles, aCrawl)
	}
	return Articles
}

func CrawlTheArticles_Rod(links []string, articles []ArticleCrawl) []ArticleCrawl {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	for idx, link := range links {
		if articles[idx].Content != "" {
			fmt.Println("Skipping already crawled article at URL:", link)
			continue
		}

		fmt.Println("Crawling article at URL:", link)
		page := browser.MustPage(link)
		page.MustWaitLoad()
		fmt.Println("Page loaded successfully:", link)
		if page.MustElement("h1").MustText() == "" {
			fmt.Println("No title found for article at URL:", link)
			continue
		}
		var aCrawl ArticleCrawl
		aCrawl.URL = link
		aCrawl.Title = page.MustElement("h1").MustText()
		aCrawl.Content = page.MustElement("p").MustText()
		articles[idx]=aCrawl
	}

	return articles
}

func main() {
	links := fetchNews()
	if len(links) == 0 {
		fmt.Println("No links fetched, exiting.")
		return
	}
	Articles:=CrawlTheArticles_Colly(links)

	Articles = CrawlTheArticles_Rod(links, Articles)

	fmt.Println("Crawling completed.")
	fmt.Println("Total articles crawled:", len(links))
	fmt.Println("Crawling finished successfully.")
	fmt.Println("Exiting the program.")
}
