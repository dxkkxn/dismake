#!/usr/bin/sh
set -xe

USER=ybenjellounelkbibi
SSHKEY=~/.ssh/g5k
PINGPONG_DIR=~/ensimag-parmake/pingpong
SITE=grenoble

# copy files and connect
scp -i $SSHKEY -r $PINGPONG_DIR alloc.sh install_go.sh deploy.sh $USER@access.grid5000.fr:$SITE/
ssh -i $SSHKEY $USER@access.grid5000.fr
