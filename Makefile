build-image:
	nix-build '<nixpkgs/nixos>' -A vm -I nixpkgs=channel:nixos-24.05 -I nixos-config=./configuration.nix

run-vm:
	./result/bin/run-nixos-vm

generate-qcow:
	nixos-generate -f qcow -c ./configuration/configuration.nix -o ./output.qcow2

run-image:
	QEMU_KERNEL_PARAMS=console=ttyS0 ./result/bin/run-nixos-vm

fuck:
	nixos-generate -f qcow -c ./configuration.nix
