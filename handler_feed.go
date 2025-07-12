package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"net/http"
	"rssagg/internal/database"
	"time"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding params: %s", err.Error()))
		return
	}

	feed, err := apiCfg.DB.CreatedFeed(r.Context(), database.CreatedFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating feed: %s", err.Error()))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedToAPIFeed(feed))
}

func (apiCfg *apiConfig) getAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting feeds: %s", err.Error()))
	}

	respondWithJSON(w, http.StatusOK, databaseFeedsToAPIFeeds(feeds))

}
