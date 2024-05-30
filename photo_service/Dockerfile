# Compile protos and build binary
FROM golang:alpine as build

WORKDIR /go/src/photo_service
COPY . .
 
# Install dependencies
RUN go get -d -v ./pkg/app
RUN go install -v ./pkg/app
RUN go build -o photo_service ./pkg/app/main.go

# Run
FROM alpine

WORKDIR /go/src/photo_service

COPY --from=build /go/src/photo_service/ .

ENTRYPOINT [ "./photo_service" ]
