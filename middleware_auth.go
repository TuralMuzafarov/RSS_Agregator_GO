package main

import (
	"fmt"
	"net/http"

	"github.com/TuralMuzafarov/RSS_Agregator_GO/internal/auth"
	"github.com/TuralMuzafarov/RSS_Agregator_GO/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Authentication Failed with error: %v", err))
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Could not get user with given apiKey: %v", err))
			return
		}
		handler(w, r, user)
	}
}
