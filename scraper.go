package main

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"log"
	"rssagg/internal/database"
	"strings"
	"sync"
	"time"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines with every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("Error getting feeds to scrape: %s", err)
			continue
		}
		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, feed, wg)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, feed database.Feed, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("Scraping feed %s (%s)", feed.Name, feed.Url)
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)

	if err != nil {
		log.Printf("Error marking feed %s as fetched: %s", feed.ID, err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching feed %s: %s", feed.Url, err)
		return
	}

	for _, item := range rssFeed.Channel.Item {

		description := sql.NullString{}

		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}

		t, err2 := time.Parse(time.RFC1123Z, item.PubDate)

		if err2 != nil {
			log.Printf("Error parsing date for item %s: %s", item.Title, err2)
			continue
		}

		_, err2 = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: description,
			FeedID:      feed.ID,
			PublishedAt: t.UTC(),
		})
		if err2 != nil {
			if strings.Contains(err2.Error(), "duplicate key") {
				continue
			}
			log.Printf("Error creating post for feed %s: %s", feed.ID, err2)
		}
	}
}
