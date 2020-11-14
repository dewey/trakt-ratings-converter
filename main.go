package main

import (
	"fmt"
	"os"
	"trakt-watchlist-converter/trakt"

	"github.com/gocarina/gocsv"
)

func main() {
	clientID := os.Getenv("CLIENT_ID")
	if clientID == "" {
		fmt.Println("CLIENT_ID can't be empty, set environment variable")
		return
	}
	username := os.Getenv("TRAKT_USERNAME")
	traktClient, err := trakt.NewTraktClient(username, clientID)
	if err != nil {
		fmt.Println(err)
		return
	}
	ratedMovies, err := traktClient.Rated("movies")
	if err != nil {
		fmt.Println(err)
		return
	}
	var wi []RatedItemExport
	for _, item := range ratedMovies {
		// This case should never happen
		if item.Rating == 0 {
			continue
		}
		wi = append(wi, RatedItemExport{
			IMDb:   item.Movie.Ids.Imdb,
			Rating: item.Rating,
			RatedItem: RatedItem{
				DateRated: item.RatedAt.Format("2006-01-02"),
				Title:     item.Movie.Title,
				Year:      item.Movie.Year,
			},
		})
	}

	// Create IMDb compatible ratings file
	watchlistFile, err := os.OpenFile("trakt-ratings.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer watchlistFile.Close()

	if err := gocsv.MarshalFile(wi, watchlistFile); err != nil {
		fmt.Println(err)
	}
}

// RatedItem contains a rated item as defined by the IMDb export
type RatedItem struct {
	Position    string `csv:"Position"`
	ID          string `csv:"Const"`
	Created     string `csv:"Created"`
	Modified    string `csv:"Modified"`
	Description string `csv:"Description"`
	Title       string `csv:"Title"`
	URL         string `csv:"URL"`
	TitleType   string `csv:"Title Type"`
	IMDbRating  string `csv:"IMDb Rating"`
	RuntimeMins string `csv:"Runtime (mins)"`
	Year        int    `csv:"Year"`
	Genres      string `csv:"Genres"`
	NumVotes    string `csv:"NumVotes"`
	ReleaseDate string `csv:"Release Date"`
	Directors   string `csv:"Directors"`
	YourRating  int    `csv:"Your Rating"`
	DateRated   string `csv:"Date Rated"`
}

// RatedItemExport contains a rated item from IMDb with better naming. This is used to import the exported file into other
// services that expect well-named columns.
type RatedItemExport struct {
	IMDb   string `csv:"IMDb"`
	Rating int    `csv:"Rating"`
	RatedItem
}
