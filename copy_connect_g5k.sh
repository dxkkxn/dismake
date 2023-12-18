#!/bin/sh
set -xe

USER=nessalihi
SSHKEY=~/.ssh/id_rsa
SITE=$1

DISMAKE_DIR=/home/nabil/Documents/SYSD/distrib-make/dismake
MAKEFILES_DIR=/home/nabil/Documents/SYSD/distrib-make/makefiles

# copy files and connect
scp -i $SSHKEY -r $DISMAKE_DIR $MAKEFILES_DIR alloc.sh setup.sh make.sh clean.sh auto_deploy.py benchmark.py launch_benchmark.sh $USER@access.grid5000.fr:$SITE/
ssh -i $SSHKEY $USER@access.grid5000.fr
