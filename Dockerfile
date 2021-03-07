FROM golang:1.15.1-alpine3.12 as builder
COPY ./ ./appp
WORKDIR ./appp
RUN go build -o /main main.go

FROM alpine:3.7
CMD ["./main"]
COPY --from=builder /main .