# Compile protos and build binary
FROM golang:alpine as build

WORKDIR /go/src/gateway
COPY . .

# Install dependencies
RUN go get -d -v ./pkg/app
RUN go install -v ./pkg/app
RUN go build -o gateway ./cmd/main.go

# Run
FROM alpine
WORKDIR /go/src/gateway

COPY --from=build /go/src/gateway .

ENTRYPOINT [ "./gateway" ]
EXPOSE 9999
