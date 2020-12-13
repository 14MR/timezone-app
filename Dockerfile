FROM golang:1.15.5-alpine as builder
WORKDIR /app
ENV USER=tzapp
ENV GROUP=tzapp

RUN apk --no-cache add tzdata
RUN addgroup -S $GROUP && adduser -S $USER -G $GROUP
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main

FROM scratch
WORKDIR /app

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
USER tzapp:tzapp

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/main /app/main

ENTRYPOINT ["/app/main"]