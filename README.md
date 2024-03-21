# Toggl Unattended Programming Test

---

REST API written in Go to handle deck and cards manipulation for any game like Poker or Blackjack

### Installation

Manually Install Packages: `go mod download`

### Development

`go run cmd/server/main.go`

### Production

- Build
- Run

### Testing

Run all tests in verbose mode: `go test ./... -v`

### Postman

Postman collection available: [Postman Collection](Toggl%20Assessment.postman_collection.json)

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
