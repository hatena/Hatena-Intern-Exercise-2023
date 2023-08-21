FROM golang:1.21-bookworm as builder

WORKDIR /src/

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build


FROM debian:bookworm-slim

RUN groupadd -g 1000 grpc && useradd -u 1000 -g grpc grpc
USER grpc

COPY --from=builder /src/bin/server /bin/server

CMD ["/bin/server"]
