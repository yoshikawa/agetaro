FROM golang:1.17-bullseye as builder

WORKDIR /go/src/github.com/yoshikawa/agetaro/server
COPY ./server .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o agetaro main.go

FROM alpine:latest
EXPOSE 8080
COPY --from=builder /go/src/github.com/yoshikawa/agetaro/server/agetaro /app/agetaro
CMD ["/app/agetaro"]
