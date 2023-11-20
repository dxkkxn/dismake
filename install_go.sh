#!/usr/bin/sh
set -xe

wget https://go.dev/dl/go1.21.4.linux-amd64.tar.gz;
rm -rf /usr/local/go && sudo-g5k tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz;
PATH=$PATH:/usr/local/go/bin;
