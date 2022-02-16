FROM golang:1.15-alpine as builder
ENV CGO_ENABLED=0
RUN mkdir /build
ADD ./ /build/
WORKDIR /build
RUN go build -o app main.go
FROM golang:1.15-alpine
RUN adduser -S -D -H -h /app app
USER app
COPY --from=builder /build/app /app/
WORKDIR /app
ENTRYPOINT ["./app"]