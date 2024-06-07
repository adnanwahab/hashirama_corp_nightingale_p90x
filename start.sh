#!/usr/bin/env sh

# Exit immediately if a command exits with a non-zero status
set -e


start_service() {
    service_dir=$1
    log_file="$service_dir/start.log"
    cd "$service_dir"
    #nix-shell
    #bash start.sh > "$log_file" 2>&1 &
    bash start.sh &
    cd ~/hashirama

}


./result/bin/gomod2nix-template

#nix build .

#echo "All services started successfully."


# sudo apt-get install supervisor

# sudo cp ./supervisor.conf.d /etc/supervisor/conf.d/

# sudo supervisorctl reread
# sudo supervisorctl update
# sudo supervisorctl start all


# cd ~/hashirama/services/caddy/
# sudo ~/caddy run &
# cd ~/hashirama
# echo "started caddy..."
# # Start other services in parallel
# start_service "services/bot-party"
# start_service "services/planning-prediction"
# start_service "services/hooks"

wait  # Wait for all background jobs to complete

