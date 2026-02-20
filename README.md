# gator
RSS blog aggregator written in Go


# Initial Setup

## Clone the repo
```
git clone https://github.com/pjjimiso/gator.git
```

## Postgresql

### Install - Ubuntu/Debian/WSL
```
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo service postgresql start
```

### Enter psql shell
```
psql
```

### Create a Database and connect to it 
```
CREATE DATABASE gator;
\c gator
```

### Test psql connection string
Replace the user:password with your own
```
psql "postgres://postgres:postgres@localhost:5432/gator"
```

## Gatorconfig

### Create a `.gatorconfig.json` in your home directory
```
touch ~/.gatorconfig.json
```

Add a `"db_url": "<db_connection_string>"` to your config file
Example:
```
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"
}
```

## Goose

### Install 
Install the goose module
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Run goose up migrations to create the tables
Replace `<db_url>` with the url from your `.gatorconfig.json`. This command must be run from within the `schema` dir
```
cd sql/schema
goose postgres "<db_url>" up
```

# App Usage

## Commands
These are the supported commands: 

### register
Registers the `user`. This creates a new user in the database and adds them to `~/.gatorconfig.json`.
usage: `gator register <user>`
```
gator register this_is_patrick
```

### login
Update the `user` in `~/.gatorconfig.json`. The user must already be registered for this command to work. Note that the `register` command runs login, so you only need to run this when switching between different users.
usage: `gator login <user>`
```
gator login this_is_patrick
```

### addfeed
Adds a new feed to the database using the `feed_name` and `feed_url`. This command also automatically follows the feed for user currently logged in.
usage: `gator addfeed <feed_name> <feed_url>`
```
gator addfeed "Hacker News RSS" "https://hnrss.org/newest"
```

### agg
Scrapes all feeds for posts every `time_between_scans` (1s, 1m, 1h, etc.)
usage: `gator agg <time_between_scans>`
```
gator agg 5m
```

### feeds
Lists all of the added feeds
usage: `gator feeds`
```
gator feeds
```

### follow
Accepts a `feed_name` and `feed_url` and adds it to the current user's followed feeds
usage: `gator follow "<feed_name>" "<feed_url>"`
```
gator follow "Hacker News RSS" "https://hnrss.org/newest"
```

### following
Lists all of the feeds that you currently follow
usage: `gator following`
```
gator following
```

### unfollow
usage: `gator unfollow "<feed_url>"`
Unfollows a feed using the specified URL
```
gator unfollow "https://hnrss.org/newest"
```

### browse
Lists posts from all feeds that you are following, *optional* `limit` argument controls how many results are returned
usage: `gator browse [limit]`
```
gator browse 10
```

