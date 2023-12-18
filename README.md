### Distributed Makefile

- [Class website](http://systemes.pages.ensimag.fr/www-sysd-isi3a)

# Launch tests application in g5k
First of all change in copy_connect_g5k.sh your login and ssh private key

```
./copy_connect_g5k.sh
ssh $SITE # site specified in copyToG5k
./setup.sh
./alloc.sh
cd makefiles
make your_test 
```
