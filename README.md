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

after running, create sample data for oauth2_client
```
INSERT INTO oauth2_client (id,created_at,updated_at,deleted_at,secret,"domain","data") VALUES
	 (1,'2021-12-21 22:47:48.060504+07','2021-12-21 22:47:48.060504+07',NULL,'999999','http://localhost','{"ClientID":"1","UserID":"","RedirectURI":"","Scope":"","Code":"","CodeCreateAt":"0001-01-01T00:00:00Z","CodeExpiresIn":0,"Access":"_S0OOXJUPZOGTFAA__SPIG","AccessCreateAt":"2021-12-21T20:03:26.427334+07:00","AccessExpiresIn":7200000000000,"Refresh":"","RefreshCreateAt":"0001-01-01T00:00:00Z","RefreshExpiresIn":0}');
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