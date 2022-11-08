FROM golang:latest

WORKDIR /app
 
COPY . .
RUN go mod download
RUN go install github.com/rubenv/sql-migrate/...@latest

EXPOSE 8080

CMD ["go", "run", "main.go"]