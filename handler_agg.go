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
		return errors.New("Agg command expects a single duration string argument (1s, 1m, 1h, etc.)")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	fmt.Printf("Collecting feeds every %s\n\n", timeBetweenRequests)
	if err != nil {
		return errors.Wrap(err, "Invalid time duration string")
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)

	}
}
