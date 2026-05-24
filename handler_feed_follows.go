package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bukkaa/rss_aggregator/internal/database"
	"github.com/google/uuid"
)

type feedFollowsParams struct {
	FeedId uuid.UUID `json:"feed_id"`
}

func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, userEntity database.User) {
	params := parseFeedFollowsParams(r, w)

	feedFollowEntity, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: userEntity.ID,
		FeedID: params.FeedId,
	})
	if err != nil {
		respondWithError(w, 512, fmt.Sprintf("Couldn't follow the feed [%v] by the user [%v]: %v", params.FeedId, userEntity.ID, err))
		return
	}

	respondWithJson(w, 201, newFeedFollowsDto(feedFollowEntity))
}

func parseFeedFollowsParams(r *http.Request, w http.ResponseWriter) feedFollowsParams {
	params := feedParams{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		var zeroParams feedFollowsParams
		return zeroParams
	}

	return params
}
