# IAM

## Get Started

### Prerequisites

- [golang](https://golang.org/)
- postgresql
- [air](https://github.com/cosmtrek/air)


### Run project

##### Install module

```
cd iam
go mod download
```

##### Add `.env`

```
DB_HOST=
DB_NAME=
DB_USER=
DB_PASSWORD=
DB_PORT=
```

#### Run for Development with Live Reload

```
air
```

visit http://localhost:8080/

### Code Formatting

View files to be formatted

```
gofmt -d .
```

Format project files

```
gofmt -w .
```

## Build and Deploy

```
go build -o dist main.go
```

## Endpoints
- http://localhost:8080/oauth2/token?grant_type=client_credentials&client_id=1&client_secret=999999   (GET TOKEN)
- http://localhost:8080/api/test  (TEST AUTH with Header Authorization Bearer)