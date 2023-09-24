{
  description = "Nix flake for development";

  inputs = {
    nixpkgs = {
      url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    };

    utils = {
      url = "github:numtide/flake-utils";
    };
  };

  outputs = { self, nixpkgs, utils }:
    utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
          config = { allowUnfree = true; };
        };
      in
      {
        devShell = pkgs.mkShell {
          buildInputs = with pkgs; [
            chart-testing
            helm-docs
            kind
            kubeconform
            kubernetes-helm
            yamllint
            go
          ];
        };
      }
    );
}
