package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

type QuotePageInfo struct {
	Count   int `json:"count"`
	Results []struct {
		Quote  string `json:"quote"`
		Author string `json:"author"`
	} `json:"results"`
}

// FetchRandomQuote fetches a single random quote
func FetchRandomQuote() (string, error) {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// First, fetch the first page to get the total count and calculate the total number of pages
	pageInfo, err := fetchPageInfo(1) // Fetch first page to get total count
	if err != nil {
		return "", err
	}

	totalCount := pageInfo.Count
	pageSize := len(pageInfo.Results)
	totalPages := totalCount / pageSize
	if totalCount%pageSize != 0 {
		totalPages++ // Add one more page if there's a remainder
	}

	// Select a random page
	randomPage := rand.Intn(totalPages) + 1

	// Fetch the random page selected
	randomPageInfo, err := fetchPageInfo(randomPage)
	if err != nil {
		return "", err
	}

	// Select a random quote from this page
	randomQuoteIndex := rand.Intn(len(randomPageInfo.Results))
	selectedQuote := randomPageInfo.Results[randomQuoteIndex]

	quote := fmt.Sprintf("\"%s\" - %s", selectedQuote.Quote, selectedQuote.Author)
	return quote, nil
}

// fetchPageInfo fetches information for a given page and returns the page info including the total count of quotes
func fetchPageInfo(page int) (QuotePageInfo, error) {
	var pageInfo QuotePageInfo
	url := fmt.Sprintf("https://philosophyapi.pythonanywhere.com/api/ideas/?page=%d", page)
	resp, err := http.Get(url)
	if err != nil {
		return pageInfo, fmt.Errorf("error fetching page %d: %w", page, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return pageInfo, fmt.Errorf("error reading response body: %w", err)
	}

	err = json.Unmarshal(body, &pageInfo)
	if err != nil {
		return pageInfo, fmt.Errorf("error parsing JSON: %w", err)
	}

	return pageInfo, nil
}

// newPhilosophicalQuoteCmd creates a new command for fetching a single random philosophical quote
func newPhilosophicalQuoteCmd() *cobra.Command {
	quoteCmd := &cobra.Command{
		Use:   "philosophy",
		Short: "Fetches and prints a random philosophical quote",
		Long:  "Fetches and prints a random philosophical quote from the philosophy API.",
		Run: func(cmd *cobra.Command, args []string) {
			quote, err := FetchRandomQuote()
			if err != nil {
				fmt.Printf("Error fetching quote: %v\n", err)
				return
			}

			fmt.Println(quote)
		},
	}

	return quoteCmd
}
