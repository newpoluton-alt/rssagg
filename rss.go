package main

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"time"
)

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

func urlToFeed(url string) (RSSFeed, error) {

	httpClient := &http.Client{Timeout: 10 * time.Second}

	resp, err := httpClient.Get(url)

	if err != nil {
		return RSSFeed{}, err
	}

	defer func(Body io.ReadCloser) {
		err2 := Body.Close()
		if err2 != nil {
			log.Println("Error closing body ", err2)
		}
	}(resp.Body)

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return RSSFeed{}, err
	}
	rssFeed := RSSFeed{}
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return RSSFeed{}, err
	}

	return rssFeed, nil
}
