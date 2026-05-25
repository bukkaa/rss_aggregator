package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bukkaa/rss_aggregator/internal/auth"
	"github.com/bukkaa/rss_aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type userParams struct {
	Name string `json:"name"`
}

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	params := parseUserParams(r, w)

	now := time.Now().UTC()
	userEntity, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 512, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJson(w, 201, newUserDto(userEntity))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, userEntity database.User) {
	respondWithJson(w, 200, newUserDto(userEntity))
}

func (apiCfg *apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request, _ database.User) {
	params := parseUserParams(r, w)
	apiKey, _ := auth.GetAPIKey(r.Header)

	userEntity, err := apiCfg.DB.UpdateUserByAPIKey(r.Context(), database.UpdateUserByAPIKeyParams{
		ApiKey: apiKey,
		Name:   params.Name,
	})
	if err != nil {
		respondWithError(w, 512, fmt.Sprintf("Couldn't update user: %v", err))
		return
	}

	respondWithJson(w, 200, newUserDto(userEntity))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, userEntity database.User) {
	amount := 5

	amountStr := chi.URLParam(r, "amount")
	if amountStr != "" {
		amount, _ = strconv.Atoi(amountStr)
	}

	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: userEntity.ID,
		Limit:  int32(amount),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't load fresh posts for user [%v]: %v", userEntity.ID, err))
		return
	}
	respondWithJson(w, 200, newListPostsDto(posts))
}

func parseUserParams(r *http.Request, w http.ResponseWriter) userParams {
	params := userParams{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		var zeroParams userParams
		return zeroParams
	}

	return params
}
