package main

import (
	"encoding/csv"
	"os"

	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
)

type InstaInfluencers struct {
	rank       string
	influencer string
	category   string
	followers  string
	country    string
	engAuth    string
	engAvg     string
}

func main() {
	logger := logrus.StandardLogger()

	var instaInfluencers []InstaInfluencers

	pageToScrape := "https://hypeauditor.com/top-instagram-all-russia/"

	c := colly.NewCollector()

	c.OnHTML(".row__top", func(e *colly.HTMLElement) {
		instaInfluencer := InstaInfluencers{}

		instaInfluencer.rank = e.ChildText(".row .row-cell.rank span[data-v-2e6a30b8]")
		instaInfluencer.influencer = e.ChildText(".contributor__title")
		instaInfluencer.category = e.ChildText(".category")
		instaInfluencer.followers = e.ChildText(".subscribers")
		instaInfluencer.country = e.ChildText(".audience")
		instaInfluencer.engAuth = e.ChildText(".authentic")
		instaInfluencer.engAvg = e.ChildText(".engagement")

		instaInfluencers = append(instaInfluencers, instaInfluencer)
	})

	err := c.Visit(pageToScrape)
	if err != nil {
		logger.Infof("c.Visit: %s", err)
	}

	file, err := os.Create("parsed_data.csv")
	if err != nil {
		logger.Warningf("os.Create: %s", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Infof("file.Close: %s", err)
		}
	}(file)

	writer := csv.NewWriter(file)

	headers := []string{
		"Rank",
		"Influencer",
		"Category",
		"Subscribers",
		"Audience",
		"Authentic",
		"Engagement",
	}

	err = writer.Write(headers)
	if err != nil {
		logger.Warningf("writer.Write(headers): %s", err)
	}

	for _, instaInfluencer := range instaInfluencers {
		record := []string{
			instaInfluencer.rank,
			instaInfluencer.influencer,
			instaInfluencer.category,
			instaInfluencer.followers,
			instaInfluencer.country,
			instaInfluencer.engAuth,
			instaInfluencer.engAvg,
		}

		err = writer.Write(record)
		if err != nil {
			logger.Warningf("writer.Write(record): %s", err)
		}
	}

	defer writer.Flush()
}
