# cycle-twitter-bot

The twitter bot repeats given tweets randomly.

## SQL Setup

This bot fetches tweets from `PostgreSQL`.
First of all, you need to store tweets into a table in the SQL.

```console
$ pgsql -c "CREATE tweets (tweet text, not_yet boolean)"
$ pgsql -c "INSERT INTO tweets VALUES ('$tweet1', TRUE)"
$ pgsql -c "INSERT INTO tweets VALUES ('$tweet2', TRUE)"
$ # ...
$ # you can also copy tweets from a file (e.g. 'tweets.csv')
$ # please refer the documentation of PostgreSQL or some others.
```

## Environment Variables

This bot get some values from the environment variables.
You need to set the below variables before running this.

| variable            | value                                       |
|---------------------|---------------------------------------------|
| CONSUMER_KEY        | Consumer API key of your twitter app        |
| CONSUMER_SECRET     | Consumer API secret key of your twitter app |
| ACCESS_TOKEN        | Access token of your account                |
| ACCESS_TOKEN_SECRET | Access token secret of your account         |
| DATABASE_URL        | URL to your `PostgreSQL` database           |

## Run

```console
$ go build
$ env # Ensure environemnt variables are correctly set before running
...
CONSUMER_KEY=...
CONSUMER_SECRET_KEY=...
ACCESS_TOKEN=...
ACCESS_TOKEN_SECRET=...
DATABASE_URL=...
...
$ ./cycle-twitter-bot
```
