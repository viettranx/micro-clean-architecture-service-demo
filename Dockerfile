FROM golang:1.19-alpine as builder

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o demo_app .

FROM alpine
WORKDIR /app/
COPY --from=builder /app/demo_app .
ENTRYPOINT ["./demo_app"]