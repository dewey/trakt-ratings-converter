package trakt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type traktClient struct {
	c        *http.Client
	username string
	clientID string
}

// NewTraktClient returns a trakt.tv API client
func NewTraktClient(username string, clientID string) (*traktClient, error) {
	return &traktClient{
		c:        http.DefaultClient,
		username: username,
		clientID: clientID,
	}, nil
}

func (r *traktClient) Watched(objectType string) ([]WatchedItem, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.trakt.tv/users/%s/watched/%s", r.username, objectType), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("trakt-api-version", "2")
	req.Header.Add("trakt-api-key", r.clientID)
	resp, err := r.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var wis []WatchedItem
	if err := json.NewDecoder(resp.Body).Decode(&wis); err != nil {
		return nil, err
	}
	return wis, nil
}

func (r *traktClient) Rated(objectType string) ([]RatedItem, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.trakt.tv/users/%s/ratings/%s", r.username, objectType), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("trakt-api-version", "2")
	req.Header.Add("trakt-api-key", r.clientID)
	resp, err := r.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var ris []RatedItem
	if err := json.NewDecoder(resp.Body).Decode(&ris); err != nil {
		return nil, err
	}
	return ris, nil
}

// WatchedItem is a movie or show that was watched with additional meta data
type WatchedItem struct {
	Plays         int       `json:"plays"`
	LastWatchedAt time.Time `json:"last_watched_at"`
	LastUpdatedAt time.Time `json:"last_updated_at"`
	Movie         Movie     `json:"movie"`
}

type Movie struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	Ids   struct {
		Trakt int    `json:"trakt"`
		Slug  string `json:"slug"`
		Imdb  string `json:"imdb"`
		Tmdb  int    `json:"tmdb"`
	} `json:"ids"`
}

// RatedItem is a movie or show that was rated with additional meta data
type RatedItem struct {
	RatedAt time.Time `json:"rated_at"`
	Rating  int       `json:"rating"`
	Type    string    `json:"type"`
	Movie   Movie     `json:"movie"`
}
