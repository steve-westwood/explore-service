FROM golang
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o explore main.go
EXPOSE 9000
CMD ["/app/explore"]