package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"net/http"
	"rssagg/internal/database"
	"time"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding params: %s", err.Error()))
		return
	}

	feed, err := apiCfg.DB.CreatedFeedFollow(r.Context(), database.CreatedFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error creating feed: %s", err.Error()))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseFeedFollowToAPIFeedFollow(feed))
}

func (apiCfg *apiConfig) getAllFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error getting feeds: %s", err.Error()))
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowsToAPIFeedFollows(feedFollows))

}

func (apiCfg *apiConfig) deleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid feed follow ID: %s", feedFollowIDStr))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		FeedID: feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error deleting feed follow: %s", err.Error()))
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Feed follow deleted successfully"})

}
