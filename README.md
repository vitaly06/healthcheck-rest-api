# Healthcheck REST API

Simple REST API written in Go using Fiber.

This project provides basic endpoints for checking whether the service is running, responding to requests, and exposing the current application version.

---

# Tech Stack

- Go
- Fiber

---

## How to Run

Clone the repository

```sh
git clone https://github.com/vitaly06/healthcheck-rest-api
cd healthcheck-rest-api
```

Run the application:

```sh
go run ./cmd/api
```

The server will start on:

```text
http://localhost:3000
```

---

## Endpoints

| Method | Endpoint   | Description                 |
| ------ | ---------- | --------------------------- |
| GET    | `/health`  | Check service health        |
| GET    | `/ping`    | Check server response       |
| GET    | `/version` | Returns application version |

---

## Healthcheck

Request:

```sh
curl http://localhost:3000/health
```

Expected response:

```json
{
  "status": "ok"
}
```

---

## Ping

Request:

```sh
curl http://localhost:3000/ping
```

Expected response:

```json
{
  "message": "pong"
}
```

---

## Version

Request:

```sh
curl http://localhost:3000/version
```

Expected response:

```json
{
  "name": "healthcheck-rest-api",
  "status": "ok",
  "version": "1.0.0"
}
```

---

## Build

To check that the project builds correctly, run:

```sh
go build ./cmd/api
```
