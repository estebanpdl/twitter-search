package main

import (
    "fmt"
	"log"
	"time"
	"strings"
	"io/ioutil"
	"encoding/json"
	kingpin "github.com/alecthomas/kingpin/v2"
    twitterscraper "github.com/n0madic/twitter-scraper"
)

type RetweetedStatus struct {
	ConversationID  		string
	ID 						string
	HTML					string
    Username 				string
    Name 					string
    UserID 					string
	IsPin 					bool
	IsQuoted				bool
	IsReply					bool
	IsRetweet 				bool
	IsSelfThread 			bool
	Likes 					int
	Retweets 				int
	Replies 				int
	Views 					int
	Text 					string
	Timestamp 				int64
	TimeParsed 				time.Time
}

type QuotedStatus struct {
	ConversationID  		string
	ID 						string
	HTML					string
    Username 				string
    Name 					string
    UserID 					string
	IsPin 					bool
	IsQuoted				bool
	IsReply					bool
	IsRetweet 				bool
	IsSelfThread 			bool
	Likes 					int
	Retweets 				int
	Replies 				int
	Views 					int
	Text 					string
	Timestamp 				int64
	TimeParsed 				time.Time
}

// SimpleTweet struct. It defines the keys to be collected.
type SimpleTweet struct {
    ConversationID 		string
	Hashtags 			[]string
	HTML				string
	ID 					string
	InReplyToStatusID 	string
	QuotedStatusID 		string
	RetweetedStatusID 	string
	IsPin 				bool
	IsQuoted			bool
	IsReply				bool
	IsRetweet 			bool
	IsSelfThread 		bool
	RetweetedStatus 	RetweetedStatus
	QuotedStatus 		QuotedStatus
	Likes 				int
	Retweets 			int
	Replies 			int
	Views 				int
	Text 				string
	Timestamp 			int64
	TimeParsed 			time.Time
	URLs 				[]string
	UserID 				string
	Username 			string
	Name 				string
}

// arguments
var (
	username = kingpin.Flag("username", "Twitter username").Short('u').Required().String()
	password = kingpin.Flag("password", "Twitter password").Short('p').Required().String()
	tweetIDsFile = kingpin.Flag("batch-file", "Path to the batch file containing tweets ids").Short('b').Required().String()
	output = kingpin.Flag("output", "Path to the output file").Short('o').Required().String()
)

