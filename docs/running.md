# Running the server

## Running from source

- Install `go` in a compatible version. See [go.mod](../go.mod)
- Open a terminal in the repo and run `go run cmd/main.go`

## Running from a binary

Get a (statically linked) binary for the server, e.g. a binary called `start-server`.

The static files are bundled into the binary.
Ensure that the other files are available:

your-working-directory/
├── start-server
├── [config.toml](../config.toml)
└── media/
    ├── video1.mp4
    ├── ...
    └── other-video.webp

Run the binary `start-server`


# Configuring the server

Adjust the file [config.toml](../config.toml) as needed.
