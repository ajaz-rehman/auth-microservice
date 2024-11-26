# Auth Microservice

## Description

This project provides all authentication backend functionality and relevant endpoints. It also handles access and refresh tokens. Additionally, users can authenticate using Google, Microsoft, and GitHub accounts.

## Why?

I've created a lot of projects from scratch and authentication has always been a primary element of pretty much every commercial API I've built. Rather than developing custom auth every time or using existing libraries, I wanted to have an easy way of adding solid authentication to my projects. That's why I decided to build this auth microservice that I can easily use with any future project.

## Prerequisites

-   [Go 1.23+](https://go.dev)
-   [PostgreSQL 13+](https://postgresql.org)

## Installation

Clone the repository and navigate to the project directory, then follow the steps below:

### Create a `.env` file:

```dosini
PORT=5000
GO_ENV=development
DATABASE_URL=postgres://admin:password@localhost:5432/test?sslmode=disable
```

Replace the `DATABASE_URL` value with your own PostgreSQL connection string.

### Install Goose:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Install SQLC:

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### Install dependencies:

```bash
go mod download
```

### Run the migrate script:

```bash
scripts/migrate.sh up
```

## Usage

### To start the server locally:

```bash
go run .
```
