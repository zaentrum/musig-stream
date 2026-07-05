# musig-stream

Audio streaming origin for the `musig` product in the stube platform. Serves
gapless Opus/HLS playback from pre-packaged media, with ffmpeg used only as a
fallback path.

## Status

**Early scaffold.** A fresh Go skeleton with a chi router, `/healthz`,
`/metrics` (Prometheus `promhttp`), and a stub `/play/*` handler is in place.
The HTTP server compiles and serves `/healthz` + `/metrics`, so the service can
be deployed as a placeholder while playback lands in a later phase.

## Design

- One stream origin per product; this service is the `musig` (audio) origin.
- Playback serves pre-packaged assets directly and reaches for on-the-fly
  transcoding only as a fallback.
- Stateless and horizontally scalable; observability via the `/metrics`
  endpoint.

## Local development

```bash
go run ./cmd/server
curl http://localhost:8080/healthz
curl http://localhost:8080/metrics
```

| Env    | Default  |
| ------ | -------- |
| `ADDR` | `:8080`  |

## Build the container

```bash
docker build -t zaentrum/musig-stream .
```

Build and push the image to your own registry and update the image reference in
the `k8s/` manifests for your environment.

## License

[MPL-2.0](LICENSE).
