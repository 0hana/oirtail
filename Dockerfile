# `oirtail` GitHub Repo Fiber+Golang API Demo Dockerfile
# ------------------------------------------------------------------------------
# Multi-stage build (builder)
#
FROM golang:1.26-alpine AS builder

# - Create and `cd` to `/app/` directory
WORKDIR /app/

# - Layered Caching: Download dependencies 1st
COPY go.mod go.sum ./
RUN  go mod download

# - Layered Caching: Download source files 2nd
COPY ./ ./

# - Build using Standard Go Project Layout: `cmd/server/main.go`
RUN  go build -o main cmd/server/main.go

# ------------------------------------------------------------------------------
# Multi-stage build (runner)
#
FROM alpine:latest AS runner
WORKDIR /app/
COPY --from=builder /app/main ./
EXPOSE 8080
CMD ["./main"]
