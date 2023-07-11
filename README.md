<div align="center">

# Twitter Search

---

[![GitHub forks](https://img.shields.io/github/forks/estebanpdl/twitter_search.svg?style=social&label=Fork&maxAge=2592000)](https://GitHub.com/estebanpdl/twitter_search/network/)
[![GitHub stars](https://img.shields.io/github/stars/estebanpdl/twitter_search?style=social)](https://github.com/estebanpdl/twitter_search/stargazers)
[![Open Source](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://twitter.com/estebanpdl)
[![Twitter estebanpdl](https://badgen.net/badge/icon/twitter?icon=twitter&label)](https://twitter.com/estebanpdl)

---

</div>

## Overview

Twitter Search is an OSINT (Open-Source Intelligence) tool that allows you to fetch data from Twitter via the command line. It's built using Go and leverages the twitter-scraper package to interact with Twitter's API.

This tool aims to simplify the process of gathering data from Twitter for research and data analysis purposes.

## Getting Started

### **Precompiled Binaries**

For convenience, precompiled binaries are available for both Linux and Windows.

### **Downloading and Cloning the Repository**

You can clone this repository to your local machine by using the following command in your terminal you need to have [git](https://git-scm.com/downloads) installed:

`git clone https://github.com/estebanpdl/twitter_search.git`

Alternatively, you can download a zip file of the repository from the GitHub page by clicking the green "Code" button and then clicking "Download ZIP". Once downloaded, you'll need to unzip the file to access the repository contents.

### **Navigating the Repository**

After cloning or downloading the repository, you can navigate to the directory containing the tools using your terminal or command line interface:

`cd twitter_search/cmd`

### **Features**

Currently, this tool supports the following features:

- **Get Tweets by IDs**: Fetches tweets based on their unique IDs.
- **Get User Info**: Retrieves information for a specific Twitter account.
- **Get User Info in Batch**: Fetches information for multiple Twitter accounts at once.

In this directory, you will find a directory for each tool (`get_tweets_from_ids`, `get_user_timelines`, `get_tweets_by_keywords`). Each tool's directory contains a main.go file with the Go source code, and two directories, linux/ and windows/, holding the respective executables for each operating system.

Refer to the **Directory Structure** section for more details.

### **Running the Tools**

Before running any of the tools, you must ensure that you have the correct permissions. On Unix-based systems like Linux or macOS, you might need to make the scripts executable by running `chmod +x filename` where `filename` is the name of the script.

You can then run the tools by navigating to their respective directories and executing them. 

Examples:

```
cd get_tweets_from_ids/linux
./get_tweets_from_ids
```

```
cd get_user_timelines/windows
get_user_timelines.exe
```

Refer to the usage sections for more detailed instructions on running each tool.

## **Example usage**

### **Get Tweets by IDs**

To fetch specific tweets using their unique IDs, use the `get_tweets_from_ids` command:

`./get_tweets_from_ids username password ids-text-file output_path`

Replace the following placeholders with your actual information:

- **`username`**: Your Twitter username.
- **`password`**: Your Twitter password.
- **`ids_text_file`**: A text file where each line is a unique ID of a tweet you wish to fetch.
- **`output_path`**: The path where the resulting JSON file should be written.

For example:

`./get_tweets_from_ids johndoe mypassword tweet_ids.txt /home/johndoe/twitter_data.json`

This command will fetch the tweets specified in `tweet_ids.txt` and write the resulting JSON file in the following directory: `/home/johndoe/twitter_data.json`.


### **Get User Info**

To retrieve information about a specific Twitter account, use the `get_user_info` command:

`./get_user_info username password target_username output_path`

Replace the following placeholders with your actual information:

- **`username`**: Your Twitter username.
- **`password`**: Your Twitter password.
- **`target_username`**: The username of the Twitter account for which you want to fetch information.
- **`output_path`**: The path where the resulting JSON file should be written.

For example:

`./get_user_info johndoe mypassword elonmusk /home/johndoe/twitter_data.json`

This command will fetch information about the Twitter account `@elonmusk` and write the resulting JSON file in the following directory: `/home/johndoe/twitter_data.json`.


### **Get User Info in Batch**

To retrieve information about multiple Twitter accounts at once, use the `get_user_info_batch` command:

`./get_user_info_batch username password accounts_text_file output_path`

Replace the following placeholders with your actual information:

- **`username`**: Your Twitter username.
- **`password`**: Your Twitter password.
- **`accounts_text_file`**: A text file where each line is a unique username of the Twitter accounts you wish to fetch information about.
- **`output_path`**: The path where the resulting JSON files should be written.

For example:

`./get_user_info_batch johndoe mypassword accounts.txt /home/johndoe/twitter_data.json`

This command will fetch information about the Twitter accounts specified in `accounts.txt` and write the resulting JSON file in the following directory: `/home/johndoe/twitter_data.json`.

## Authentication

Please note: Due to restrictions on Twitter's API, authentication is a must for this tool. You will need to provide your Twitter username and password to interact with the API. Your credentials are used exclusively for authenticating with the API and are not stored or used in any other way.

## Data Structure

```
twitter_search/
├─ cmd/
│  ├─ get_tweets_from_ids/
│  │  ├─ linux/
│  │  ├─ windows/
│  │  ├─ main.go
│  ├─ get_user_timelines/
│  │  ├─ linux/
│  │  ├─ windows/
│  │  ├─ main.go
│  ├─ get_tweets_by_keywords/
│  │  ├─ linux/
│  │  ├─ windows/
│  │  ├─ main.go
├─ go.mod
├─ go.sum
```

- The `cmd/` directory contains the Go source code for each command-line tool. Each tool has its own subdirectory.
- Inside each tool's subdirectory (e.g., `get_tweets_from_ids/`), there is a `main.go` file containing the Go source code, and two additional directories: `linux/` and `windows/`. These directories contain the compiled executables for Linux and Windows, respectively.
- The `go.mod` and `go.sum` files are used by Go's dependency management system. They ensure that the project uses the correct versions of its dependencies.

## Disclaimer

Please use this tool responsibly and in accordance with Twitter's API usage rules and guidelines. Misuse of this tool and/or violation of Twitter's rules may result in your Twitter account being temporarily or permanently suspended from accessing the API.
