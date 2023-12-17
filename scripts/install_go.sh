#!/usr/bin/sh
set -xe
wget -O go1.21.4.linux-amd64.tar.gz https://go.dev/dl/go1.21.4.linux-amd64.tar.gz;
chmod -R u+w ~/go && rm -rf ~/go
tar -C ~ -xzf go1.21.4.linux-amd64.tar.gz;
