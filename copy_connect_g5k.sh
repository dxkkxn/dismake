#!/usr/bin/sh
set -xe

USER=ybenjellounelkbibi
SSHKEY=~/.ssh/g5k
SITE=lyon

DISMAKE_DIR=~/ensimag-dismake/dismake
SCRIPTS_DIR=~/ensimag-dismake/scripts
MAKEFILES_DIR=~/ensimag-dismake/makefiles

# copy files and connect
scp -i $SSHKEY -r $DISMAKE_DIR $SCRIPTS_DIR $MAKEFILES_DIR alloc.sh setup.sh make.sh $USER@access.grid5000.fr:$SITE/
ssh -i $SSHKEY $USER@access.grid5000.fr
