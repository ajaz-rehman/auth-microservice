#!/usr/bin/env bash

# This script is used to run database migrations using goose.

# Take up or down as the first argument

# Example:
# migrate.sh up
# migrate.sh down

# Check if goose is installed
if ! command -v goose &> /dev/null
then
    echo "goose could not be found"
    echo "Please install goose by running: go install github.com/pressly/goose/v3/cmd/goose@latest"
    exit 1
fi

# Check if the first argument is passed
if [ -z "$1" ]
then
    echo "Please provide up or down as the first argument"
    exit 1
fi

# Check if the first argument is up or down
if [ "$1" != "up" ] && [ "$1" != "down" ]
then
    echo "Please provide up or down as the first argument"
    exit 1
fi

# Check if the .env file exists and source it, else exit
if [ -f .env ]
then
    source .env
else
    echo "Please create a .env file"
    exit 1
fi

# Check if DATABASE_URL is set
if [ -z "$DATABASE_URL" ]
then
    echo "DATABASE_URL is not set in the .env file"
    exit 1
fi

# Run the migration
goose -dir sql/schema postgres "$DATABASE_URL" "$1"