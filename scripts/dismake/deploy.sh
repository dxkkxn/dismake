#!/usr/bin/sh
set -xe

remote_exec="
export PATH=$PATH:~/go/bin;
cd dismake;
go run server/main.go
"

for host in $(uniq $OAR_NODEFILE); do
    if [ $host != $(hostname) ]; then
        ssh $host $remote_exec &
    fi
done
