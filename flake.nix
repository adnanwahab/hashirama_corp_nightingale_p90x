{
  description = "NixOS configurations for my VMs";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
  };

  outputs = { self, nixpkgs }: {
    nixosConfigurations = {
      jetson = nixpkgs.lib.nixosSystem {
        system = "aarch64linux";
        modules = [
          ./hosts/jetson/configuration.nix
        ];
      };
      macbook = nixpkgs.lib.nixosSystem {
        system = "aarch64linux";
        modules = [
          ./hosts/macbook/configuration.nix
        ];
      };

      thinkpad = nixpkgs.lib.nixosSystem {
        system = "x86_64-linux";
        modules = [
          ./hosts/thinkpad/configuration.nix
        ];
      };
    };
  };
}
