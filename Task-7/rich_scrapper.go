package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	fName := "data.csv"
	file, err := os.Create(fName)
	checkError(err)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(colly.AllowedDomains("forbes.com", "www.forbes.com"))
	c.OnHTML(".base ng-scope", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText("tr:.name"),
			e.ChildText("span"),
		})
	})
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", string(r.Body))
	})
	c.Visit("https://forbes.com/real-time-billionaires/")
}
