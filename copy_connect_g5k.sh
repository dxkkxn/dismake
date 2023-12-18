#!/bin/sh
set -xe

USER=gfaccinhuth
SSHKEY=~/.ssh/g5k
SITE=lyon


# copy files and connect
scp -i $SSHKEY -r $DISMAKE_DIR $SCRIPTS_DIR $MAKEFILES_DIR alloc.sh setup.sh make.sh clean.sh $USER@access.grid5000.fr:$SITE/
ssh -i $SSHKEY $USER@access.grid5000.fr
