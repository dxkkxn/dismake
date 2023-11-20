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

```

# Deploy in g5k
First of all change in copyToG5k.sh your login and ssh private key

```
./copyToG5k.sh
./alloc.sh
./deploy.sh
```

