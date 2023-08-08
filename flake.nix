{
  description = "Flake utils demo";

  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        packages = rec {
          default = import ./default.nix { inherit pkgs; };
        };
        devShells.default = pkgs.mkShell {
          packages = [
            pkgs.go
          ];
        };
      }
    );
}
