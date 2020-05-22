package main

import (
	"os"
	"fmt"

	"github.com/mmcdole/gofeed"
)

func mapToUrl(feed *gofeed.Feed) map[string]string {
	output := map[string]string{}
	for _, post := range feed.Items {
		output[post.Link] = post.Title
	}
	return output
}

func readFeed(path string) *gofeed.Feed {
	fp := gofeed.NewParser()
	//remtoteFeed, _ := fpR.ParseURL("https://chollinger.com/blog/feed")
	file, _ := os.Open(path)
	defer file.Close()
	feed, _ := fp.Parse(file)
	return feed
}

func compareUrls(local, remote map[string]string) bool {
	allFound := true
	for k, _ := range remote {
		fmt.Println(k)
		if title, ok := local[k]; ok {
			fmt.Printf("Found %s\n", title)
		} else {
			fmt.Printf("Could not find %s locally, url: %s\n", title, k)
			allFound = false
		}
	}
	return allFound
}

func main() {
	// Remote
	remoteFeed := readFeed("./wordpress.xml")
	remote := mapToUrl(remoteFeed)

	// Local
	localFeed := readFeed("../public/index.xml")
	local := mapToUrl(localFeed)

	// Compare
	fmt.Println(compareUrls(local, remote))

	for k, _ := range local {
		fmt.Println(k)
	}

	fmt.Println("Remote")	
	for k, _ := range remote {
		fmt.Println(k)
	}
}
