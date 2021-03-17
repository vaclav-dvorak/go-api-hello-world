# Hello World API in golang<!-- omit in toc -->

![Go version of a Go module](https://img.shields.io/github/go-mod/go-version/vaclav-dvorak/go-api-hello-world.svg)

Simple hello world api written in golang

## 1. API

Just run this and api will start on `localhost:8080`

```bash
go run main.go
```

### 1.1. Endpoints

List of api endpoints

#### 1.1.1. Hello

Outputs simple json with hello message

URL: `/hello/{name}`

## 2. Swagger

You can run swagger with this command

```bash
make serve-swagger
```

## 3. Docker

You can build and run this in docker

```bash
docker build -t go-api-hello-world:0.1.0 .
docker run -p 8080:8080 go-api-hello-world:0.1.0
```
