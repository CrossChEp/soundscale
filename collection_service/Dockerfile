# Compile protos and build binary
FROM golang:alpine as build

WORKDIR /go/src/collection_service
COPY . .


# Install dependencies
RUN go get -d -v ./pkg/app
RUN go install -v ./pkg/app
RUN go build -o collection_service ./pkg/app/main.go

# Run
FROM alpine
WORKDIR /go/src/collection_service

COPY --from=build /go/src/collection_service/ .

EXPOSE 9385
ENTRYPOINT [ "./collection_service" ]