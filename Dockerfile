FROM golang:1.21 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o release-notes-publisher ./main.go

FROM alpine
COPY --from=builder /app/release-notes-publisher /release-notes-publisher
ENTRYPOINT ["/release-notes-publisher"]