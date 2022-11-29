#Build stage
FROM golang:1.19.3-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o deliveroo deliverooserver.go

#Run stage

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/deliveroo .
COPY start.sh .
COPY wait-for.sh .
EXPOSE 3000
CMD ["/app/deliveroo"]
ENTRYPOINT [ "/app/start.sh" ]