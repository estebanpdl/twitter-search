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

// Account data struct. It defines the keys to be collected.
type AccountData struct {
	Avatar 			string
	Banner 			string
	Biography 		string
	Birthday 		string
	FollowersCount 	int
	FollowingCount 	int
	FriendsCount 	int
	IsPrivate 		bool
	IsVerified 		bool
	Joined 			*time.Time
	LikesCount 		int
	ListedCount 	int
	Location 		string
	Name 			string
	PinnedTweetIDs 	[]string
	TweetsCount 	int
	URL 			string
	UserID 			string
	Username 		string
	Website 		string
}

// arguments
var (
	username = kingpin.Flag("username", "Twitter username").Short('u').Required().String()
	password = kingpin.Flag("password", "Twitter password").Short('p').Required().String()
	accountsFile = kingpin.Flag("batch-file", "Path to the batch file containing Twitter account names").Short('b').Required().String()
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

	// Read file with user accounts
	fileBytes, err := ioutil.ReadFile(*accountsFile)
	if err != nil {
		panic(err)
	}

	// Split file into lines
	accountsList := strings.Split(string(fileBytes), "\n")
	accountsData := []AccountData{}

	// defer function
	defer func() {
		// write json file
		accountJSON, err := json.MarshalIndent(accountsData, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling JSON: ", err)
		} else {
			err = ioutil.WriteFile(*output, accountJSON, 0644)
			if err != nil {
				fmt.Println("Error writing file: ", err)
			}
		}
	}()

	for _, accountName := range accountsList {
		// Remove leading and trailing whitespace
		accountName = strings.TrimSpace(accountName)

		// Skip empty lines
		if accountName == "" {
			continue
		}

		// retry mechanism
		retryCount := 0
		maxRetry := 5

		for retryCount < maxRetry {
			// GetProfile function from scraper
			profile, err := scraper.GetProfile(accountName)
			if err != nil {
				if strings.Contains(err.Error(), "429 Too Many Requests") {
					log.Println("Rate limit exceeded. Retrying after 300 seconds...")

					// collected accounts
					fmt.Println("Accounts collected: ", len(accountsData))

					// handle rate limit exceeded
					time.Sleep(300 * time.Second)
					retryCount++
				} else {
					log.Println(err)
					break
				}
			} else {
				// Success, process the account

				// struct
				accountData := AccountData {
					Avatar: profile.Avatar,
					Banner: profile.Banner,
					Biography: profile.Biography,
					Birthday: profile.Birthday,
					FollowersCount: profile.FollowersCount,
					FollowingCount: profile.FollowingCount,
					FriendsCount: profile.FriendsCount,
					IsPrivate: profile.IsPrivate,
					IsVerified: profile.IsVerified,
					Joined: profile.Joined,
					LikesCount: profile.LikesCount,
					ListedCount: profile.ListedCount,
					Location: profile.Location,
					Name: profile.Name,
					PinnedTweetIDs: profile.PinnedTweetIDs,
					TweetsCount: profile.TweetsCount,
					URL: profile.URL,
					UserID: profile.UserID,
					Username: profile.Username,
					Website: profile.Website,
				}

				accountsData = append(accountsData, accountData)
				break
			}
		}
	}
}
