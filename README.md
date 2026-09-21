# Go Login CLI

Interactive command-line application for registering users and managing login sessions.

## Requirements

- Go 1.27 or newer
- Docker and Docker Compose, if you prefer to run the app in a container
- (optional) make

## Run with Docker

First build the docker image

```sh
docker compose build
```
and then run it with

```sh
docker compose run --rm app
```


## Run locally

Install the dependecies with:

```sh
go mod tidy
```
then run it with

```sh
go run ./cmd/main.go
```
Alternatively you can also use the Makefile if you have make installed:

```sh
make build
make run
```

## Commands

Once the application starts, use:

- `/register` to create a user
- `/login` to sign in
- `/whoami` to show the current user
- `/logout` to end the current session
- `/help` to show the command list
- `/clear` to clear the terminal
- `/exit` to quit

