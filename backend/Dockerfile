FROM golang:1.11.4-alpine3.8

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN apk add --update --no-cache ca-certificates git
RUN go mod download
# COPY the source code as the last step
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main

CMD ["/app/main"]
