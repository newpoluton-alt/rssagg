package main

import (
	"fmt"
	"net/http"
	"rssagg/internal/auth"
	"rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, error := auth.GetApiKey(r.Header)
		if error != nil {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Error getting API key: %s", error.Error()))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				respondWithError(w, http.StatusNotFound, "User not found")
				return
			}
			respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting user: %s", err.Error()))
			return
		}

		handler(w, r, user)
	}
}
