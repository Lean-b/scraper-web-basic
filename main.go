package main

import (
	"fmt"
	"time"
	"github.com/gocolly/colly"
)

func main(){

	var input string
	fmt.Scanln(&input)
	link := input

	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)


	//Web sites
	c.OnHTML("a[href]", func(r *colly.HTMLElement) {
		r.Request.Visit(r.Attr("href"))
	})

	//Telefono
	c.OnHTML("a[href^='phone:']",func(r *colly.HTMLElement)  {
		r.Request.Visit(r.Attr("href"))
	})

	//Linkedin
	c.OnHTML("a[href^='https://www.linkedin.com/']",func(r *colly.HTMLElement)  {
		r.Request.Visit(r.Attr("href"))
	})

	//Twitter
	c.OnHTML("a[href^='https://twitter.com/']",func(r *colly.HTMLElement)  {
		r.Request.Visit(r.Attr("href"))
	})

	//Email
	c.OnHTML("a[href*=\"@gmail.com\"]",func(r *colly.HTMLElement)  {
		r.Request.Visit(r.Attr("href"))
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visint", r.StatusCode , r.Request.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error", err)
	})

	c.Visit(link)

}