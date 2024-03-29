FROM golang:alpine as builder

# Install git + SSL ca certificates
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser
ENV USER=appuser
ENV UID=10001

# See https://stackoverflow.com/a/55757473/12429735
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /opt/trading-ai/

# use modules
COPY go.mod .

RUN git config --global url."https://randi.apriansyah:n1c9nvyc47ZJe7UYD_Ji@andromeda.ottopay.id".insteadOf "https://andromeda.ottopay.id"

ENV GO111MODULE=on
RUN go env -w GOPRIVATE=andromeda.ottopay.id/*
RUN go mod download
RUN go mod verify

COPY . .
COPY config.production.yml config.yml

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags='-w -s -extldflags "-static"' -a \
      -o /go/bin/trading-ai .


############################
# STEP 2 build a small image
############################
FROM go:alpine

# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable
COPY --from=builder /go/bin/trading-ai .
# Copy .env
COPY --from=builder /opt/trading-ai/config.yml .
COPY --from=builder /opt/trading-ai/config.production.yml .
COPY --from=builder /opt/trading-ai/config.staging.yml .

ENV CONFIGOR_ENV=staging


# Use an unprivileged user.
USER appuser:appuser

# Run the scraper binary.
ENTRYPOINT ["/trading-ai"]
