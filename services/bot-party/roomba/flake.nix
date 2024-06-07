{
  description = "A simple flake";

  # Nixpkgs / NixOS version to use.
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      # Generate a user-friendly version number.
      version = builtins.substring 0 8 self.lastModifiedDate;

      # System types to support.
      supportedSystems = [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];

      # Helper function to generate an attrset '{ x86_64-linux = f "x86_64-linux"; ... }'.
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;

      # Nixpkgs instantiated for supported system types.
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });

    in
    {
      # Provide some binary packages for selected system types.
      packages = forAllSystems (system:
        let
          pkgs = nixpkgsFor.${system};
        in
        {
          # The default package for 'nix build'.
          default = pkgs.mkShell {
            buildInputs = [
              # Add packages you need here
              pkgs.python39
              pkgs.uv
#              pkgs.rye
            ];
          };
        });

      # Additional settings can be added here.
      defaultPackage = forAllSystems (system: self.packages.${system}.default);

      # If you want to provide a default application, uncomment and adjust the following lines
      # defaultApp = forAllSystems (system: {
      #   type = "app";
      #   program = "${self.packages.${system}.default}";
      # });
    };
}
