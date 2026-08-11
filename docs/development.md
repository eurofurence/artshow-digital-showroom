# Architecture

This project implements its main server in go.

Some additional user interaction on the webpage is enabled through HTMX (not yet).

# Dependencies

To run please install a compatible version of `go`.
See the version in the file [go.mod](../go.mod).

## Installing dependencies with nix

If you have `nix` running on your server then you can install dependencies with the following steps:

- Open a shell in the root of the repository
- Run `nix develop dev/`
- Dependencies are now installed in that shell

You can skip having to type `nix develop dev/` for each new shell by using `direnv`.
Run `direnv allow` in the repository.

The file `.envrc` will instruct `direnv` to enter the dev shell of the flake as long as you stay inside any folder of the repository.

# Running the server

To compile and start the server use:

    go run cmd/main.go

For hot reloading on any change use [https://github.com/air-verse/air] and just run

    air

Now the server is updated and restarted whenever you change a file.
See the config in [.air.toml](../.air.toml).

# Code style

## editorconfig

Basic indent style can be set in the file [.editorconfig](../.editorconfig).
Most IDEs and advanced editors will respect this config by default.

## go fmt

After making changes to the `go` code, use the included formatter

    go fmt ./...
