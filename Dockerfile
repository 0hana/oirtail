# `oirtail` GitHub Repo Fiber+Golang API Demo Dockerfile
# ------------------------------------------------------------------------------
# Multi-stage build (builder)
#
FROM golang:1.26-alpine AS builder

# - Install CA certificates for `go` app to use HTTPS
RUN apk add --no-cache ca-certificates

# - Create and `cd` to `/app/` directory
WORKDIR /app/

# - Layered Caching: Download dependencies 1st
COPY go.mod go.sum ./
RUN  go mod download

# - Layered Caching: Download source files 2nd
COPY ./ ./

# - Build (static binary)
#   + `CGO_ENABLED=0`    diables C bindings to make static binary
#   + `GOOS=linux`       tells go compiler to target linux kernel
#   + `-ldflags="-s -w"` tells go compiler to omit:
#     * symbol table
#     * DWARF debug information
#     to further slim-down production executable (and potentially obscure)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" \
 -o main cmd/server/main.go

# ------------------------------------------------------------------------------
# Multi-stage build (runner)
#
FROM scratch AS runner
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main /main
EXPOSE 80
ENTRYPOINT ["/main"]
