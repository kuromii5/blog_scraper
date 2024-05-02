package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title string    `xml:"title"`
		Link  string    `xml:"link"`
		Desc  string    `xml:"description"`
		Lang  string    `xml:"language"`
		Item  []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	Desc    string `xml:"description"`
	PubDate string `xml:"pubDate"`
}

func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	r, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
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
