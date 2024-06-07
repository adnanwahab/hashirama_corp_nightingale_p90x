#!/usr/bin/env sh




# nix develop
#   rye init
#   . .venv/bin/activate
#   rye sync
#   uv pip install -r requirements.txt
#   python docking.py



# nix develop <<EOF
# uv pip install -r requirements.txt
# uv venv
# source .venv/bin/activate
# python docking.py
# EOF


#
#python3.9 -m venv venv
#source venv/bin/activate
#
#   source "$HOME/.rye/env"

uv venv
uv pip install -r requirements.txt
python3 docking.py
