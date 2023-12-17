#!/usr/bin/sh
set -xe

remote_exec="
export PATH=$PATH:~/go/bin;
cd dismake;
go run server/main.go
"
servers=""
for host in $(uniq $OAR_NODEFILE); do
    if [ $host != $(hostname) ]; then
        ssh $host $remote_exec &
        if [ ! $servers ]; then
            servers="$host:50051"
        else
            servers="${servers} $host:50051"
        fi
    fi
done

sleep 10 # wait for the servers to run
export PATH=$PATH:~/go/bin;
go build -C dismake/client
./dismake/client/client -server "$servers" $1