// main function
func main () {
	// Help support -> Parse arguments < kingpin >
	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.Parse()

	// twitter scraper
	scraper := twitterscraper.New()

	// login process < required >
	err := scraper.Login(*username, *password)
	if err != nil {
		panic(err)
	}

	// Read file with tweet IDs
	fileBytes, err := ioutil.ReadFile(*tweetIDsFile)
	if err != nil {
		panic(err)
	}

	// Split file into lines
	tweetIDs := strings.Split(string(fileBytes), "\n")
	simpleTweets := []SimpleTweet{}


	// defer function
	defer func() {
		// write json file
		tweetJSON, err := json.MarshalIndent(simpleTweets, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling JSON: ", err)
		} else {
			err = ioutil.WriteFile(*output, tweetJSON, 0644)
			if err != nil {
				fmt.Println("Error writing file: ", err)
			}
		}
	}()

	for _, tweetID := range tweetIDs {
		// remove leading and trailing whitespace
		tweetID = strings.TrimSpace(tweetID)

		// skip empty lines
		if tweetID == "" {
			continue
		}

		// retry mechanism
		retryCount := 0
		maxRetry := 5

		for retryCount < maxRetry {
			// GetTweet function from scraper
			tweet, err := scraper.GetTweet(tweetID)
			if err != nil {
				if strings.Contains(err.Error(), "429 Too Many Requests") {
					log.Println("Rate limit exceeded. Retrying after 300 seconds...")

					// collected tweets
					fmt.Println("Tweets collected: ", len(simpleTweets))

					// handle rate limit exceeded
					time.Sleep(300 * time.Second)
					retryCount++
				} else {
					log.Println(err)
					break
				}
			} else {
				// Success, process the tweet
				
				// retweeted status struct
				retweetedStatus := RetweetedStatus{}
				if tweet.RetweetedStatus != nil {
					retweetedStatus = RetweetedStatus {
						ConversationID: tweet.RetweetedStatus.ConversationID,
						ID: tweet.RetweetedStatus.ID,
						HTML: tweet.RetweetedStatus.HTML,
						Username: tweet.RetweetedStatus.Username,
						Name: tweet.RetweetedStatus.Name,
						UserID: tweet.RetweetedStatus.UserID,
						IsPin: tweet.RetweetedStatus.IsPin,
						IsQuoted: tweet.RetweetedStatus.IsQuoted,
						IsReply: tweet.RetweetedStatus.IsReply,
						IsRetweet: tweet.RetweetedStatus.IsRetweet,
						IsSelfThread: tweet.RetweetedStatus.IsSelfThread,
						Likes: tweet.RetweetedStatus.Likes,
						Retweets: tweet.RetweetedStatus.Retweets,
						Replies: tweet.RetweetedStatus.Replies,
						Views: tweet.RetweetedStatus.Views,
						Text: tweet.RetweetedStatus.Text,
						Timestamp: tweet.RetweetedStatus.Timestamp,
						TimeParsed: tweet.RetweetedStatus.TimeParsed,
					}
				}

				// quoted status struct
				quotedStatus := QuotedStatus{}
				if tweet.QuotedStatus != nil {
					quotedStatus = QuotedStatus {
						ConversationID: tweet.QuotedStatus.ConversationID,
						ID: tweet.QuotedStatus.ID,
						HTML: tweet.QuotedStatus.HTML,
						Username: tweet.QuotedStatus.Username,
						Name: tweet.QuotedStatus.Name,
						UserID: tweet.QuotedStatus.UserID,
						IsPin: tweet.QuotedStatus.IsPin,
						IsQuoted: tweet.QuotedStatus.IsQuoted,
						IsReply: tweet.QuotedStatus.IsReply,
						IsRetweet: tweet.QuotedStatus.IsRetweet,
						IsSelfThread: tweet.QuotedStatus.IsSelfThread,
						Likes: tweet.QuotedStatus.Likes,
						Retweets: tweet.QuotedStatus.Retweets,
						Replies: tweet.QuotedStatus.Replies,
						Views: tweet.QuotedStatus.Views,
						Text: tweet.QuotedStatus.Text,
						Timestamp: tweet.QuotedStatus.Timestamp,
						TimeParsed: tweet.QuotedStatus.TimeParsed,
					}
				}

				simpleTweet := SimpleTweet {
					ConversationID: tweet.ConversationID,
					HTML: tweet.HTML,
					Hashtags: tweet.Hashtags,
					ID: tweet.ID,
					InReplyToStatusID: tweet.InReplyToStatusID,
					QuotedStatusID: tweet.QuotedStatusID,
					RetweetedStatusID: tweet.RetweetedStatusID,
					IsPin: tweet.IsPin,
					IsQuoted: tweet.IsQuoted,
					IsReply: tweet.IsReply,
					IsRetweet: tweet.IsRetweet,
					IsSelfThread: tweet.IsSelfThread,
					RetweetedStatus: retweetedStatus,
					QuotedStatus: quotedStatus,
					Likes: tweet.Likes,
					Retweets: tweet.Retweets,
					Replies: tweet.Replies,
					Views: tweet.Views,
					Text: tweet.Text,
					Timestamp: tweet.Timestamp,
					TimeParsed: tweet.TimeParsed,
					URLs: tweet.URLs,
					UserID: tweet.UserID,
					Username: tweet.Username,
					Name: tweet.Name,
				}

				simpleTweets = append(simpleTweets, simpleTweet)
				break
			}
		}		
	}
}
