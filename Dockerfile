FROM golang:latest

WORKDIR /app
 
COPY . .
RUN go mod download
RUN go install github.com/rubenv/sql-migrate/...@latest
RUN go install github.com/cosmtrek/air@latest

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]