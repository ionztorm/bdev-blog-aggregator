# bdev-blog-aggregator

<!--toc:start-->

- [bdev-blog-aggregator](#bdev-blog-aggregator)
  - [Introduction](#introduction)
    - [Learning](#learning)
  - [How It Works](#how-it-works)
  - [Requirements](#requirements)
  - [Installing Requirements](#installing-requirements)
    - [Install Go](#install-go)
    - [Install PostgreSQL](#install-postgresql)
      - [Installation](#installation)
      - [Verify Installation](#verify-installation)
      - [Linux Only: Set a Password](#linux-only-set-a-password)
      - [Start the Server](#start-the-server)
      - [Connect to PostgreSQL](#connect-to-postgresql)
      - [Create a Database](#create-a-database)
  - [Create the Config File](#create-the-config-file)
  - [Installing the Application](#installing-the-application)
  - [Usage / Commands](#usage--commands)
  <!--toc:end-->

## Introduction

**bdev-blog-aggregator** is a customizable RSS feed aggregator written in Go that allows multiple users to manage their own personalized feed subscriptions on the same machine.

Users can register, add or remove their favorite RSS feeds, and browse the latest posts aggregated from those feeds. The app fetches and stores posts in a database, providing a simple command-line interface to interact with personalized content.

### Learning

This project was completed as part of the Golang backend path with [Boot.dev](https://boot.dev).

Key learnings include:

- **Writing SQL with Goose and SQLC**: Managing migrations using Goose and generating type-safe database code with SQLC.
- **Installing and running PostgreSQL locally**: Practical experience setting up a relational database and integrating it with Go.
- **Command-line interface design**: Parsing CLI arguments with `os.Args` and designing a multi-command tool.
- **Middleware concepts**: Implemented simple middleware logic for user login state and command protection.

Additionally, I introduced **directory structuring** to better organize the codebase, enhancing maintainability and scalability.

> Note: The dependency on Goose was later removed, replaced with a custom up/down migration parser embedding SQL files into the build.

## How It Works

This project includes:

- A **PostgreSQL database** to persist user accounts, feeds, and posts.
- A **Go backend** that:
  - Parses RSS feeds using XML.
  - Stores new posts while avoiding duplicates.
  - Handles user authentication and feed management.
- A **command-line interface** that acts as the primary user interface.
- A **background aggregator loop** that fetches posts from followed feeds at configurable intervals.

## Requirements

- Go (version 1.20 or later recommended)
- PostgreSQL database
- Unix-like environment (Linux, macOS) or Windows with a compatible terminal
- Network access to fetch RSS feeds

## Installing Requirements

### Install Go

Download the latest version from the [official Go installation page](https://go.dev/doc/install) and follow the platform-specific instructions.

### Install PostgreSQL

#### Installation

```bash
# macOS
brew install postgresql@15

# Linux (Debian/Ubuntu)
sudo apt update
sudo apt install postgresql postgresql-contrib
```

#### Verify Installation

```bash
psql --version
```

#### Linux Only: Set a Password for the `postgres` User

```bash
sudo passwd postgres
```

#### Start the Server

```bash
# macOS
brew services start postgresql@15

# Linux
sudo service postgresql start
```

#### Connect to PostgreSQL

```bash
# macOS
psql postgres

# Linux
sudo -u postgres psql
```

#### Create a Database

```sql
CREATE DATABASE gator;
```

## Create the Config File

```bash
touch ~/.gatorconfig.json
```

Open the file and add the following content:

```json
{
  "db_url": "<database address>"
}
```

Example database URLs:

- macOS (no password, replace `<username>` with your system username):

  ```
  postgres://<username>:@localhost:5432/gator
  ```

- Linux (replace `<username>` and `<password>` accordingly):

  ```
  postgres://<username>:<password>@localhost:5432/gator
  ```

## Installing the Application

Make sure Go is installed (see [Requirements](#requirements)).

Install the application:

```bash
go install github.com/ionztorm/bdev-blog-aggregator@latest
```

Ensure your Go binary directory is in your system PATH:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

Initialise the database tables:

```bash
gator migrate up
```

## Usage / Commands

```bash
gator help [command]             # Show help info for all commands or detailed help for a specific command
gator migrate [up|down]          # Apply or rollback SQL schema migrations
gator register <name>            # Register a new user
gator login <name>               # Log in as an existing user
gator addfeed <url>              # Add an RSS feed (automatically follows it)
gator feeds                      # List all RSS feeds added by any user
gator follow <url>               # Follow an existing feed
gator following                  # Show feeds you are currently following
gator unfollow <url>             # Unfollow a feed
gator agg <interval>             # Start the aggregator loop (e.g. 5s, 1m, 1h)
gator browse [limit]             # Browse recent posts from followed feeds (optional post limit)
gator list users                 # List all registered users
gator reset                      # Delete all users and their associated data (use with caution)
```
