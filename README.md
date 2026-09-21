# Go Login CLI

Interactive command-line application for registering users and managing login sessions.

## Requirements

- Go 1.27 or newer
- Docker and Docker Compose, if you prefer to run the app in a container
- (OPTINAL) make

## Run with Docker

Build and run docker image with:
```sh
docker compose build
docker compose run --rm app
```
Alternatively you can also use the Makefile to handle builing it
```sh
make build
make run
```
## Run locally
If you have Go 1.27 or newer installed locally
First install the dependecies
```sh
go mod tidy
```
After that run these:
```sh
go run ./cmd/main.go
```

## Commands and Usage
Once the application starts, use:
- `/register` to create a user
- `/login` to sign in
- `/whoami` to show the current user
- `/logout` to end the current session
- `/help` to show the command list
- `/clear` to clear the terminal
- `/exit` to quit

