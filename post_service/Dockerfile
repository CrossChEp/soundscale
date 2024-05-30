# Compile protos and build binary
FROM golang:alpine as build

WORKDIR /go/src/post_service
COPY . .


# Install dependencies
RUN go get -d -v ./pkg/app
RUN go install -v ./pkg/app
RUN go build -o post_service ./pkg/app/main.go

# Run
FROM alpine
WORKDIR /go/src/post_service

COPY --from=build /go/src/post_service/ .

EXPOSE 9385
ENTRYPOINT [ "./post_service" ]