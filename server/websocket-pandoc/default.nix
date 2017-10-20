{ pkgs ? import <nixpkgs>{}, ... }:

pkgs.buildGoPackage rec {
  name = "pankat-${version}";
  version = "0.0.1";

  src = ./src/github.com/nixcloud/pankat;

  goPackagePath = "github.com/nixcloud/pankat";
  goDeps = ./deps.nix;

  buildInputs = [ pkgs.pandoc ];
}
