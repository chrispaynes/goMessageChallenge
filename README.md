# goMessageChallenge
Email message parsing using Go and Angular 5

## Requirements
  - [Docker](https://www.docker.com/) and [Docker-Compose](https://docs.docker.com/compose/install/)

## Quickstart
  - The API runs on Port 3000
  - The UI runs on Port 8080 using Docker and Port 4200 using `ng serve`

#### Start UI, API and Nginx services Docker-Compose
  - Run `make dockerUp`

#### Start API Without Docker
  - Run `make start`
  - *note: this will automatically unbind any main.go files bound to Port 3000

#### Start UI Without Docker
  - Run `make ui-dev`
  - *note: this will automatically unbind any ng files bound to Port 4200

## API

## Usage

## Test
- Run `make test` for API tests
- Run `make coverage` for API test coverage
