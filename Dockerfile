# Stage 1: Build the binary
ARG GO_VERSION
FROM golang:${GO_VERSION} AS builder

WORKDIR /app
COPY cmd/ /app/cmd
COPY pkg /app/pkg
COPY go.mod /app
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o adder /app/cmd/main.go

# Stage 2: Copy to a scratch image
FROM scratch
COPY --from=builder /app/adder /adder

ENTRYPOINT ["./adder"]
