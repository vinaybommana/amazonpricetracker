package ProductDetails

import (
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly"
)

func GetProductDetails(productName string) map[string]string {
	products := make(map[string]string)

	c := colly.NewCollector(
		colly.AllowedDomains("gamenation.in"),
	)

	// Callback for links on scraped pages
	c.OnHTML("div.game-card", func(e *colly.HTMLElement) {
		Title := e.ChildText(".title")
		// fmt.Println(Title)
		Pricing := e.ChildText(".pricing b")
		// fmt.Println(Pricing)
		fmt.Fprintf(os.Stdout, "The Price of %s : %s\n", Title, Pricing)
		products[Title] = Pricing
	})

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	visitUrl := fmt.Sprintf("https://gamenation.in/Search?term=%s", productName)
	c.Visit(visitUrl)

	return products
}

// func main() {
// 	fmt.Println(GetProductDetails("metro"))
// }
