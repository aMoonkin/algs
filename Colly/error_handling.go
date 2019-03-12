package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("*", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("req url:", r.Request.URL, "failed with res", r, "\nError:", err)

	})

	c.Visit("https://squirrel.tech")
}
