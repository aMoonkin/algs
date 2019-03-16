package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	err := c.Post("http://example.com/login", map[string]string{
		"usrename": "admin",
		"password": "admin",
	})

	fmt.Println(err)

	c.OnResponse(func(r *colly.Response) {
		log.Println("123", r.StatusCode)
	})

	c.Visit("https://example.com/")
}
