#!/usr/bin/sh
set -xe
wget -O go1.21.4.linux-amd64.tar.gz https://go.dev/dl/go1.21.4.linux-amd64.tar.gz;
rm -rf ~/go && tar -C ~ -xzf go1.21.4.linux-amd64.tar.gz;
