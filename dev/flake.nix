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
            ffmpeg # generate preview videos
            imagemagick # generate images
            mpv # playback

            pre-commit # Run checks before creating git commits
            treefmt # manage formatters
            golangci-lint # fmt and lint go code
            prettier # formatting html, css, js files
            nixfmt # format nix code
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
              ${red_text}ffmpeg version: ${green_text}$(ffmpeg -version | head -n2)
              ${red_text}magick version: ${green_text}$(magick -version | head -n1)
              ${red_text}mpv version: ${green_text}$(mpv --version | head -n1)

              ${red_text}pre-commit version: ${green_text}$(pre-commit --version)
              ${red_text}treefmt version: ${green_text}$(treefmt --version)
              ${red_text}golangci-lint version: ${green_text}$(golangci-lint version)
              ${red_text}prettier version: ${green_text}$(prettier --version)
              ${red_text}nixfmt version: ${green_text}$(nixfmt --version)
              ${reset_formatting}"'';
        };
      }
    );
}
