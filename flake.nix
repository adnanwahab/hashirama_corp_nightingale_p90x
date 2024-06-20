{
  description = "Deployment for my server cluster";

  # For accessing `deploy-rs`'s utility Nix functions
  inputs.deploy-rs.url = "github:serokell/deploy-rs";

  outputs = { self, nixpkgs, deploy-rs }: {
    nixosConfigurations.some-random-system = nixpkgs.lib.nixosSystem {
      system = "x86_64-linux";
      modules = [ ./configuration/configuration.nix ];
    };

    deploy.nodes.orin = {
        hostname = "ubuntu";
        profiles.system = {
          user = "adnan";
          path = deploy-rs.lib.x86_64-linux.activate.nixos self.nixosConfigurations.some-random-system;
        };
    };

    # This is highly advised, and will prevent many possible mistakes
    checks = builtins.mapAttrs (system: deployLib: deployLib.deployChecks self.deploy) deploy-rs.lib;
  };
}
