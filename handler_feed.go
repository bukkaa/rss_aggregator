package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bukkaa/rss_aggregator/internal/database"
	"github.com/google/uuid"
)

type feedParams struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, userEntity database.User) {
	params := parseFeedParams(r, w)

	feedEntity, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:     uuid.New(),
		Name:   params.Name,
		Url:    params.Url,
		UserID: userEntity.ID,
	})
	if err != nil {
		respondWithError(w, 512, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	respondWithJson(w, 201, newFeedDto(feedEntity))
}

func (apiCfg *apiConfig) handleGetFeedsForUser(w http.ResponseWriter, r *http.Request, userEntity database.User) {
	feedEntities, err := apiCfg.DB.GetFeedsByUserId(r.Context(), userEntity.ID)
	if err != nil {
		respondWithError(w, 512, fmt.Sprintf("Couldn't find feed for user {%v}: %v", userEntity.ID, err))
		return
	}
	var code = 200
	if len(feedEntities) == 0 {
		code = 204
	}
	respondWithJson(w, code, newListFeedDto(feedEntities))
}

func (apiCfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feedEntities, err := apiCfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, 512, fmt.Sprintf("Couldn't find feeds: %v", err))
		return
	}
	var code = 200
	if len(feedEntities) == 0 {
		code = 204
	}
	respondWithJson(w, code, newListFeedDto(feedEntities))
}

func parseFeedParams(r *http.Request, w http.ResponseWriter) feedParams {
	params := feedParams{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		var zeroParams feedParams
		return zeroParams
	}

	return params
}
