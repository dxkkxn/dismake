### Distributed Makefile

- [Class website](http://systemes.pages.ensimag.fr/www-sysd-isi3a)

# Launch Ping Pong 

```
# dans un terminal
cd pingpong
go run server/main.go

# dans un autre terminal
cd pingpong
go run client/main.go

# ou
Pour faire le histogramme:
cd pingpong
go run client/main.go 2>&1 | python metrics.py

```

# Deploy in g5k
First of all change in copyToG5k.sh your login and ssh private key

```
./copyToG5k.sh
./alloc.sh
./deploy.sh # this should deploy the servers in all avaible nodes
# to launch the client
cd pingpong
/usr/local/go/bin/go run client/main.go -server ${YOUR_SEVER_ADDR}

```

