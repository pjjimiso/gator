package main

import ( 
	"context"
	"net/http"
	"encoding/xml"
	"time"
	"io"
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

	fmt.Printf("Feed results:\n\n")

	fmt.Printf("Title: %s\n", feed.Channel.Title)
	fmt.Printf("Description: %s\n", feed.Channel.Description)
	for i, item := range feed.Channel.Item { 
		feed.Channel.Item[i].Title = fmt.Sprintf(html.UnescapeString(item.Title))
		feed.Channel.Item[i].Description = fmt.Sprintf(html.UnescapeString(item.Description))
		fmt.Printf("\tTitle: %s\n", feed.Channel.Item[i].Title)
		fmt.Printf("\tLink: %s\n", feed.Channel.Item[i].Link)
		fmt.Printf("\tDescription: %s\n", feed.Channel.Item[i].Description)
		fmt.Printf("\tPubDate: %s\n", feed.Channel.Item[i].PubDate)
		fmt.Println()
	}

	return nil
}

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
