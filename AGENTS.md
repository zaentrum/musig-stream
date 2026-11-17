# musig-stream

musig-stream is a scaffold of the stream origin for musig, the platform's planned music
app, which is to serve its audio as HLS. Today it is a Go service with a chi router that
answers `/healthz`, Prometheus metrics at `/metrics`, and 501 Not Implemented on
`/play/*`. No app and no other service calls it, and the platform's chart does not
run it.
How a change is proposed and what a pull request has to pass is in the
[guide to contributing](https://github.com/zaentrum/.github/blob/main/CONTRIBUTING.md);
what may be built is in [building from the specification](https://github.com/zaentrum/zaentrum/blob/main/AGENTS.md).

## Build

```sh
go build ./...
docker build -t zaentrum/musig-stream .
```

Go 1.23.0 (`go.mod`); its `toolchain go1.24.5` line has an older Go fetch and use
go1.24.5. The image is the static binary on distroless, and
`--build-arg BASE=<registry prefix>` takes the Go image from a mirror instead of Docker
Hub. `image.yml` builds the image on a push to `main` and publishes it as
`ghcr.io/zaentrum/musig-stream:latest` and `:<commit>`.

## Test

As `build.yml` runs them on a pull request and on `main`:

```sh
go build ./...
go vet ./...
go test ./...
```

The test needs Go alone: it starts the server on a free port of `127.0.0.1`.

## Run it

```sh
go run ./cmd/server
curl http://localhost:8080/healthz
```

It reads `ADDR` alone (default `:8080`). Nothing on the reference demo calls it, and
there are no recorded answers to run it against.

## What the tests leave unchecked

- `cmd/server/main_test.go` starts the server, checks that `/healthz` answers 200, and
  that it stops on SIGTERM. Nothing more has a test.
- `/metrics` and the `/play/*` placeholder.
- The image: no workflow builds it on a pull request.
- The manifests in `k8s/`: nothing applies or checks them.
