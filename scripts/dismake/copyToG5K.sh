#!/bin/sh
set -xe

# Dont forget to send the ssh keys
USER=yourusername
SSHKEY=~/.ssh/g5k
DISMAKE_DIR=~/ensimag-dismake/dismake
SCRIPTS_DIR=~/ensimag-dismake/scripts
SITE=lyon

# copy files and connect
rsync -e "ssh -i $SSHKEY" -avP . $USER@access.grid5000.fr:$SITE/project/
ssh -i $SSHKEY $USER@access.grid5000.fr
