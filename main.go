package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func getTurnOffQuery(tweet string) string {
	return fmt.Sprintf("UPDATE tweets SET not_yet = FALSE WHERE tweet = '%s'", tweet)
}

func getNextTweet(db *sql.DB) (tweet string, err error) {
	const SELECT_QUERY = "SELECT tweet FROM tweets WHERE not_yet ORDER BY random() LIMIT 1"
	var rows *sql.Rows
	rows, err = db.Query(SELECT_QUERY)
	if err != nil {
		return
	}
	defer rows.Close()

	if rows.Next() {
		if err = rows.Scan(&tweet); err != nil {
			return
		}

		if _, err = db.Query(getTurnOffQuery(tweet)); err != nil {
			return
		}

		return
	}

	if _, err = db.Query("UPDATE tweets SET not_yet = TRUE"); err != nil {
		return
	}

	rows, err = db.Query(SELECT_QUERY)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		// var tweet string
		if err = rows.Scan(&tweet); err != nil {
			return
		}
		if _, err = db.Query(getTurnOffQuery(tweet)); err != nil {
			return
		}

		return
	}

	err = errors.New("There is no tweet.")
	return
}

func main() {
	consumer_key := os.Getenv("CONSUMER_KEY")
	consumer_secret := os.Getenv("CONSUMER_SECRET")
	access_token := os.Getenv("ACCESS_TOKEN")
	access_token_secret := os.Getenv("ACCESS_TOKEN_SECRET")

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	tweet, err := getNextTweet(db)
	if err != nil {
		log.Fatal(err)
	}

	config := oauth1.NewConfig(consumer_key, consumer_secret)
	token := oauth1.NewToken(access_token, access_token_secret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	if _, _, err := client.Statuses.Update(tweet, nil); err != nil {
		log.Fatal(err)
	}
}
