# Url shortener

Simple url shortener implementation in go

## Getting started

Requirements:
    - make
    - sqlite3

**Rename env.template**

``` shell
    mv ./env.template ./.env
```

**Install dependencies**

``` shell
    make install-dependencies
```

**Prepare database**

``` shell
    make migrate
    make sql-gen
```
**Start application**

``` shell
    make build
    make run
```

# Notes

[Project ideas]("https://www.geeksforgeeks.org/go-language/golang-project-ideas/")

## Task

> URL shortening services provide a solution for transforming lengthy 
> URLs into simplified, easy-to-recall ones. The project typically 
> utilizes Go's HTTP package for handling HTTP requests and responses, 
> alongside a database system like SQLite or PostgreSQL.
> This undertaking involves the development of a web server with 
> tailored HTTP responses, database interactions, HTML templating, 
> and background task handling. Additionally, it necessitates robust 
> performance capabilities to manage surges in traffic, particularly 
> if a particular URL gains traction.

## Server

As a server I decided to take **gorilla** web toolkit because it is
lightweight and simple.

[Mux for routing]("https://github.com/gorilla/mux")

## Database
I took **Sqlite3** because it is lightweight and has a lot of guides.

[sqlite3]("https://sqlite.org/index.html")

## SQL compiler and migrations
I took **sqlc** because it is simple way to communicate with database.

[sqlc]("https://github.com/sqlc-dev/sqlc")

