# Compile protos and build binary
FROM golang:alpine as build

WORKDIR /go/src/recommendation_service
COPY . .


# Install dependencies
RUN go get -d -v ./pkg/app
RUN go install -v ./pkg/app
RUN go build -o recommendation_service ./pkg/app/main.go

# Run
FROM alpine

WORKDIR /go/src/recommendation_service

COPY --from=build /go/src/recommendation_service/ .

ENTRYPOINT [ "./recommendation_service" ]