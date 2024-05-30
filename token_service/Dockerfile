# Compile protos and build binary

FROM golang:alpine as build

WORKDIR /go/src/token_service
COPY . .


# Install dependencies
RUN go get -d -v .
RUN go install -v .
RUN go build -o token_service ./main.go

# Run
FROM alpine
WORKDIR /go/src/token_service

COPY --from=build /go/src/token_service/ .
COPY --from=build /go/src/token_service/config .
EXPOSE 8794
ENTRYPOINT [ "./token_service" ]
