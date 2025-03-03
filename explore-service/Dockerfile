FROM golang:1.24

WORKDIR /app

# Copy the shared proto files
COPY shared/ /app/shared/

# Copy the main service
COPY explore-service/ /app/explore-service/

# Copy go.work
COPY go.work /app/

WORKDIR /app/explore-service

RUN go build -o explore cmd/main.go

EXPOSE 9000

CMD ["/app/explore-service/explore"] 