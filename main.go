package main

import (
    "fmt"
    "os"
    "log"
    "github.com/dghubble/go-twitter/twitter"
    "github.com/dghubble/oauth1"
)

func main() {
    consumer_key := os.Getenv("CONSUMER_KEY")
    consumer_secret := os.Getenv("CONSUMER_SECRET")
    access_token := os.Getenv("ACCESS_TOKEN")
    access_token_secret := os.Getenv("ACCESS_TOKEN_SECRET")

    fmt.Println("CONSUMER_KEY:       ", consumer_key)
    fmt.Println("CONSUMER_SECRET:    ", consumer_secret)
    fmt.Println("ACCESS_TOKEN:       ", access_token)
    fmt.Println("ACCESS_TOKEN_SECRET:", access_token_secret)

    config := oauth1.NewConfig(consumer_key, consumer_secret)
    token := oauth1.NewToken(access_token, access_token_secret)
    httpClient := config.Client(oauth1.NoContext, token)

    client := twitter.NewClient(httpClient)
    _, _, err := client.Statuses.Update("test tweet", nil)
    if err != nil {
        log.Fatal(err)
    }
}
