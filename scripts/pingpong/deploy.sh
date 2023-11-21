#!/usr/bin/sh
set -xe

for host in $(uniq $OAR_NODEFILE); do
    if [ $host != $(hostname) ]; then
        ssh $host './install_go.sh 2> /dev/null && cd pingpong && /usr/local/go/bin/go run server/main.go'
    fi
done
#
#
