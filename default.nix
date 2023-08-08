{ pkgs ? import <nixpkgs> {} }:
with pkgs;
buildGoModule {
  pname = "cert-info";
  version = "0.1.0";

  src = pkgs.nix-gitignore.gitignoreSource [] ./.;

  vendorHash = null;
}
