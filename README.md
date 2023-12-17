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

# Launch dismake application in g5k
First of all change in copyToG5k.sh your login and ssh private key

```
./scripts/dismake/copyToG5k.sh
ssh $SITE # site specified in copyToG5k
./scripts/install_go.sh
./scripts/alloc.sh
./run_app.sh your_simple_makefile
# this should deploy the servers in all available nodes and launch the application

```

