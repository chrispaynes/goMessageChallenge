# goMessageChallenge
Email message parsing using Go and Angular 5

## Requirements
  - [Docker](https://www.docker.com/) and [Docker-Compose](https://docs.docker.com/compose/install/)

## Quickstart
  - The API runs on Port 3000
  - The UI runs on Port 8080 using Docker and Port 4200 using `ng serve`

#### Start UI, API and Nginx services using Docker-Compose
  - Run `make dockerUp`

#### Start API Without Docker
  - Run `make start`
  - *note: this will automatically unbind any main.go files bound to Port 3000

#### Start UI Without Docker
  - Run `make ui-dev`
  - *note: this will automatically unbind any ng files bound to Port 4200

## API
  - If you have a swagger client, such as [go-swagger](https://github.com/go-swagger/go-swagger) installed at /usr/bin/swagger, run `make swagger-serve` to view swagger documentation. Otherwise there's a simple swagger.json file with 2 API endpoints.

## Usage
  1. Start the microservices using `make dockerUp`
  2. navigate to localhost:8080
  3. Click compose to create a new email message
  4. Paste the entire email message contents into the message body and click send.
  5. Or upload a raw email file and click upload
  6. If the submission is sucessfull, steps 4 and 5 will redirect to the newly submitted message
  7. View a list of email messages by clicking the inbox on the left navigation.

## Test
- Run `make test` for API tests
- Run `make coverage` for API test coverage

## Todo (in no particular order)
- Add a datastore for the parsed email messages
- Build a more robust email client that allows user to perform CRUD operations on emails
- Increase API test coverage
- Incorporate all forms of sanitized CSS styling on the page. (Currently does not support emails with `<style>...</style>` formatting)
- POST emails using Base64 encoded messages
- Fix broken UI tests
- Add more testing for parsing multipart email messages
- Add more descriptive api messages
