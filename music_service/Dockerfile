# Compile protos and build binary
FROM golang as build

WORKDIR /go/src/music_service
COPY . .


# Install dependencies
RUN go get -d -v ./pkg/app
RUN go install -v ./pkg/app
RUN go build -o music_service ./pkg/app/main.go

# Run
FROM golang
WORKDIR /go/src/music_service

COPY --from=build /go/src/music_service/ .

ENTRYPOINT [ "./music_service" ]