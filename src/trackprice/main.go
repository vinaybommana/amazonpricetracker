package trackprice

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetProductPrice(ProductUrl string) float64 {
	res, err := http.Get(ProductUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if s, err := strconv.ParseFloat(parseHTML(res.Body), 64); err == nil {
		return s
	}
	if err != nil {
		fmt.Println(err)
	}
	return 0.0
}

/*
parseHTML for a specified tag
	.price-container span#ProductPrice

	<h2 class="price-container">
	<span id="ProductPrice">price</span>
	</h2>
*/
func parseHTML(input io.Reader) string {
	doc, err := goquery.NewDocumentFromReader(input)
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	productPrice := doc.Find(".price-container span#ProductPrice").Text()
	return productPrice
}

/*
read urls from given filename and return
in the form of slice of strings
*/
func readURLSFromFile(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error while opening file: %v", err)
	}
	defer f.Close()

	var urls []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}

	return urls
}

func GetProductName(inputUrl string) string {
	ProductSlice := strings.Split(inputUrl, "/")
	Name := ProductSlice[len(ProductSlice)-1]
	ProductName := strings.Split(Name, "?")[0]
	return ProductName
}

// /*
// usage:
// writeFile("sample.txt", "This is a sample file\n")
// */
// func writeFile(filename string, input string) {

// 	/* create file if it does not exist */
// 	if _, isFileExistError := os.Stat(filename); os.IsExist(isFileExistError) {
// 		err := os.Remove(filename)
// 		if err != nil {
// 			log.Fatalf("Error removing file: %v", err)
// 		}
// 	} else {
// 		f, createFileError := os.Create(filename)
// 		if createFileError != nil {
// 			log.Fatal(createFileError)
// 		}
// 		defer f.Close()
// 		/* write to file */
// 		_, writeFileError := f.WriteString(input)
// 		if writeFileError != nil {
// 			log.Fatalf("Error writing to file: %v", writeFileError)
// 		}
// 	}
// }

// func main() {
// 	for _, url := range readURLSFromFile("gamenationsurls.txt") {
// 		productName := getProductName(url)
// 		productPrice := getProductPrice(url)
// 		fmt.Printf("%s: %f\n", productName, productPrice)
// 	}
// }
