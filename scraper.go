package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/bukkaa/rss_aggregator/internal/database"
	"github.com/google/uuid"
)

func startScraping(db *database.Queries, concurrency int, cooldown time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, cooldown)

	ticker := time.NewTicker(cooldown)

	// run this for-loop every time the ticker would send a signal through the channel (C)
	// which happens every cooldown period.
	// first fire happens immediately
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("error fetching feeds:", err)
			continue
		}

		// allow us to accumulate as many goroutines as we need into one group
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {

			// for every feed received, increase internal counter of members in the group
			wg.Add(1)

			go scrapeFeed(db, feed, wg)

			// probably better to use this one istead:
			// wg.Go()
		}

		// the code will block and wait here until we got all Dones from every group member
		wg.Wait()
	}

	// this construction also works every cooldown period,
	// but it will fire the first run only after cooldown period too
	// for range <-ticker.C { ... }
}

func scrapeFeed(db *database.Queries, feed database.Feed, wg *sync.WaitGroup) {
	// call Done() when action performed
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed [%v] as fetched: %v", feed.ID, err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed [%v]: %v", feed.ID, err)
		return
	}

	for _, feedItem := range rssFeed.Channel.Item {
		descr := sql.NullString{}
		if feedItem.Description != "" {
			descr.String = feedItem.Description
			descr.Valid = true
		}

		publishedAt, err := time.Parse(time.RFC1123Z, feedItem.PublicationDate)
		if err != nil {
			log.Printf("couldn't parse date [%v], msg: %v\n", feedItem.PublicationDate, err)
			continue
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			Title:       feedItem.Title,
			Description: descr,
			PublishedAt: publishedAt,
			Url:         feedItem.Link,
			FeedID:      feed.ID,
		})

		if err != nil {
			if !strings.Contains(err.Error(), "duplicate key") {
				log.Println("couldn't create post", err)
			}
			continue
		}
	}
	log.Printf("Great! Feed %s collected, %d posts from! \n", feed.Name, len(rssFeed.Channel.Item))
}
