# Base-image registry prefix. Empty default = public Docker Hub; a private
# deploy mirror passes --build-arg BASE=registry.example/library/ .
ARG BASE=
# syntax=docker/dockerfile:1.7
FROM ${BASE}golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod go.sum* ./
RUN go mod download || true
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/musig-stream ./cmd/server

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/musig-stream /usr/local/bin/musig-stream
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/musig-stream"]

LABEL org.opencontainers.image.source="https://github.com/zaentrum/musig-stream"
LABEL org.opencontainers.image.title="musig-stream"
