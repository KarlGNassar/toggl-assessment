## SETUP FOR DOCKER

run (in the root directory /): `docker-compose up -d --build`

> Note: for testing, add `x-api-key` with value of `123erxn87` in the headers

# Toggl Unattended Programming Test

REST API written in Go to handle deck and cards manipulation for any game like Poker or Blackjack

Features:

- REST API w/ Custom Response
- Dependency Injection
- Error Handling & Custom Error responses
- Middleware Authentication (using x-api-key in header)

Tech Stack: Golang (w/ chi router), MongoDb

### Installation

Manually Install Packages: `go mod download`

> Note: Connection strings, api keys and other config variables are purposely not hidden and not ignored (in .gitignore) for the program to run out of the box

### Development

`go run cmd/server/main.go`

> Important: If you're having issues running the `main.go` because of the `.env` or log files, try changing the path to absolute paths in the `main.go` i.e. `.env` to `../../.env` and `logs/deck-api.log` to `../../logs/deck-api.log`

### Production

- Build: `go build cmd/server/main.go`
- Run:
  - Windows: `.\main.exe`
  - Linux: `./main`

### Testing

Run all tests in verbose mode: `go test ./... -v`

### Postman

Postman collection available: [Postman Collection](Toggl%20Assessment.postman_collection.json)

> Examples included which you can run as is

---

### GCP Architecture

![GCP Architecture](/assets/architecture.png)

> Written using eraser.io diagram as code tool.
> My Linkedin post talking about it [LinkedIn Post](https://www.linkedin.com/posts/karl-nassar_transform-your-ideas-into-appealing-diagrams-activity-7176533008865239040-aeNi?utm_source=share&utm_medium=member_desktop)

#### Deployment Setup:

- Monitoring via file logs (rotate logs each 10 days using logrotate)
  - logrotate config:
    ```bash
    /toggl-assesment/logs/\*.log {
        daily
        rotate 10
        maxsize 20M
        copytruncate
        missingok
        notifempty
        compress
        delaycompress
        sharedscripts
    }
    ```
- Optional: Add load balancers + auto scaling groups
