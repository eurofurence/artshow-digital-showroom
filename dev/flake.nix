# Activate the flake by using nix. This creates a reproducible development environment.
{
  description = "Flake for Artshow Showroom development";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    inputs@{
      nixpkgs,
      flake-utils,
      ...
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go
            air # hot reloading go server
            golangci-lint # fmt and lint go code
            ffmpeg # generate preview videos
            imagemagick # generate images
            mpv # playback
          ];

          shellHook =
            let
              red_text = ''\e[0;31m'';
              green_text = ''\e[0;32m'';
              reset_formatting = ''\e[0m'';
            in
            ''
              echo -e "\
              ${red_text}go version: ${green_text}$(go version)
              ${red_text}air version: ${green_text}$(air -v)
              ${red_text}golangci-lint version: ${green_text}$(golangci-lint version)
              ${red_text}ffmpeg version: ${green_text}$(ffmpeg -version)
              ${red_text}magick version: ${green_text}$(magick -version)
              ${red_text}mpv version: ${green_text}$(mpv --version)
              ${reset_formatting}"'';
        };
      }
    );
}
