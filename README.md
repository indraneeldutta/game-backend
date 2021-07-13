# Game Backend

A backend API service to serve for game. Built using Golang and Mongo as database.

## Pre-requisites

1. This application requires mongoDB to run.
2. Change the credentials for DB in `resources/config.yml` 

## Usage

Run the following commands to start the application

```
go mod vendor

go build

./game-backend
```