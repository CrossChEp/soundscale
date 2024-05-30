# Compile protos and build binary
FROM golang:alpine as build

WORKDIR /go/src/user_service
COPY . .

# Install dependencies
RUN go get -d -v ./pkg/app
RUN go install -v ./pkg/app
RUN go build -o user_service ./pkg/app/main.go

# Run
FROM alpine
WORKDIR /go/src/user_service

COPY --from=build /go/src/user_service/ .

EXPOSE 9385
ENTRYPOINT [ "./user_service" ]