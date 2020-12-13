FROM golang:1.15.5-alpine as builder
WORKDIR /app
RUN apk --no-cache add tzdata
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main

FROM scratch
WORKDIR /app
COPY --from=builder /app/main /app/main
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENTRYPOINT ["/app/main"]