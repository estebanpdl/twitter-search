package main

import (
	"os"
    "fmt"
	"time"
	"strings"
	"io/ioutil"
	"encoding/json"
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

// main function
func main () {
	if len(os.Args) != 5 {
		fmt.Println("Usage: ./main <username> <password> <accounts-file> <output")
		os.Exit(1)
	}

	// get arguments
	username := os.Args[1]
	password := os.Args[2]
	accountsFile := os.Args[3]
	output := os.Args[4]

	// twitter scraper
	scraper := twitterscraper.New()

	// login process < required >
	err := scraper.Login(username, password)
	if err != nil {
		panic(err)
	}

	// Read file with user accounts
	fileBytes, err := ioutil.ReadFile(accountsFile)
	if err != nil {
		panic(err)
	}

	// Split file into lines
	accountsList := strings.Split(string(fileBytes), "\n")
	accountsData := []AccountData{}

	for _, accountName := range accountsList {
		// Remove leading and trailing whitespace
		accountName = strings.TrimSpace(accountName)

		// Skip empty lines
		if accountName == "" {
			continue
		}

		// GetProfile function from scraper
		profile, err := scraper.GetProfile(accountName)
		if err != nil {
			panic(err)
		}

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
	}

	// write json file
	accountJSON, err := json.MarshalIndent(accountsData, "", "  ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(output, accountJSON, 0644)
	if err != nil {
		panic(err)
	}
}
