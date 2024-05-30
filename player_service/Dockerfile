# Compile protos and build binary
FROM golang:alpine as build

WORKDIR /go/src/player_service
COPY . .


# Install dependencies
RUN go get -d -v .
RUN go install -v .
RUN go build -o player_service ./main.go

# Run
FROM alpine

WORKDIR /go/src/player_service
RUN apk add ffmpeg
COPY --from=build /go/src/player_service/ .

ENTRYPOINT [ "./player_service" ]
