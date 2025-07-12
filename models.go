package main

import (
	"github.com/google/uuid"
	"rssagg/internal/database"
	"time"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	APIKey    string `json:"api_key"`
}

func databaseUserToAPIUser(user database.User) User {
	return User{
		ID:        user.ID.String(),
		Name:      user.Name,
		CreatedAt: user.CreatedAt.UTC().Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.UTC().Format(time.RFC3339),
		APIKey:    user.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	URL       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToAPIFeed(feed database.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		Name:      feed.Name,
		CreatedAt: feed.CreatedAt.UTC(),
		UpdatedAt: feed.UpdatedAt.UTC(),
		URL:       feed.Url,
		UserID:    feed.UserID,
	}
}

func databaseFeedsToAPIFeeds(dpFeeds []database.Feed) []Feed {
	feeds := make([]Feed, 0)

	for _, dpFeed := range dpFeeds {
		feeds = append(feeds, databaseFeedToAPIFeed(dpFeed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowToAPIFeedFollow(feedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt.UTC(),
		UpdatedAt: feedFollow.UpdatedAt.UTC(),
		UserID:    feedFollow.UserID,
		FeedID:    feedFollow.FeedID,
	}
}

func databaseFeedFollowsToAPIFeedFollows(dpFeeds []database.FeedFollow) []FeedFollow {
	feeds := make([]FeedFollow, 0)

	for _, dpFeed := range dpFeeds {
		feeds = append(feeds, databaseFeedFollowToAPIFeedFollow(dpFeed))
	}
	return feeds
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	FeedID      uuid.UUID `json:"feed_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func databasePostToAPIPost(post database.Post) Post {
	var description *string
	if post.Description.Valid {
		description = &post.Description.String
	}

	return Post{
		ID:          post.ID,
		Title:       post.Title,
		Url:         post.Url,
		Description: description,
		PublishedAt: post.PublishedAt.UTC(),
		FeedID:      post.FeedID,
		CreatedAt:   post.CreatedAt.UTC(),
		UpdatedAt:   post.UpdatedAt.UTC(),
	}
}

func databasePostsToAPIPosts(dpPosts []database.Post) []Post {
	posts := make([]Post, 0)

	for _, dpPost := range dpPosts {
		posts = append(posts, databasePostToAPIPost(dpPost))
	}
	return posts
}
