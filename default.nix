{ pkgs ? import <nixpkgs> {} }:
with pkgs;
buildGoModule {
  pname = "cert-info";
  version = "0.1.0";

  src = ./.;

  vendorHash = null;
}
