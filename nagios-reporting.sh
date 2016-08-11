#!/bin/sh
# warning: do not move this file as it starts the etherpad session using systemd, see configuration.nix (qknight)
source /etc/profile
cd /home/joachim/nagios-reporting
nix-shell --command "while true; do go run server.go ; done"
