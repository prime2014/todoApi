# FROM golang:1.24.3-bookworm

# RUN apt-get update -y && apt-get install -y netcat-openbsd && rm -rf /var/lib/apt/lists/*


# # Create a non-root user with home directory
# RUN addgroup --system prime && \
#     adduser --system --ingroup prime --home /todoApi prime

# WORKDIR /todoApi

# RUN chown -R prime:prime /todoApi/

# USER prime

# COPY --chown=prime:prime ./go.mod ./go.sum ./

# RUN go mod download 

# COPY --chown=prime:prime . .

# RUN chmod +x /todoApi/entrypoint.sh \
#     && go build -o todoApi .

# EXPOSE 8080

# ENTRYPOINT ["/todoApi/entrypoint.sh"]

# CMD ["./todoApi"]

# --- Stage 1: Builder ---
FROM golang:1.24.3-bookworm AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app
COPY . .

# Build the Go binary statically
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o todoApi .

# --- Stage 2: Minimal Final Image ---
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /todoApi

COPY --from=builder /app/todoApi .

USER nonroot

EXPOSE 8080

ENTRYPOINT ["/todoApi/todoApi"]
