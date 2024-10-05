FROM --platform=$BUILDPLATFORM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags '-s -w -extldflags "-static"' -o aria2

FROM --platform=$BUILDPLATFORM alpine AS final
RUN apk add --no-cache aria2
USER 1000
WORKDIR /data
WORKDIR /app
COPY --from=builder --chown=1000 /app/aria2 /usr/local/bin/aria2
CMD ["aria2"]