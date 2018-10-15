package scrapper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Movie struct {
	Title 		string
	URL 		string
	FoundAt 	*time.Time
}

/*
 * Open home page, parse all movies that are on show now and run their concurrent country checks
 */
func Handler() (error) {
	// Request the HTML page.
	res, err := http.Get(os.Getenv("CINEMA_URL"))
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Find movie items
	movies := doc.Find("#block-system-main .views-fluid-grid-item")

	c := make(chan error)
	defer close(c)
	movies.Each(func(i int, s *goquery.Selection) {
		// For each item found, get the url and title
		url, exists := s.Find(".card .side.back a").Attr("href")
		if exists {
			movie := Movie{ URL: url }
			movie.Title = s.Find(".views-field-title a").Text()

			go func(movie Movie) {
				fmt.Printf("Scrapping: %s - %s\n", movie.URL, movie.Title)
				c <- checkCountry(movie)
			}(movie)

		}
	})

	// collect errors from processes, return them all
	var errList []string
	for i := 0; i < movies.Length(); i++{
		err := <- c
		if err != nil {
			errList = append(errList, err.Error())
		}
	}

	if len(errList) > 0 {
		return fmt.Errorf("errors have been occured while scrapping: %v", strings.Join(errList, "; "))
	}

	return nil
}

// open movie page, find its country and write to DB if does not exists
func checkCountry(movie Movie) error {
	// Request the HTML page.
	res, err := http.Get(movie.URL)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	doc.Find(".field-name-field-country .field-item").Each(func(i int, s *goquery.Selection) {
		if s.Find("a").Text() == "Україна" {
			log.Println("Yay! UA movie found:" + movie.Title)
			// TODO write movie data to Redis or Mongo (title, url, createdAt)
		}
	})

	return nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(Handler)
}
