# gator
RSS blog aggregator written in Go


# Initial Setup

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

## Goose

### Install 
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Create an up migration in sql/schema
```
mkdir -p sql/schema
touch sql/schema/001_users.sql
```
Add the following to 001_users.sql:
```
-- +goose Up
CREATE TABLE users(
        id UUID PRIMARY KEY,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        name TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE users;
```

### Test psql connection string
Replace the user:password with your own
```
psql "postgres://postgres:postgres@localhost:5432/gator"
```

### Run up-migration to create the users table
Replace the connection string with your own
```
goose postgres "postgres://postgres:postgres@localhost:5432/gator" up
```
Verify the table was created by running \dt from your psql shell
```
gator=# \dt
              List of relations
 Schema |       Name       | Type  |  Owner
--------+------------------+-------+----------
 public | goose_db_version | table | postgres
 public | users            | table | postgres
(2 rows)
```

# App Usage

## Commands
These are the supported commands: 

### register
Registers the `user` by adding them to the users table and setting their username in `~/.gatorconfig.json`.
```
go run . register <user>
```

### login
Sets the `user` in `~/.gatorconfig.json`. The user must already be registered for this command to work.
```
go run . login <user>
```

### addfeed
Takes a feed `title` and `url` and name. This command also automatically follows the feed.
```
go run . addfeed "<title>" "<url>"
```

### agg
Scrapes all feeds for posts every `time_between_scans` (1s, 1m, 1h, etc.)
```
go run . agg <time_between_scans>
```

### feeds
Lists all of the added feeds
```
go run . feeds
```

### follow
Accepts a feed name and url and adds it to your followed list
```
go run . follow "<title>" "<url>"
```

### following
Lists all of the feeds that you currently follow
```
go run . following
```

### unfollow
Unfollows a feed using the specified URL
```
go run . unfollow "<url>"
```

### browse
Lists posts from all feeds that you are following, takes optional `limit` argument to control how many results are returned
```
go run . browse [limit]
```
