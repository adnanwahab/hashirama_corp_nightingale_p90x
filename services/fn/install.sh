#!/usr/bin/env sh
curl --proto '=https' --tlsv1.2 -sSf -L https://install.determinate.systems/nix | sh -s -- install
git clone git@github.com:adnanwahab/hashirama_corp_nightingale_p90x.git hashirama
cd hashirama
cp hosts.nix /etc/nix/host.nix
nix build
./result/bin
nix-channel --add https://github.com/nix-community/home-manager/archive/master.tar.gz home-manager
nix-channel --update
curl -fsSL https://tailscale.com/install.sh | sh
sudo tailscale up --authkey  tskey-auth-k1NyTMJ6vX11CNTRL-XJoCa9sNyMAX4NJWhSaFMAw8Aspi1TQ47 --ssh
