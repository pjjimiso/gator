package main

import (
	"fmt"
	"log"
	"html"
	"context"
	"net/http"
	"encoding/xml"
	"time"
	"io"

	"github.com/pkg/errors"
)

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := &http.Client{
		Timeout: 10* time.Second,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create HTTP request")
	}

	req.Header.Set("User-Agent", "gator")
	res, err := client.Do(req)
	if err != nil { 
		return nil, errors.Wrap(err, "Failed to get HTTP response")
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil { 
		return nil, errors.Wrap(err, "Failed to read HTTP response body")
	}

	var feed RSSFeed
	err = xml.Unmarshal(data, &feed)
	if err != nil { 
		return nil, errors.Wrap(err, "Failed to unmarshal XML data")
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i, item := range feed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		feed.Channel.Item[i] = item
	}

	return &feed, nil
}

func printFeed(feed *RSSFeed) {
	fmt.Printf("Title: %s\n", feed.Channel.Title)
	fmt.Printf("Description: %s\n", feed.Channel.Description)
	for _, item := range feed.Channel.Item { 
		if item.Title != "" {
			fmt.Printf("\tFound post: %s\n", item.Title)
		}
	}
}

func scrapeFeeds(s *state) error {
	feedData, err := s.dbQueries.GetNextFeedToFetch(context.Background())
	if err != nil { 
		return errors.Wrapf(err, "Failed to fetch feed: %s", feedData.Url)
	}

	log.Println("Scraping feed:", feedData.Url)

	err = s.dbQueries.MarkFeedFetched(context.Background(), feedData.ID)
	if err != nil {
		return errors.Wrapf(err, "Failed to mark the feed as fetched: %s", feedData.Url)
	}

	feed, err := fetchFeed(context.Background(), feedData.Url)
	if err != nil { 
		return errors.Wrapf(err, "Failed to fetch rss feed for url: %s", feedData.Url)
	}

	printFeed(feed)
	log.Printf("Feed %s collected, %v posts found\n\n", feedData.Name, len(feed.Channel.Item))
	
	return nil
}
