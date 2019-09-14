# api-go
Api in Go with go-gonic, gorm and testify

## Before Usage
* Create `.env` and `.env.testing` from `.env.example`
* Launch `postgresql` and create database `DB_NAME` in `.env` or/and `.env.testing`

## Usage
* For functional test `go test -v ./test/functional/*.go`
* For unit test `go test -v ./test/unit/*.go`
* For launch server `go run main.go`

## Technology
    - Go
    - Gin-Gonic
    - gorm
    - CORS
    - Testify
    - GolintCI-Lint