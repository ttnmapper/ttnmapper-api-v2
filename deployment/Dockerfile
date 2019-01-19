# Setup environemtn by getting required packages
FROM golang:1.11.4 as api_environment

WORKDIR /go-modules

COPY go.mod ./go.mod

RUN go get



# Build the application in the environment
FROM api_environment as api_builder 

WORKDIR /go-modules

COPY . ./
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -v -mod=vendor -o ttnmapper_api


# Create runnable container
FROM alpine:3.8

WORKDIR /root/

COPY --from=api_builder /go-modules/ttnmapper_api .

CMD ["./ttnmapper_api"]