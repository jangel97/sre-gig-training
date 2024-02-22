package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

type JokeResponse struct {
	Joke string `json:"joke"`
}

func IsValidCategory(category string) bool {
	switch category {
	case
		"Programming",
		"Miscellaneous":
		return true
	}
	return false
}

func GenerateJoke(category string) (string, error) {
	if !IsValidCategory(category) {
		return "", errors.New("invalid category")
	}

	url := fmt.Sprintf("https://v2.jokeapi.dev/joke/%s?blacklistFlags=nsfw,religious,political,racist,sexist,explicit&type=single", category)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error fetching joke: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	var jokeResp JokeResponse
	err = json.Unmarshal(body, &jokeResp)
	if err != nil {
		return "", fmt.Errorf("error parsing JSON: %w", err)
	}

	return jokeResp.Joke, nil
}

func newJokeCmd() *cobra.Command {
	var category string

	jokeCmd := &cobra.Command{
		Use:   "joke",
		Short: "Fetches and prints a joke",
		Long:  "Fetches and prints a joke from a specified category. Valid categories are 'Programming' and 'Miscellaneous'.",
		Run: func(cmd *cobra.Command, args []string) {
			joke, err := GenerateJoke(category)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println(joke)
		},
	}

	// Adding a string flag "category" with default value "Programming"
	jokeCmd.Flags().StringVarP(&category, "category", "c", "Programming", "The category of the joke (Programming or Miscellaneous)")

	return jokeCmd
}
