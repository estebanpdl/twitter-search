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

## Features

Currently, this tool supports the following features:

- **Get Tweets by IDs**: Fetches tweets based on their unique IDs.
- **Get User Info**: Retrieves information for a specific Twitter account.
- **Get User Info in Batch**: Fetches information for multiple Twitter accounts at once.

## Getting Started

### **Precompiled Binaries**

For convenience, precompiled binaries are available for both Linux and Windows.

Example usage:

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

## Disclaimer

Please use this tool responsibly and in accordance with Twitter's API usage rules and guidelines. Misuse of this tool and/or violation of Twitter's rules may result in your Twitter account being temporarily or permanently suspended from accessing the API.
