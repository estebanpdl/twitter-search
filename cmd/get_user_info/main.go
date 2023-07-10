package main

import (
	"os"
    "fmt"
	"time"
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
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./main <username> <password> <account>")
		os.Exit(1)
	}

	// get arguments
	username := os.Args[1]
	password := os.Args[2]
	account := os.Args[3]

	// twitter scraper
	scraper := twitterscraper.New()

	// login process < required >
	err := scraper.Login(username, password)
	if err != nil {
		panic(err)
	}

	profile, err := scraper.GetProfile(account)
	if err != nil {
        panic(err)
    }

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
	
	accountJSON, err := json.MarshalIndent(accountData, "", "  ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("../../../test.json", accountJSON, 0644)
	if err != nil {
		panic(err)
	}
}
