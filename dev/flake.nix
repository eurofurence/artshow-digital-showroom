# Activate the flake by using nix. This creates a reproducible development environment.
{
  description = "Flake for Artshow Showroom development";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = inputs @ {nixpkgs, ...}: let
    system = "x86_64-linux";
    pkgs = import nixpkgs {inherit system;};
  in {
    devShells.${system} = {
      default = pkgs.mkShell {
        packages = with pkgs; [
          go
        ];

        shellHook = let
          red_text = ''\e[0;31m'';
          green_text = ''\e[0;32m'';
          reset_formatting = ''\e[0m'';
        in ''
          echo -e "\
          ${red_text}go version: ${green_text}$(go version)
          ${reset_formatting}"'';
      };
    };
  };
}
