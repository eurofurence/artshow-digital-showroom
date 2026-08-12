# Setup

## Dependencies

Make sure the following programs are installed in a recent version

- `mpv` - for media playback
- `go` - for running from source, otherwise ask for a compiled binary
- `ffmpeg` - for video preview generation
- `magick` - part of imagemagick, to generate some images

## Configuring the application

The application searches for a file `config.toml` in the folder `Media`.
If this is missing it tries again under `MediaExample`.
The server starts with that configuration and searches any Videos in the same folder.

The sample folder `MediaExample` is provided.
There is however only a config file and no videos.\
Hence the application will mostly work, but only show Fallback images and display no videos.\
You can adjust the file [config.toml](../MediaExample/config.toml) as desired.

For a preconfigured setup download the folder [Media](https://cloud.eurofurence.org/index.php/f/3244131) from Nextcloud.

# Running the application

> [!IMPORTANT]
> The html server reads the `config.toml` and connects to `mpv` at launch.\
> If you make changes to the config or restart `mpv` then also restart the server.

> [!IMPORTANT]
> Thumbnails, QR-Codes and previews are placed in subfolders of `Media` / `MediaExample`
> and only generated if they are missing
> - You can overwrite them with different content and this will render after the next page reload.\
>   You might need to force clear the browser cache.
> - However when updating titles or contact info in `config.toml` the images stay outdated.\
>   Delete them to have them regenerated with the correct info on the next server launch.

## Running the playback mpv server

Run the following instruction.\
Also make sure the path listed for "input-ipc-server"
exactly matches the path `config.toml` at `playout -> mpv-socket`.

    mpv \
    --idle=yes \
    --force-window=yes \
    --fullscreen \
    --input-ipc-server=/tmp/mpv.sock

For testing you can also remove the line with `--fullscreen` to instead only display as windowed.

## Running the server

### Running from source

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

### Running from a binary

Get a (statically linked) binary for the server, e.g. a binary called `start-server`.

The static files are bundled into the binary.\
Ensure that the other files are available:

    your-working-directory/
    ├── start-server
    └── Media/
        ├── config.toml
        ├── video1.mp4
        ├── ...
        └── other-video.webp

Run the binary `./start-server`.

## Connect to the server

Open the webpage [http://localhost:8080/](http://localhost:8080/) in your browser.

If you adjust the port in the `config.toml` at `media-interface -> port` then also update the port in the URL accordingly.

For the full experience, start your browser in kiosk mode, e.g.:

    firefox -kiosk http://localhost:8080
