package main

import ( 
	"time"
	"fmt"

	"github.com/pkg/errors"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func handlerAgg(s *state, cmd command) error { 
	if len(cmd.args) != 1 {
		return errors.New("usage: agg <time_between_requests>\ntime duration format is 1s, 1m, 1h, etc.")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	fmt.Printf("Collecting feeds every %s\n\n", timeBetweenRequests)
	if err != nil {
		return errors.Wrap(err, "Invalid time duration string")
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil { 
			return errors.Wrap(err, "An error occured while scraping feeds list")
		}
	}
}
