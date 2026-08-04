# Configuring the application

Adjust the file [config.toml](../config.toml) as needed.


# Running the server

## Running from source

Install `go` in a compatible version. See the version in [go.mod](../go.mod)

Download the repository and place your Media files:

    artshow-digital-showroom/
    ├── cmd/
    ├── docs/
    ├── README.md
    ├── ...
    └── Media/    <-- create this
        ├── config.toml
        ├── video1.mp4
        ├── ...
        └── other-video.webp

Open a terminal in the repo and run `go run cmd/main.go`

## Running from a binary

Get a (statically linked) binary for the server, e.g. a binary called `start-server`.

The static files are bundled into the binary.
Ensure that the other files are available:

    your-working-directory/
    ├── start-server
    └── Media/
        ├── config.toml
        ├── video1.mp4
        ├── ...
        └── other-video.webp

Run the binary `start-server`.

# Running the playback mpv server

Set up mpv. Run

    mpv \
    --idle=yes \
    --force-window=yes \
    --fullscreen \
    --input-ipc-server=/tmp/mpv.sock
