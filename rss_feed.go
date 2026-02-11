package main

import (
	"fmt"
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
		return &RSSFeed{}, errors.Wrap(err, "Failed to create HTTP request")
	}
	req.Header.Set("User-Agent", "gator")
	
	res, err := client.Do(req)
	if err != nil { 
		return &RSSFeed{}, errors.Wrap(err, "Failed to get HTTP response")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil { 
		return &RSSFeed{}, errors.Wrap(err, "Failed to read HTTP response body")
	}

	var feed *RSSFeed
	err = xml.Unmarshal(data, &feed)
	if err != nil { 
		return &RSSFeed{}, errors.Wrap(err, "Failed to unmarshal XML data")
	}

	return feed, nil
}

func printFeed(feed *RSSFeed) {
	fmt.Printf("Title: %s\n", feed.Channel.Title)
	fmt.Printf("Description: %s\n", feed.Channel.Description)
	for _, item := range feed.Channel.Item { 
		fmt.Printf("\tTitle: %s\n", item.Title)
		fmt.Printf("\tLink: %s\n", item.Link)
		fmt.Printf("\tDescription: %s\n", item.Description)
		fmt.Printf("\tPubDate: %s\n\n", item.PubDate)
	}
}

