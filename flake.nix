{
  description = "NixOS configurations for my VMs";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
        flake-utils.url = "github:numtide/flake-utils";

  };

    outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        packages = {
          default = pkgs.stdenv.mkDerivation {
            pname = "my-package";
            version = "1.0.0";
            src = ./.;

            buildInputs = [ pkgs.hello ];

            buildPhase = ''
              echo "Building my package..."
              # Add your build steps here
   # Simulate build process (replace this with actual build steps)
              echo "#!/bin/sh" > my-executable
              echo "echo Hello, World!" >> my-executable
              chmod +x my-executable
            '';

            installPhase = ''
              mkdir -p $out/bin
              cp my-executable $out/bin/
            '';
          };
        };
      }
    );
  #   nixosConfigurations = {
  #     jetson = nixpkgs.lib.nixosSystem {
  #       system = "aarch64linux";
  #       modules = [
  #         ./hosts/jetson/configuration.nix
  #       ];
  #     };
  #     macbook = nixpkgs.lib.nixosSystem {
  #       system = "aarch64linux";
  #       modules = [
  #         ./hosts/macbook/configuration.nix
  #       ];
  #     };

  #     thinkpad = nixpkgs.lib.nixosSystem {
  #       system = "x86_64-linux";
  #       modules = [
  #         ./hosts/thinkpad/configuration.nix
  #       ];
  #     };
  #   };
  # };
}
