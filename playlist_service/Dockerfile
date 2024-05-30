# Compile protos and build binary
FROM golang:alpine as build

WORKDIR /go/src/playlist_service
COPY . .

# Install dependencies
RUN go get -d -v ./pkg/app
RUN go install -v ./pkg/app
RUN go build -o playlist_service ./pkg/app/main.go

# Run
FROM alpine
WORKDIR /go/src/playlist_service

COPY --from=build /go/src/playlist_service/ .

ENTRYPOINT [ "./playlist_service" ]
