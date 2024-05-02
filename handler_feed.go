package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kuromii5/osu-parser/internal/database"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type params struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	p := params{}
	err := decoder.Decode(&p)
	if err != nil {
		respondWithErr(w, 400, "Parsing JSON error")
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      p.Name,
		Url:       p.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithErr(w, 400, "Creating user error")
		return
	}

	respondWithJSON(w, 201, convertDBFeedToFeed(feed))
}

func (apiCfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithErr(w, 400, "Couldn't get feeds")
		return
	}

	respondWithJSON(w, 200, convertDBFeedsToFeeds(feeds))
}
