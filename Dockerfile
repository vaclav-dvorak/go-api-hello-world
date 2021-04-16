# Start by building the application.
FROM golang:1.16-stretch as build

WORKDIR /go/src/app
COPY . /go/src/app

RUN go get -d -v ./...

# Build package with statically linked dependencies
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/app

# Now copy it into our base image.
FROM gcr.io/distroless/static-debian9:latest-amd64
COPY --from=build /go/bin/app /
EXPOSE 8080
CMD ["/app"]
