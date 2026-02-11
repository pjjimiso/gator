package main

import ( 
	"context"
	"fmt"
	"html"

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
	if len(cmd.args) > 0 {
		return errors.New("Agg command doesn't expect any arguments")
	}

	fetchURL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), fetchURL)
	if err != nil { 
		return errors.Wrap(err, "Failed to fetch feed")
	}

	feed.Channel.Title = fmt.Sprintf(html.UnescapeString(feed.Channel.Title))
	feed.Channel.Description = fmt.Sprintf(html.UnescapeString(feed.Channel.Description))
	for i, item := range feed.Channel.Item {
		item.Title = fmt.Sprintf(html.UnescapeString(item.Title))
		item.Description = fmt.Sprintf(html.UnescapeString(item.Description))
		feed.Channel.Item[i] = item
	}

	printFeed(feed)

	return nil
}
