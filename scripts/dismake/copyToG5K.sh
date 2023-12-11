#!/usr/bin/sh
set -xe

USER=ybenjellounelkbibi
SSHKEY=~/.ssh/g5k
DISMAKE_DIR=~/ensimag-dismake/dismake
SCRIPTS_DIR=~/ensimag-dismake/scripts
SITE=grenoble

# copy files and connect
scp -i $SSHKEY -r $DISMAKE_DIR $SCRIPTS_DIR $USER@access.grid5000.fr:$SITE/
ssh -i $SSHKEY $USER@access.grid5000.fr
