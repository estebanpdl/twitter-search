package main

import (
	"os"
    "fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
    twitterscraper "github.com/n0madic/twitter-scraper"
)

// SimpleTweet struct. It defines the keys to be collected.
type SimpleTweet struct {
    ConversationID 		string
	Hashtags 			[]string
	ID 					string
	InReplyToStatusID 	string
	QuotedStatusID 		string
	RetweetedStatusID 	string
	IsQuoted			bool
	IsReply				bool
	IsRetweet 			bool
	IsSelfThread 		bool
	Likes 				int
	Retweets 			int
	Replies 			int
	Views 				int
	Text 				string
	Timestamp 			int64
	URLs 				[]string
	UserID 				string
	Username 			string
}

// main function
func main () {
	if len(os.Args) != 5 {
		fmt.Println("Usage: ./main <username> <password> <ids-text-file> <output>")
		os.Exit(1)
	}

	// get arguments
	username := os.Args[1]
	password := os.Args[2]
	tweetIDsFile := os.Args[3]
	output := os.Args[4]

	// twitter scraper
	scraper := twitterscraper.New()

	// login process < required >
	err := scraper.Login(username, password)
	if err != nil {
		panic(err)
	}

	// Read file with tweet IDs
	fileBytes, err := ioutil.ReadFile(tweetIDsFile)
	if err != nil {
		panic(err)
	}

	// Split file into lines
	tweetIDs := strings.Split(string(fileBytes), "\n")
	simpleTweets := []SimpleTweet{}

	for _, tweetID := range tweetIDs {
		// Remove leading and trailing whitespace
		tweetID = strings.TrimSpace(tweetID)

		// Skip empty lines
		if tweetID == "" {
			continue
		}

		// GetTweet function from scraper
		tweet, err := scraper.GetTweet(tweetID)
		if err != nil {
			fmt.Printf("Failed to get tweet %s: %v\n", tweetID, err)
			continue
		}

		// struct
		simpleTweet := SimpleTweet {
			ConversationID: tweet.ConversationID,
			Hashtags: tweet.Hashtags,
			ID: tweet.ID,
			InReplyToStatusID: tweet.InReplyToStatusID,
			QuotedStatusID: tweet.QuotedStatusID,
			RetweetedStatusID: tweet.RetweetedStatusID,
			IsQuoted: tweet.IsQuoted,
			IsReply: tweet.IsReply,
			IsRetweet: tweet.IsRetweet,
			IsSelfThread: tweet.IsSelfThread,
			Likes: tweet.Likes,
			Retweets: tweet.Retweets,
			Replies: tweet.Replies,
			Views: tweet.Views,
			Text: tweet.Text,
			Timestamp: tweet.Timestamp,
			URLs: tweet.URLs,
			UserID: tweet.UserID,
			Username: tweet.Username,
		}

		simpleTweets = append(simpleTweets, simpleTweet)
	}

	tweetJSON, err := json.MarshalIndent(simpleTweets, "", "  ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(output, tweetJSON, 0644)
	if err != nil {
		panic(err)
	}
}
