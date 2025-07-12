# bdev-blog-aggregator

## Introduction

A customisable RSS feed aggregator written in Go that allows multiple users to manage their own personalised feed subscriptions on the same machine.

Users can register, add or remove their favourite RSS feeds, and browse the latest posts aggregated from those feeds. The app fetches and stores posts in a database, providing a simple command-line interface to interact with personalised content.

### Learning

This project was completed as part of the Golang backend path with [Boot.dev](https://boot.dev).

This project taught me a great deal about building backend systems and working with real-world tools and data. Specifically, I learned:

- **Writing SQL with Goose and SQLC**: I became familiar with managing migrations using Goose and generating type-safe database code with SQLC.
- **Installing and running PostgreSQL locally**: Setting up the database and integrating it with Go gave me practical experience with managing a relational database.
- **Command-line interface design**: I used `os.Args` to parse CLI arguments and designed a multi-command tool with clear functionality.
- **Middleware concepts**: I introduced simple middleware logic to handle user login state and command protection.

In addition to the core curriculum from Boot.dev, I took the opportunity to **introduce directory structuring** into the project. This helped me better organise the codebase and understand how to split responsibilities across packages — a useful skill for building scalable Go applications.

## ⚙️ How It Works

This project includes:

- A **PostgreSQL database** to persist user accounts, feeds, and posts.
- A **Go backend** that:
  - Parses RSS feeds using XML.
  - Stores new posts and avoids duplicates.
  - Handles user authentication and feed management.
- A **command-line interface** that acts as the main user interface for interacting with the aggregator.
- A **background aggregator loop** that fetches posts from followed feeds at a configured interval.

## Requirements

- Go (version 1.20 or later recommended)
- PostgreSQL database
- Unix-like environment (Linux, macOS) or Windows with compatible terminal
- Network access to fetch RSS feeds from the internet

## Installing requirments

### Install Go

Download the latest binary from the [official go webpage](https://go.dev/doc/install) and follow the installation instructions.

### Install postgres

#### Install

```bash
# Mac
brew install postgresql@15

# Linux
sudo apt update
sudo apt install postgresql postgresql-contrib
```

#### Check it worked

```bash
psql --version
```

#### Linux Only: Set a Password

```bash
sudo passwd postgres
```

#### Start the server

```bash
# Mac
brew services start postgresql@15

# Linux
sudo service postgresql start
```

#### Connect

```bash
# Mac
psql postgres

# Linux
sudo -u postgres psql
```

#### Create a Database

```bash
CREATE DATABASE gator;
```

## Create the config file

```bash
touch ~/.gatorconfig.json
```

Open the config, and add:

```json
{
  "db_url": "<database address>"
}
```

The database address should look like this:

macOS (no password, your username):

```text
postgres://<username>:@localhost:5432/gator
```

Linux (password from last lesson, postgres user):

```text
postgres://<username>:<password>@localhost:5432/gator
```

## Installing the Application

Make sure you have Go installed (see [Requirements](#requirements)).

Then, install the application by running:

```bash
go install github.com/ionztorm/bdev-blog-aggregator@latest
```

Make sure your Go bin directory is in your system PATH so you can run the command from anywhere:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

## Usage / Commands

```bash
gator register <name>        # Create a user
gator login <name>           # Log in
gator addfeed <url>          # Add an RSS feed (auto-follows it)
gator feeds                  # List all feeds added by any user
gator follow <url>           # Follow a feed
gator following              # List feeds you’re currently following
gator unfollow <url>         # Unfollow a feed
gator agg <interval>         # Start the aggregator loop (interval: 5s, 1m, 1h, etc.)
gator browse <limit>         # Browse posts from followed feeds (limit is optional)
gator list users             # List all registered users
gator reset                  # Delete all users and cascade-delete related data
```
