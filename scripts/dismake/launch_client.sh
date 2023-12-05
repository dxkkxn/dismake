#!/usr/bin/sh
set -xe
~/scripts/install_go.sh
cd ~/dismake/client
PATH=$PATH:/usr/local/go/bin
make
./client -server test-makefile
